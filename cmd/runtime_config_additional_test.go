package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRuntimeConfigStageError_NilReceiverAndNilInnerError(t *testing.T) {
	var nilErr *runtimeConfigStageError
	if got := nilErr.Error(); got != "" {
		t.Fatalf("expected nil receiver Error() to return empty string, got %q", got)
	}
	if unwrapped := nilErr.Unwrap(); unwrapped != nil {
		t.Fatalf("expected nil receiver Unwrap() to return nil, got %v", unwrapped)
	}
	if nilErr.Is(errRuntimeConfigLoad) {
		t.Fatalf("expected nil receiver Is(...) to return false")
	}

	nilInner := &runtimeConfigStageError{stage: errRuntimeConfigLoad}
	if got := nilInner.Error(); got != "" {
		t.Fatalf("expected nil inner error to return empty message, got %q", got)
	}
}

func TestWrapRuntimeConfigError_ReturnsNilWhenInnerErrorIsNil(t *testing.T) {
	if wrapped := wrapRuntimeConfigError(errRuntimeConfigLoad, nil); wrapped != nil {
		t.Fatalf("expected nil error to stay nil, got %v", wrapped)
	}
}

func TestResolveConfigPathFromArgs_MissingConfigValueFallsBackToDefault(t *testing.T) {
	if got := resolveConfigPathFromArgs([]string{"info", "--config"}, "config.yaml"); got != "config.yaml" {
		t.Fatalf("expected long --config without value to fall back to default, got %q", got)
	}

	if got := resolveConfigPathFromArgs([]string{"info", "-c"}, "config.yaml"); got != "config.yaml" {
		t.Fatalf("expected short -c without value to fall back to default, got %q", got)
	}
}

func TestRegisterOverrideLikeFlagsFromArgs_NilCommandIsNoOp(t *testing.T) {
	registerOverrideLikeFlagsFromArgs(nil, []string{"--defaults.revision=release"})
}

func TestRegisterDynamicOverrideFlags_NilInputsAreNoOps(t *testing.T) {
	registerDynamicOverrideFlags(nil, nil)
	registerDynamicOverrideFlags(&cobra.Command{Use: "test"}, nil)
}

func TestApplyDynamicOverrides_NilInputsAreNoOps(t *testing.T) {
	applyDynamicOverrides(nil, nil)
	applyDynamicOverrides(&cobra.Command{Use: "test"}, nil)
}

func TestGetChangedStringFlag_SupportsCommandFlagsAndAbsentFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.Flags().String("repos.tool.path", "", "")

	if err := cmd.Flags().Set("repos.tool.path", "workspace/tool"); err != nil {
		t.Fatalf("setting local command flag failed: %v", err)
	}

	if value, ok := getChangedStringFlag(cmd, "repos.tool.path"); !ok || value != "workspace/tool" {
		t.Fatalf("expected changed local command flag value, got value=%q changed=%v", value, ok)
	}

	if value, ok := getChangedStringFlag(cmd, "repos.tool.revision"); ok || value != "" {
		t.Fatalf("expected missing flag lookup to return not changed, got value=%q changed=%v", value, ok)
	}
}

func TestCommands_PrintRuntimeConfigLoadErrorWhenConfigMissing(t *testing.T) {
	missingConfig := filepath.Join(t.TempDir(), "missing.yaml")
	previousCfgFile := cfgFile
	cfgFile = missingConfig
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	tests := []struct {
		name string
		run  func(*cobra.Command, []string)
	}{
		{name: "info", run: infoCmd.Run},
		{name: "clone", run: cloneCmd.Run},
		{name: "update", run: updateCmd.Run},
		{name: "cleanup", run: cleanupCmd.Run},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			output := captureStdout(t, func() {
				tc.run(&cobra.Command{Use: tc.name}, nil)
			})

			if !strings.Contains(output, "Error loading config") {
				t.Fatalf("expected %s command to print config load failure, got:\n%s", tc.name, output)
			}
		})
	}
}
