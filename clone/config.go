package clone

import (
    "fmt"
    "path/filepath"
    "git-downloader-tool/config"
)

func ProcessConfigAndClone(cfg *config.Config, options *CloneOptions) error {
    // For each repo in the config, clone it
    for _, repo := range cfg.Repos {
        // Determine remote URL based on repo's remote setting
        var remoteURL string
        if repo.Remote != nil {
            if remote, exists := cfg.Remotes[*repo.Remote]; exists {
                remoteURL = remote.URL
            } else {
                return fmt.Errorf("remote %s not found", *repo.Remote)
            }
        } else {
            // Use the default remote from config
            if remote, exists := cfg.Remotes[cfg.Defaults.Remote]; exists {
                remoteURL = remote.URL
            } else {
                return fmt.Errorf("default remote %s not found", cfg.Defaults.Remote)
            }
        }
        
        // Determine revision to use
        var revision string
        if repo.Revision != nil {
            revision = *repo.Revision
        } else {
            revision = cfg.Defaults.Revision
        }
        
        // Determine destination path
        var destPath string
        if repo.Path != nil {
            destPath = *repo.Path
        } else {
            destPath = cfg.Defaults.Path
        }
        
        // Create the full destination path
        fullPath := filepath.Join(destPath, repo.Name)
        
        if options.DryRun {
            fmt.Printf("Would clone %s to %s\n", repo.Name, fullPath)
        } else {
            fmt.Printf("Cloning %s to %s\n", repo.Name, fullPath)
            err := CloneRepository(remoteURL, repo.Name, revision, fullPath)
            if err != nil {
                return fmt.Errorf("failed to clone %s: %v", repo.Name, err)
            }
        }
    }
    
    return nil
}