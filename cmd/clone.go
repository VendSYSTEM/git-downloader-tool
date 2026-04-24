package cmd

import (
	"fmt"
	"git-downloader-tool/config"
	"git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone repositories from config",
	Long:  `Clone repositories according to the configuration file.`,
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
			fmt.Printf("Cloning repository %s...\n", name)

			// Construct full repository URL
			fullURL := repo.Remote + repo.Repository

			// Clone the repository
			err := git.CloneRepository(fullURL, repo.Path+"/"+name)
			if err != nil {
				fmt.Printf("Error cloning %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Successfully cloned %s to %s\n", fullURL, repo.Path+"/"+name)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
