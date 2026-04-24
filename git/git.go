package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// CloneRepository clones a repository at the specified path
func CloneRepository(url, path string) error {
	// Check if the target path already exists
	if _, err := exec.Command("test", "-d", path).CombinedOutput(); err == nil {
		fmt.Printf("Target path '%s' already exists...\n", path)
		return nil
	}
	cmd := exec.Command("git", "clone", url, path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repository: %s - %v", strings.TrimSpace(string(output)), err)
	}
	return nil
}

// UpdateRepository updates an existing repository to the latest commit
func UpdateRepository(path string) error {
	cmd := exec.Command("git", "-C", path, "pull")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update repository: %s - %v", strings.TrimSpace(string(output)), err)
	}
	return nil
}

// CleanupRepository removes a repository directory
func CleanupRepository(path string) error {
	// Check if the target path exists
	if _, err := exec.Command("test", "-d", path).CombinedOutput(); err != nil {
		fmt.Printf("Target path '%s' does not exist...\n", path)
		return nil
	}
	cmd := exec.Command("rm", "-rf", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to cleanup repository: %s - %v", strings.TrimSpace(string(output)), err)
	}
	return nil
}
