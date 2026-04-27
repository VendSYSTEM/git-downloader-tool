package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/VendSYSTEM/git-downloader-tool/config"
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
			resolvedRepo := resolveRepositoryRemote(cfg, repo)

			// Clone the repository
			err := git.CloneRepository(resolvedRepo, name)
			if err != nil {
				fmt.Printf("Error cloning %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Successfully cloned %s to %s\n", resolvedRepo.Remote+resolvedRepo.Repository, repositoryClonePath(repo, name))
		}
	},
}

func resolveRepositoryRemote(cfg *config.Config, repo config.Repository) config.Repository {
	if cfg == nil {
		return repo
	}

	remote, ok := cfg.Remotes[repo.Remote]
	if !ok {
		return repo
	}

	repo.Remote = remote.URL
	return repo
}

func repositoryClonePath(repo config.Repository, name string) string {
	return filepath.Join(repo.Path, name)
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
