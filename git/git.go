package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/VendSYSTEM/git-downloader-tool/config"
)

// CloneRepository clones a repository at the specified path. repo.Remote must be a resolved URL or base path.
func CloneRepository(repo config.Repository, name string) error {
	url := repo.Remote + repo.Repository
	path := repositoryPath(repo, name)

	// Check if the target path already exists
	if _, err := exec.Command("test", "-d", path).CombinedOutput(); err == nil {
		fmt.Printf("Target path '%s' already exists, syncing...\n", path)
		return UpdateRepository(repo, name)
	}
	cmd := exec.Command("git", "clone", url, path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repository: %s - %v", strings.TrimSpace(string(output)), err)
	}

	cmd = exec.Command("git", "-C", path, "switch", repo.Revision)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to switch revision: %s - %v", strings.TrimSpace(string(output)), err)
	}

	return nil
}

// UpdateRepository updates an existing repository to the latest commit
func UpdateRepository(repo config.Repository, name string) error {
	path := repositoryPath(repo, name)

	cmd := exec.Command("git", "-C", path, "switch", repo.Revision)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to switch revision: %s - %v", strings.TrimSpace(string(output)), err)
	}

	cmd = exec.Command("git", "-C", path, "pull")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to update repository: %s - %v", strings.TrimSpace(string(output)), err)
	}
	return nil
}

// CleanupRepository removes a repository directory
func CleanupRepository(repo config.Repository, name string) error {
	path := repositoryPath(repo, name)

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

func repositoryPath(repo config.Repository, name string) string {
	return filepath.Join(repo.Path, name)
}
