package cmd

import (
	"fmt"
	"git-downloader-tool/config"
	"git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update repositories from config",
	Long:  `Update existing repositories according to the configuration file.`,
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
			fmt.Printf("Updating repository %s...\n", name)

			// Update the repository
			err := git.UpdateRepository(repo.Path + "/" + name)
			if err != nil {
				fmt.Printf("Error updating %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Successfully updated %s\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
