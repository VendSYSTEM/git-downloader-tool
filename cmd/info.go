package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about repositories",
	Long:  `Show details about the repositories defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		fmt.Printf("Config loaded from %s\n", cfgFile)
		fmt.Printf("Defaults: remote=%s, revision=%s, path=%s\n", cfg.Defaults.Remote, cfg.Defaults.Revision, cfg.Defaults.Path)
		fmt.Printf("Repos: %d found\n", len(cfg.Repos))
		for name, repo := range cfg.Repos {
			fmt.Printf("- %s: remote=%s, revision=%s, path=%s\n", name, repo.Remote, repo.Revision, repo.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
