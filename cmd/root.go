package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/system66/git-downloader-tool/config"
)

var (
	cfgFile string
	verbose bool
	tmpCfg  *config.Config
	cfg     = config.Config{
		Remotes: make(map[string]config.Remote),
		Default: config.Defaults{},
		Repos:   make(map[string]config.Repository),
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
	// rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")

	// // Load config
	// tmpCfg, err := config.LoadConfig(cfgFile)
	// if err == nil {
	// 	for remoteName, remoteConfig := range tmpCfg.Remotes {
	// 		rootCmd.PersistentFlags().StringVar(&remoteConfig.URL, fmt.Sprintf("remote.%s.url", remoteName), remoteConfig.URL, fmt.Sprintf("[DYNAMIC] Override remote.%s.url", remoteName))
	// 		cfg.Remotes[remoteName] = remoteConfig
	// 	}

	// 	rootCmd.PersistentFlags().StringVar(&cfg.Default.Remote, "defaults.remote", tmpCfg.Default.Remote, "[DYNAMIC] Override defaults.remote")
	// 	rootCmd.PersistentFlags().StringVar(&cfg.Default.Revision, "defaults.revision", tmpCfg.Default.Revision, "[DYNAMIC] Override defaults.revision")
	// 	rootCmd.PersistentFlags().StringVar(&cfg.Default.Path, "defaults.path", tmpCfg.Default.Path, "[DYNAMIC] Override defaults.path")

	// 	for repoName, repoConfig := range tmpCfg.Repos {
	// 		rootCmd.PersistentFlags().StringVar(&repoConfig.Remote, fmt.Sprintf("repo.%s.remote", repoName), repoConfig.Remote, fmt.Sprintf("[DYNAMIC] Override repo.%s.remote", repoName))
	// 		rootCmd.PersistentFlags().StringVar(&repoConfig.Revision, fmt.Sprintf("repo.%s.revision", repoName), repoConfig.Revision, fmt.Sprintf("[DYNAMIC] Override repo.%s.revision", repoName))
	// 		rootCmd.PersistentFlags().StringVar(&repoConfig.Path, fmt.Sprintf("repo.%s.path", repoName), repoConfig.Path, fmt.Sprintf("[DYNAMIC] Override repo.%s.path", repoName))
	// 		cfg.Repos[repoName] = repoConfig
	// 	}

	// 	// Merge with defaults and apply CLI overrides
	// 	err = cfg.MergeWithDefaults()
	// 	if err != nil {
	// 		fmt.Printf("Error merging config: %v\n", err)
	// 		return
	// 	}

	// 	err = cfg.Validate()
	// 	if err != nil {
	// 		fmt.Printf("Error validating config: %v\n", err)
	// 		return
	// 	}
	// }
}

func initConfig() {
}

func Execute() error {
	return rootCmd.Execute()
}
