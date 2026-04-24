package cmd

import (
	"fmt"
	"git-downloader-tool/config"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about repositories",
	Long:  `Show details about the repositories defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		cfg, err := config.LoadConfig(cfgFile)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		// Merge with defaults
		err = cfg.MergeWithDefaults()
		if err != nil {
			fmt.Printf("Error merging config: %v\n", err)
			return
		}

		// Validate config
		err = cfg.Validate()
		if err != nil {
			fmt.Printf("Error validating config: %v\n", err)
			return
		}

		fmt.Printf("Config loaded from %s\n", cfgFile)
		fmt.Printf("Defaults: remote=%s, revision=%s, path=%s\n", cfg.Default.Remote, cfg.Default.Revision, cfg.Default.Path)
		fmt.Printf("Repos: %d found\n", len(cfg.Repos))
		for name, repo := range cfg.Repos {
			fmt.Printf("- %s: remote=%s, revision=%s, path=%s\n", name, repo.Remote, repo.Revision, repo.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
