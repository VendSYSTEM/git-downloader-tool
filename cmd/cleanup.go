package cmd

import (
	"fmt"

	"github.com/VendSYSTEM/git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup repositories that are no longer defined in config",
	Long:  `Remove repositories that are no longer defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
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
