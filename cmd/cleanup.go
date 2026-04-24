package cmd

import (
	"fmt"
	"git-downloader-tool/config"
	"git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup repositories that are no longer defined in config",
	Long:  `Remove repositories that are no longer defined in the configuration file.`,
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

		// Process each repository
		for name, repo := range cfg.Repos {
			fmt.Printf("Cleaning repository %s...\n", name)

			err := git.CleanupRepository(repo.Path + "/" + name)
			if err != nil {
				fmt.Printf("Error cleaning %s: %v\n", name, err)
				continue
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
