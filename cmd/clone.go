package cmd

import (
	"fmt"

	"github.com/VendSYSTEM/git-downloader-tool/git"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone repositories from config",
	Long:  `Clone repositories according to the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
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
