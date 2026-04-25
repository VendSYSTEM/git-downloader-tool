package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/VendSYSTEM/git-downloader-tool/config"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func TestPreregisterOverrideFlagsFromArgs_RegistersMatchingLongFlags(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}
	registerOverrideLikeFlagsFromArgs(cmd, []string{
		"info",
		"--defaults.revision=release",
		"--remote.origin.url",
		"https://example.com/mirror.git",
		"--repos.some-tool.path=workspace/some-tool",
		"--not-an-override=value",
	})

	for _, flagName := range []string{"defaults.revision", "remote.origin.url", "repos.some-tool.path"} {
		if cmd.PersistentFlags().Lookup(flagName) == nil {
			t.Fatalf("expected preregistered override-like flag %q", flagName)
		}
	}

	if cmd.PersistentFlags().Lookup("not-an-override") != nil {
		t.Fatalf("did not expect non override flag to be preregistered")
	}
}

func TestPreregisterOverrideFlagsFromArgs_StopsAtEndOfOptionsMarker(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}
	registerOverrideLikeFlagsFromArgs(cmd, []string{
		"info",
		"--defaults.revision=release",
		"--",
		"--repos.post-marker.path=ignored",
	})

	if cmd.PersistentFlags().Lookup("defaults.revision") == nil {
		t.Fatalf("expected defaults.revision to be preregistered before -- marker")
	}

	if cmd.PersistentFlags().Lookup("repos.post-marker.path") != nil {
		t.Fatalf("did not expect override-like flags after -- marker to be preregistered")
	}
}

func TestExecute_AllowsOverrideFlagsWhenConfigLoadFails(t *testing.T) {
	tempDir := t.TempDir()
	missingConfigPath := filepath.Join(tempDir, "missing-config.yaml")

	originalArgs := os.Args
	os.Args = []string{
		"git-downloader-tool",
		"info",
		"--config", missingConfigPath,
		"--defaults.revision=release",
		"--remote.origin.url=https://example.com/mirror.git",
		"--repos.lifecycle_check.path=workspace/lifecycle-check",
	}
	t.Cleanup(func() {
		os.Args = originalArgs
	})

	output := captureStdout(t, func() {
		if err := Execute(); err != nil {
			t.Fatalf("Execute returned error: %v", err)
		}
	})

	if strings.Contains(output, "unknown flag") {
		t.Fatalf("expected override flags to be parseable even when config loading fails, got output:\n%s", output)
	}

	if !strings.Contains(output, "Error loading config") {
		t.Fatalf("expected runtime config load error output, got:\n%s", output)
	}
}

func TestResolveConfigPathFromArgs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		args         []string
		defaultValue string
		expected     string
	}{
		{
			name:         "defaults when no config flag provided",
			args:         []string{"info"},
			defaultValue: "config.yaml",
			expected:     "config.yaml",
		},
		{
			name:         "supports long config with value",
			args:         []string{"info", "--config", "custom.yaml"},
			defaultValue: "config.yaml",
			expected:     "custom.yaml",
		},
		{
			name:         "supports short config alias",
			args:         []string{"info", "-c", "custom-short.yaml"},
			defaultValue: "config.yaml",
			expected:     "custom-short.yaml",
		},
		{
			name:         "supports long config equals syntax",
			args:         []string{"info", "--config=inline.yaml"},
			defaultValue: "config.yaml",
			expected:     "inline.yaml",
		},
		{
			name:         "supports short config equals syntax",
			args:         []string{"info", "-c=inline-short.yaml"},
			defaultValue: "config.yaml",
			expected:     "inline-short.yaml",
		},
		{
			name:         "uses last repeated config flag value",
			args:         []string{"info", "--config", "first.yaml", "--config", "second.yaml"},
			defaultValue: "config.yaml",
			expected:     "second.yaml",
		},
		{
			name:         "uses last repeated mixed config flags value",
			args:         []string{"info", "--config=first.yaml", "-c", "second.yaml", "-c=third.yaml"},
			defaultValue: "config.yaml",
			expected:     "third.yaml",
		},
		{
			name:         "ignores config flags after end of options marker",
			args:         []string{"info", "--config", "before.yaml", "--", "--config", "after.yaml"},
			defaultValue: "config.yaml",
			expected:     "before.yaml",
		},
		{
			name:         "returns default when config appears only after end of options marker",
			args:         []string{"info", "--", "--config", "after.yaml"},
			defaultValue: "config.yaml",
			expected:     "config.yaml",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := resolveConfigPathFromArgs(tc.args, tc.defaultValue)
			if actual != tc.expected {
				t.Fatalf("expected config path %q, got %q", tc.expected, actual)
			}
		})
	}
}

func TestRegisterDynamicOverrideFlags_UsesReposNamespace(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}
	loadedConfig := &config.Config{
		Remotes: map[string]config.Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: config.Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]config.Repository{
			"tool": {
				Remote:     "origin",
				Repository: "team/tool.git",
				Revision:   "main",
				Path:       "repos/tool",
			},
		},
	}

	registerDynamicOverrideFlags(cmd, loadedConfig)

	expectedFlags := []string{
		"remote.origin.url",
		"defaults.remote",
		"defaults.revision",
		"defaults.path",
		"repos.tool.remote",
		"repos.tool.repository",
		"repos.tool.revision",
		"repos.tool.path",
	}

	for _, flagName := range expectedFlags {
		if cmd.PersistentFlags().Lookup(flagName) == nil {
			t.Fatalf("expected dynamic flag %q to be registered", flagName)
		}
	}

	if cmd.PersistentFlags().Lookup("repo.tool.remote") != nil {
		t.Fatalf("did not expect legacy repo.* namespace flag to be registered")
	}
}

func TestRegisterDynamicOverrideFlags_CanBeCalledTwice(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}
	loadedConfig := &config.Config{
		Remotes: map[string]config.Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: config.Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]config.Repository{
			"tool": {
				Remote:     "origin",
				Repository: "team/tool.git",
				Revision:   "main",
				Path:       "repos/tool",
			},
		},
	}

	registerDynamicOverrideFlags(cmd, loadedConfig)
	registerDynamicOverrideFlags(cmd, loadedConfig)

	flagCount := 0
	cmd.PersistentFlags().VisitAll(func(_ *pflag.Flag) {
		flagCount++
	})

	const expectedFlagCount = 8
	if flagCount != expectedFlagCount {
		t.Fatalf("expected %d dynamic flags after repeated registration, got %d", expectedFlagCount, flagCount)
	}
}

func TestRegisterStringPersistentFlag_ReRegisterDoesNotMarkChanged(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}

	registerStringPersistentFlag(cmd, "defaults.remote", "origin", "dynamic defaults remote")
	if cmd.PersistentFlags().Changed("defaults.remote") {
		t.Fatalf("expected defaults.remote to be unchanged after first registration")
	}

	registerStringPersistentFlag(cmd, "defaults.remote", "upstream", "dynamic defaults remote")
	if cmd.PersistentFlags().Changed("defaults.remote") {
		t.Fatalf("expected defaults.remote to remain unchanged after re-registration")
	}
}

func TestApplyDynamicOverrides_OnlyChangedFlags(t *testing.T) {
	t.Parallel()

	cmd := &cobra.Command{Use: "test"}
	loadedConfig := &config.Config{
		Remotes: map[string]config.Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: config.Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]config.Repository{
			"tool": {
				Remote:     "origin",
				Repository: "team/tool.git",
				Revision:   "main",
				Path:       "repos/tool",
			},
		},
	}

	registerDynamicOverrideFlags(cmd, loadedConfig)

	if err := cmd.PersistentFlags().Set("remote.origin.url", "https://mirror.example.com/"); err != nil {
		t.Fatalf("setting changed remote flag failed: %v", err)
	}
	if err := cmd.PersistentFlags().Set("repos.tool.revision", "release"); err != nil {
		t.Fatalf("setting changed repo flag failed: %v", err)
	}

	applyDynamicOverrides(cmd, loadedConfig)

	if got := loadedConfig.Remotes["origin"].URL; got != "https://mirror.example.com/" {
		t.Fatalf("expected changed remote URL to be applied, got %q", got)
	}

	if got := loadedConfig.Repos["tool"].Revision; got != "release" {
		t.Fatalf("expected changed repo revision to be applied, got %q", got)
	}

	if got := loadedConfig.Defaults.Path; got != "repos" {
		t.Fatalf("expected unchanged defaults.path to stay original, got %q", got)
	}

	if got := loadedConfig.Repos["tool"].Path; got != "repos/tool" {
		t.Fatalf("expected unchanged repos.tool.path to stay original, got %q", got)
	}
}

func TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags(t *testing.T) {
	t.Parallel()

	rootCmd := &cobra.Command{Use: "root"}
	subCmd := &cobra.Command{Use: "sub"}
	rootCmd.AddCommand(subCmd)

	loadedConfig := &config.Config{
		Remotes: map[string]config.Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: config.Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]config.Repository{
			"tool": {
				Remote:     "origin",
				Repository: "team/tool.git",
				Revision:   "main",
				Path:       "repos/tool",
			},
		},
	}

	registerDynamicOverrideFlags(rootCmd, loadedConfig)

	if err := rootCmd.PersistentFlags().Set("defaults.revision", "release"); err != nil {
		t.Fatalf("setting changed root persistent flag failed: %v", err)
	}

	applyDynamicOverrides(subCmd, loadedConfig)

	if got := loadedConfig.Defaults.Revision; got != "release" {
		t.Fatalf("expected defaults.revision override from inherited flag, got %q", got)
	}
}

func TestBuildEffectiveConfig_LoadApplyMergeValidate(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	configContents := []byte(`remotes:
  origin:
    url: https://example.com/
defaults:
  remote: origin
  revision: main
  path: repos
repos:
  tool:
    repository: team/tool.git
`)

	if err := os.WriteFile(configPath, configContents, 0o644); err != nil {
		t.Fatalf("failed writing test config: %v", err)
	}

	previousCfgFile := cfgFile
	cfgFile = configPath
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	cmd := &cobra.Command{Use: "test"}
	loadedConfig, err := config.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("failed loading config fixture: %v", err)
	}

	registerDynamicOverrideFlags(cmd, loadedConfig)

	if err := cmd.PersistentFlags().Set("defaults.revision", "release"); err != nil {
		t.Fatalf("setting changed defaults flag failed: %v", err)
	}
	if err := cmd.PersistentFlags().Set("repos.tool.path", "workspace/tool"); err != nil {
		t.Fatalf("setting changed repo path flag failed: %v", err)
	}

	effective, err := buildEffectiveConfig(cmd)
	if err != nil {
		t.Fatalf("expected buildEffectiveConfig to succeed, got error: %v", err)
	}

	if got := effective.Defaults.Revision; got != "release" {
		t.Fatalf("expected defaults.revision override to be applied, got %q", got)
	}

	if got := effective.Repos["tool"].Revision; got != "release" {
		t.Fatalf("expected defaults.revision to flow into repo during merge, got %q", got)
	}

	if got := effective.Repos["tool"].Path; got != "workspace/tool" {
		t.Fatalf("expected repos.tool.path override to be applied before merge, got %q", got)
	}
}

func TestBuildEffectiveConfig_WrapsLoadStageErrors(t *testing.T) {
	previousCfgFile := cfgFile
	cfgFile = filepath.Join(t.TempDir(), "missing-config.yaml")
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	_, err := buildEffectiveConfig(&cobra.Command{Use: "test"})
	if err == nil {
		t.Fatalf("expected load-stage error, got nil")
	}

	if !errors.Is(err, errRuntimeConfigLoad) {
		t.Fatalf("expected load-stage sentinel, got %v", err)
	}
}

func TestBuildEffectiveConfig_WrapsValidateStageErrors(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	configContents := []byte(`remotes:
  origin:
    url: https://example.com/
defaults:
  remote: origin
  revision: main
  path: repos
repos:
  broken:
    repository: ""
`)

	if err := os.WriteFile(configPath, configContents, 0o644); err != nil {
		t.Fatalf("failed writing test config: %v", err)
	}

	previousCfgFile := cfgFile
	cfgFile = configPath
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	_, err := buildEffectiveConfig(&cobra.Command{Use: "test"})
	if err == nil {
		t.Fatalf("expected validate-stage error, got nil")
	}

	if !errors.Is(err, errRuntimeConfigValidate) {
		t.Fatalf("expected validate-stage sentinel, got %v", err)
	}
}

func TestBuildEffectiveConfig_RegressionCasesFromFixture(t *testing.T) {
	fixtureConfigPath := filepath.Join("testdata", "config.yaml")

	tests := []struct {
		name   string
		setup  func(t *testing.T, cmd *cobra.Command)
		assert func(t *testing.T, effective *config.Config, err error)
	}{
		{
			name: "repo field override applied for api-gateway path",
			setup: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()

				if err := cmd.PersistentFlags().Set("repos.api-gateway.path", "services/edge"); err != nil {
					t.Fatalf("setting repos.api-gateway.path failed: %v", err)
				}
			},
			assert: func(t *testing.T, effective *config.Config, err error) {
				t.Helper()

				if err != nil {
					t.Fatalf("expected buildEffectiveConfig to succeed, got error: %v", err)
				}

				if got := effective.Repos["api-gateway"].Path; got != "services/edge" {
					t.Fatalf("expected repos.api-gateway.path override, got %q", got)
				}
			},
		},
		{
			name: "repo field overrides applied for revision and path",
			setup: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()

				if err := cmd.PersistentFlags().Set("repos.beta.revision", "release"); err != nil {
					t.Fatalf("setting repos.beta.revision failed: %v", err)
				}
				if err := cmd.PersistentFlags().Set("repos.beta.path", "workspace/beta"); err != nil {
					t.Fatalf("setting repos.beta.path failed: %v", err)
				}
			},
			assert: func(t *testing.T, effective *config.Config, err error) {
				t.Helper()

				if err != nil {
					t.Fatalf("expected buildEffectiveConfig to succeed, got error: %v", err)
				}

				if got := effective.Repos["beta"].Revision; got != "release" {
					t.Fatalf("expected repos.beta.revision override, got %q", got)
				}

				if got := effective.Repos["beta"].Path; got != "workspace/beta" {
					t.Fatalf("expected repos.beta.path override, got %q", got)
				}
			},
		},
		{
			name: "validation fails when defaults.remote points to unknown key",
			setup: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()

				if err := cmd.PersistentFlags().Set("defaults.remote", "missing-remote"); err != nil {
					t.Fatalf("setting defaults.remote failed: %v", err)
				}
			},
			assert: func(t *testing.T, _ *config.Config, err error) {
				t.Helper()

				if err == nil {
					t.Fatalf("expected validation-stage error, got nil")
				}

				if !errors.Is(err, errRuntimeConfigValidate) {
					t.Fatalf("expected validate-stage sentinel, got %v", err)
				}

				if !strings.Contains(err.Error(), "default remote 'missing-remote' is not defined") {
					t.Fatalf("expected unknown default remote validation message, got: %v", err)
				}
			},
		},
		{
			name: "unchanged flags preserve yaml values",
			setup: func(t *testing.T, cmd *cobra.Command) {
				t.Helper()

				if err := cmd.PersistentFlags().Set("defaults.revision", "hotfix"); err != nil {
					t.Fatalf("setting defaults.revision failed: %v", err)
				}
			},
			assert: func(t *testing.T, effective *config.Config, err error) {
				t.Helper()

				if err != nil {
					t.Fatalf("expected buildEffectiveConfig to succeed, got error: %v", err)
				}

				if got := effective.Repos["api-gateway"].Remote; got != "origin" {
					t.Fatalf("expected unchanged repos.api-gateway.remote yaml value, got %q", got)
				}

				if got := effective.Repos["api-gateway"].Revision; got != "main" {
					t.Fatalf("expected unchanged repos.api-gateway.revision yaml value, got %q", got)
				}

				if got := effective.Repos["api-gateway"].Path; got != "services/api-gateway" {
					t.Fatalf("expected unchanged repos.api-gateway.path yaml value, got %q", got)
				}

				if got := effective.Repos["alpha"].Revision; got != "release-candidate" {
					t.Fatalf("expected unchanged repos.alpha.revision yaml value, got %q", got)
				}

				if got := effective.Repos["alpha"].Path; got != "custom/alpha" {
					t.Fatalf("expected unchanged repos.alpha.path yaml value, got %q", got)
				}
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			previousCfgFile := cfgFile
			cfgFile = fixtureConfigPath
			t.Cleanup(func() {
				cfgFile = previousCfgFile
			})

			cmd := &cobra.Command{Use: "test"}
			loadedConfig, err := config.LoadConfig(fixtureConfigPath)
			if err != nil {
				t.Fatalf("failed loading fixture config: %v", err)
			}

			registerDynamicOverrideFlags(cmd, loadedConfig)
			tc.setup(t, cmd)

			effective, err := buildEffectiveConfig(cmd)
			tc.assert(t, effective, err)
		})
	}
}

func TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression(t *testing.T) {
	fixtureConfigPath := filepath.Join("testdata", "config.yaml")

	previousCfgFile := cfgFile
	cfgFile = fixtureConfigPath
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	cmd := &cobra.Command{Use: "test"}
	registerOverrideLikeFlagsFromArgs(cmd, []string{
		"info",
		"--repos.beta.path=workspace/from-arg",
		"--defaults.revision=release-from-arg",
	})

	loadedConfig, err := config.LoadConfig(fixtureConfigPath)
	if err != nil {
		t.Fatalf("failed loading fixture config: %v", err)
	}

	registerDynamicOverrideFlags(cmd, loadedConfig)

	if err := cmd.PersistentFlags().Set("repos.beta.path", "workspace/from-arg"); err != nil {
		t.Fatalf("setting preregistered repos.beta.path failed: %v", err)
	}
	if err := cmd.PersistentFlags().Set("defaults.revision", "release-from-arg"); err != nil {
		t.Fatalf("setting preregistered defaults.revision failed: %v", err)
	}

	effective, err := buildEffectiveConfig(cmd)
	if err != nil {
		t.Fatalf("expected buildEffectiveConfig to succeed, got error: %v", err)
	}

	if got := effective.Repos["beta"].Path; got != "workspace/from-arg" {
		t.Fatalf("expected preregistered repos.beta.path override to win, got %q", got)
	}

	if got := effective.Defaults.Revision; got != "release-from-arg" {
		t.Fatalf("expected preregistered defaults.revision override to win, got %q", got)
	}

	if got := effective.Repos["alpha"].Path; got != "custom/alpha" {
		t.Fatalf("expected untouched alpha path to keep yaml value, got %q", got)
	}
}

func TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix(t *testing.T) {
	t.Parallel()

	mergeErr := wrapRuntimeConfigError(errRuntimeConfigMerge, errors.New("merge failed"))
	validateErr := wrapRuntimeConfigError(errRuntimeConfigValidate, errors.New("validate failed"))
	loadErr := wrapRuntimeConfigError(errRuntimeConfigLoad, errors.New("load failed"))

	if got := runtimeConfigErrorMessage(loadErr); got != "Error loading config: load failed" {
		t.Fatalf("unexpected load message: %q", got)
	}

	if got := runtimeConfigErrorMessage(mergeErr); got != "Error merging config: merge failed" {
		t.Fatalf("unexpected merge message: %q", got)
	}

	if got := runtimeConfigErrorMessage(validateErr); got != "Error validating config: validate failed" {
		t.Fatalf("unexpected validate message: %q", got)
	}
}
