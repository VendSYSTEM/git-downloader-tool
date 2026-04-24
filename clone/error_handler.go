package clone

import (
    "errors"
    "fmt"
    "os/exec"
    "git-downloader-tool/config"
)

func HandleCloneError(repoName string, err error) error {
    if err != nil {
        return fmt.Errorf("failed to clone repository %s: %v", repoName, err)
    }
    return nil
}

// Check if git command is available
func IsGitAvailable() bool {
    _, err := exec.LookPath("git")
    return err == nil
}

func ValidateCloneOptions(cfg *config.Config, options *CloneOptions) error {
    if !IsGitAvailable() {
        return errors.New("git command not found in PATH")
    }
    
    if cfg == nil {
        return errors.New("config cannot be nil")
    }
    
    return nil
}