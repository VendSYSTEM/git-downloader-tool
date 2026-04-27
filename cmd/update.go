package cmd

import (
	"fmt"

	"github.com/VendSYSTEM/git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update repositories from config",
	Long:  `Update existing repositories according to the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		// Process each repository
		for name, repo := range cfg.Repos {
			fmt.Printf("Updating repository %s...\n", name)

			// Update the repository
			err := git.UpdateRepository(repo, name)
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
