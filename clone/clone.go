package clone

import (
    "fmt"
    "os"
    "os/exec"
)

func CloneRepository(remoteURL, repoName, revision, destPath string) error {
    // Create destination directory if it doesn't exist
    err := os.MkdirAll(destPath, 0755)
    if err != nil {
        return fmt.Errorf("failed to create directory %s: %v", destPath, err)
    }
    
    // Construct the full repository URL
    repoURL := fmt.Sprintf("%s/%s", remoteURL, repoName)
    
    // Build the git clone command
    cmd := exec.Command("git", "clone", "-b", revision, "--single-branch", "--depth", "1", repoURL, destPath)
    
    // Run the command
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to clone %s: %v\nOutput: %s", repoName, err, output)
    }
    
    return nil
}

func CloneRepositoryWithSubdir(remoteURL, repoName, revision, destPath, subPath string) error {
    // Create destination directory if it doesn't exist
    err := os.MkdirAll(destPath, 0755)
    if err != nil {
        return fmt.Errorf("failed to create directory %s: %v", destPath, err)
    }
    
    // Construct the full repository URL
    repoURL := fmt.Sprintf("%s/%s", remoteURL, repoName)
    
    // Build the git clone command with subdirectory
    cmd := exec.Command("git", "clone", "-b", revision, "--single-branch", "--depth", "1", repoURL, destPath)
    
    // Run the command
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to clone %s: %v\nOutput: %s", repoName, err, output)
    }
    
    return nil
}