package cmd

import (
    "fmt"
    "log"

    "github.com/spf13/cobra"
    "git-downloader-tool/config"
    "git-downloader-tool/output"
)

var (
    configFile string
    remote     string
    revision   string
    path       string
)

var rootCmd = &cobra.Command{
    Use:   "git-downloader",
    Short: "Git Downloader Configuration Tool",
    Long:  "A tool to consume git-downloader.yaml configuration",
    Run: func(cmd *cobra.Command, args []string) {
        cfg, err := config.ParseConfig(configFile)
        if err != nil {
            log.Fatalf("Error parsing config: %v", err)
        }
        
        // Apply CLI overrides
        if remote != "" {
            cfg.Defaults.Remote = remote
        }
        if revision != "" {
            cfg.Defaults.Revision = revision
        }
        if path != "" {
            cfg.Defaults.Path = path
        }
        
        cfg.ApplyDefaults()
        
        fmt.Println(output.FormatRemotes(cfg.Remotes))
        fmt.Println(output.FormatDefaults(cfg.Defaults))
        fmt.Println(output.FormatRepos(cfg.Repos))
    },
}

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    rootCmd.Flags().StringVar(&configFile, "config", "git-downloader.yaml", "Path to config file")
    rootCmd.Flags().StringVar(&remote, "remote", "", "Override remote")
    rootCmd.Flags().StringVar(&revision, "revision", "", "Override revision")
    rootCmd.Flags().StringVar(&path, "path", "", "Override path")
}