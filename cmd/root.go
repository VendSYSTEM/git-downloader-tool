package cmd

import (
	"os"

	"git-downloader-tool/config"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	cfg     = config.Config{
		Remotes:  make(map[string]config.Remote),
		Defaults: config.Defaults{},
		Repos:    make(map[string]config.Repository),
	}

	rootCmd = &cobra.Command{
		Use:   "git-downloader-tool",
		Short: "A tool to manage git repositories via YAML configuration",
		Long: `git-downloader-tool is a command-line utility that manages multiple git repositories
through YAML configuration files. It supports clone, update, and cleanup operations.`,
		Run: func(cmd *cobra.Command, args []string) {
			// If no command is provided, show help
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "[DYNAMIC] config file")
}

func initConfig() {
}

func Execute() error {
	resolvedCfgPath := resolveConfigPathFromArgs(os.Args[1:], "config.yaml")
	cfgFile = resolvedCfgPath
	_ = rootCmd.PersistentFlags().Set("config", resolvedCfgPath)
	registerOverrideLikeFlagsFromArgs(rootCmd, os.Args[1:])

	loadedConfig, err := config.LoadConfig(resolvedCfgPath)
	if err == nil {
		registerDynamicOverrideFlags(rootCmd, loadedConfig)
	}

	return rootCmd.Execute()
}
