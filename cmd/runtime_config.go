package cmd

import (
	"errors"
	"fmt"
	"strings"

	"git-downloader-tool/config"
	"github.com/spf13/cobra"
)

var (
	errRuntimeConfigLoad     = errors.New("runtime config load stage")
	errRuntimeConfigMerge    = errors.New("runtime config merge stage")
	errRuntimeConfigValidate = errors.New("runtime config validate stage")
)

type runtimeConfigStageError struct {
	stage error
	err   error
}

func (e *runtimeConfigStageError) Error() string {
	if e == nil || e.err == nil {
		return ""
	}

	return e.err.Error()
}

func (e *runtimeConfigStageError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.err
}

func (e *runtimeConfigStageError) Is(target error) bool {
	if e == nil {
		return false
	}

	return target == e.stage
}

func resolveConfigPathFromArgs(args []string, defaultValue string) string {
	resolved := defaultValue

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--" {
			break
		}

		if strings.HasPrefix(arg, "--config=") {
			resolved = strings.TrimPrefix(arg, "--config=")
			continue
		}

		if strings.HasPrefix(arg, "-c=") {
			resolved = strings.TrimPrefix(arg, "-c=")
			continue
		}

		if arg == "--config" || arg == "-c" {
			if i+1 < len(args) {
				resolved = args[i+1]
				i++
				continue
			}
			resolved = defaultValue
		}
	}

	return resolved
}

func registerOverrideLikeFlagsFromArgs(cmd *cobra.Command, args []string) {
	if cmd == nil {
		return
	}

	for _, arg := range args {
		if arg == "--" {
			break
		}

		if !strings.HasPrefix(arg, "--") {
			continue
		}

		flagName := strings.TrimPrefix(arg, "--")
		if index := strings.Index(flagName, "="); index >= 0 {
			flagName = flagName[:index]
		}

		if !isOverrideLikeFlagName(flagName) {
			continue
		}

		registerStringPersistentFlag(
			cmd,
			flagName,
			"",
			fmt.Sprintf("[DYNAMIC] Override %s", flagName),
		)
	}
}

func isOverrideLikeFlagName(flagName string) bool {
	return strings.HasPrefix(flagName, "defaults.") ||
		strings.HasPrefix(flagName, "remote.") ||
		strings.HasPrefix(flagName, "repos.")
}

func registerDynamicOverrideFlags(cmd *cobra.Command, loadedConfig *config.Config) {
	if cmd == nil || loadedConfig == nil {
		return
	}

	for remoteName, remoteConfig := range loadedConfig.Remotes {
		registerStringPersistentFlag(
			cmd,
			fmt.Sprintf("remote.%s.url", remoteName),
			remoteConfig.URL,
			fmt.Sprintf("[DYNAMIC] Override remote.%s.url", remoteName),
		)
	}

	registerStringPersistentFlag(cmd, "defaults.remote", loadedConfig.Default.Remote, "[DYNAMIC] Override defaults.remote")
	registerStringPersistentFlag(cmd, "defaults.revision", loadedConfig.Default.Revision, "[DYNAMIC] Override defaults.revision")
	registerStringPersistentFlag(cmd, "defaults.path", loadedConfig.Default.Path, "[DYNAMIC] Override defaults.path")

	for repoName, repoConfig := range loadedConfig.Repos {
		registerStringPersistentFlag(
			cmd,
			fmt.Sprintf("repos.%s.remote", repoName),
			repoConfig.Remote,
			fmt.Sprintf("[DYNAMIC] Override repos.%s.remote", repoName),
		)
		registerStringPersistentFlag(
			cmd,
			fmt.Sprintf("repos.%s.repository", repoName),
			repoConfig.Repository,
			fmt.Sprintf("[DYNAMIC] Override repos.%s.repository", repoName),
		)
		registerStringPersistentFlag(
			cmd,
			fmt.Sprintf("repos.%s.revision", repoName),
			repoConfig.Revision,
			fmt.Sprintf("[DYNAMIC] Override repos.%s.revision", repoName),
		)
		registerStringPersistentFlag(
			cmd,
			fmt.Sprintf("repos.%s.path", repoName),
			repoConfig.Path,
			fmt.Sprintf("[DYNAMIC] Override repos.%s.path", repoName),
		)
	}
}

func registerStringPersistentFlag(cmd *cobra.Command, name, value, usage string) {
	if cmd.PersistentFlags().Lookup(name) == nil {
		cmd.PersistentFlags().String(name, value, usage)
	}
}

func applyDynamicOverrides(cmd *cobra.Command, loadedConfig *config.Config) {
	if cmd == nil || loadedConfig == nil {
		return
	}

	for remoteName, remoteConfig := range loadedConfig.Remotes {
		flagName := fmt.Sprintf("remote.%s.url", remoteName)
		flagValue, ok := getChangedStringFlag(cmd, flagName)
		if !ok {
			continue
		}

		remoteConfig.URL = flagValue
		loadedConfig.Remotes[remoteName] = remoteConfig
	}

	if value, ok := getChangedStringFlag(cmd, "defaults.remote"); ok {
		loadedConfig.Default.Remote = value
	}

	if value, ok := getChangedStringFlag(cmd, "defaults.revision"); ok {
		loadedConfig.Default.Revision = value
	}

	if value, ok := getChangedStringFlag(cmd, "defaults.path"); ok {
		loadedConfig.Default.Path = value
	}

	for repoName, repoConfig := range loadedConfig.Repos {
		remoteName := fmt.Sprintf("repos.%s.remote", repoName)
		if value, ok := getChangedStringFlag(cmd, remoteName); ok {
			repoConfig.Remote = value
		}

		repositoryName := fmt.Sprintf("repos.%s.repository", repoName)
		if value, ok := getChangedStringFlag(cmd, repositoryName); ok {
			repoConfig.Repository = value
		}

		revisionName := fmt.Sprintf("repos.%s.revision", repoName)
		if value, ok := getChangedStringFlag(cmd, revisionName); ok {
			repoConfig.Revision = value
		}

		pathName := fmt.Sprintf("repos.%s.path", repoName)
		if value, ok := getChangedStringFlag(cmd, pathName); ok {
			repoConfig.Path = value
		}

		loadedConfig.Repos[repoName] = repoConfig
	}
}

func buildEffectiveConfig(cmd *cobra.Command) (*config.Config, error) {
	loadedConfig, err := config.LoadConfig(cfgFile)
	if err != nil {
		return nil, wrapRuntimeConfigError(errRuntimeConfigLoad, err)
	}

	applyDynamicOverrides(cmd, loadedConfig)

	err = loadedConfig.MergeWithDefaults()
	if err != nil {
		return nil, wrapRuntimeConfigError(errRuntimeConfigMerge, err)
	}

	err = loadedConfig.Validate()
	if err != nil {
		return nil, wrapRuntimeConfigError(errRuntimeConfigValidate, err)
	}

	return loadedConfig, nil
}

func wrapRuntimeConfigError(stage error, err error) error {
	if err == nil {
		return nil
	}

	return &runtimeConfigStageError{stage: stage, err: err}
}

func runtimeConfigErrorMessage(err error) string {
	prefix := "Error loading config"

	switch {
	case errors.Is(err, errRuntimeConfigMerge):
		prefix = "Error merging config"
	case errors.Is(err, errRuntimeConfigValidate):
		prefix = "Error validating config"
	case errors.Is(err, errRuntimeConfigLoad):
		prefix = "Error loading config"
	}

	return fmt.Sprintf("%s: %v", prefix, err)
}

func getChangedStringFlag(cmd *cobra.Command, name string) (string, bool) {
	flags := cmd.Flags()
	if flags.Lookup(name) != nil && flags.Changed(name) {
		value, err := flags.GetString(name)
		if err == nil {
			return value, true
		}
	}

	persistentFlags := cmd.PersistentFlags()
	if persistentFlags.Lookup(name) != nil && persistentFlags.Changed(name) {
		value, err := persistentFlags.GetString(name)
		if err == nil {
			return value, true
		}
	}

	inheritedFlags := cmd.InheritedFlags()
	if inheritedFlags.Lookup(name) != nil && inheritedFlags.Changed(name) {
		value, err := inheritedFlags.GetString(name)
		if err == nil {
			return value, true
		}
	}

	return "", false
}
