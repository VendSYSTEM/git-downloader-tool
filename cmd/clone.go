package cmd

import (
    "log"

    "github.com/spf13/cobra"
    "git-downloader-tool/config"
    "git-downloader-tool/clone"
)

var (
    cloneConfigFile string
    cloneRevision   string
    dryRun          bool
)

var cloneCmd = &cobra.Command{
    Use:   "clone",
    Short: "Clone repositories based on config",
    Long:  "Clone all repositories defined in the git-downloader.yaml configuration",
    Run: func(cmd *cobra.Command, args []string) {
        // Parse the config file
        cfg, err := config.ParseConfig(cloneConfigFile)
        if err != nil {
            log.Fatalf("Error parsing config: %v", err)
        }
        
        // Apply any CLI overrides to defaults
        if cloneRevision != "" {
            cfg.Defaults.Revision = cloneRevision
        }
        
        // Create clone options
        options := &clone.CloneOptions{
            ConfigFile: cloneConfigFile,
            Revision:   cloneRevision,
            DryRun:     dryRun,
        }
        
        // Process and clone repos
        err = clone.ProcessConfigAndClone(cfg, options)
        if err != nil {
            log.Fatalf("Error during cloning: %v", err)
        }
    },
}

func init() {
    rootCmd.AddCommand(cloneCmd)
    
    // Add flags to clone command
    cloneCmd.Flags().StringVar(&cloneConfigFile, "config", "git-downloader.yaml", "Path to config file")
    cloneCmd.Flags().StringVar(&cloneRevision, "revision", "", "Override revision for cloning")
    cloneCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Show what would be cloned without actually cloning")
}