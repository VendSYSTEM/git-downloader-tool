package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/VendSYSTEM/git-downloader-tool/config"
)

func TestCloneRepository_ReturnsNilWhenTargetExists(t *testing.T) {
	tempDir := t.TempDir()
	existingPath := filepath.Join(tempDir, "existing")
	if err := os.MkdirAll(existingPath, 0o755); err != nil {
		t.Fatalf("failed creating existing target path: %v", err)
	}

	err := CloneRepository(config.Repository{
		Remote:     "https://example.com/",
		Repository: "does-not-matter.git",
		Path:       tempDir,
	}, "existing")
	if err != nil {
		t.Fatalf("expected existing target path to short-circuit clone, got %v", err)
	}
}

func TestCloneRepository_SucceedsWithLocalRepository(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)
	clonePath := filepath.Join(tempDir, "cloned")

	err := CloneRepository(config.Repository{
		Remote: sourceRepo,
		Path:   tempDir,
	}, "cloned")
	if err != nil {
		t.Fatalf("expected clone from local repository to succeed, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(clonePath, "README.md")); err != nil {
		t.Fatalf("expected cloned repository to contain committed file: %v", err)
	}
}

func TestCloneRepository_ReturnsErrorWhenCloneFails(t *testing.T) {
	tempDir := t.TempDir()
	err := CloneRepository(config.Repository{
		Remote: filepath.Join(tempDir, "missing-source"),
		Path:   tempDir,
	}, "dest")
	if err == nil {
		t.Fatalf("expected clone failure for missing source path, got nil")
	}

	if !strings.Contains(err.Error(), "failed to clone repository") {
		t.Fatalf("expected clone error prefix, got %q", err.Error())
	}
}

func TestUpdateRepository_SucceedsForValidGitRepository(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)

	if err := CloneRepository(config.Repository{
		Remote: sourceRepo,
		Path:   tempDir,
	}, "cloned"); err != nil {
		t.Fatalf("failed cloning local fixture repository: %v", err)
	}

	err := UpdateRepository(config.Repository{Path: tempDir}, "cloned")
	if err != nil {
		t.Fatalf("expected git pull to succeed for valid clone, got %v", err)
	}
}

func TestUpdateRepository_ReturnsErrorForNonRepositoryPath(t *testing.T) {
	nonRepoPath := filepath.Join(t.TempDir(), "not-a-repo")
	if err := os.MkdirAll(nonRepoPath, 0o755); err != nil {
		t.Fatalf("failed creating non-repo path: %v", err)
	}

	err := UpdateRepository(config.Repository{Path: filepath.Dir(nonRepoPath)}, filepath.Base(nonRepoPath))
	if err == nil {
		t.Fatalf("expected update failure for non-repository path, got nil")
	}

	if !strings.Contains(err.Error(), "failed to update repository") {
		t.Fatalf("expected update error prefix, got %q", err.Error())
	}
}

func TestCleanupRepository_ReturnsNilWhenPathDoesNotExist(t *testing.T) {
	err := CleanupRepository(config.Repository{Path: t.TempDir()}, "missing")
	if err != nil {
		t.Fatalf("expected cleanup of missing path to be a no-op, got %v", err)
	}
}

func TestCleanupRepository_RemovesExistingDirectory(t *testing.T) {
	tempDir := t.TempDir()
	targetPath := filepath.Join(tempDir, "to-delete")
	if err := os.MkdirAll(targetPath, 0o755); err != nil {
		t.Fatalf("failed creating target path: %v", err)
	}

	if err := os.WriteFile(filepath.Join(targetPath, "artifact.txt"), []byte("content"), 0o644); err != nil {
		t.Fatalf("failed writing artifact file: %v", err)
	}

	err := CleanupRepository(config.Repository{Path: tempDir}, "to-delete")
	if err != nil {
		t.Fatalf("expected cleanup to remove existing directory, got %v", err)
	}

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		t.Fatalf("expected cleanup target to be removed, stat err: %v", err)
	}
}

func TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory(t *testing.T) {
	tempDir := t.TempDir()
	previousDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed getting current working directory: %v", err)
	}
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed changing working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(previousDir); err != nil {
			t.Fatalf("failed restoring working directory: %v", err)
		}
	})

	targetPath := filepath.Join(tempDir, "to-delete")
	if err := os.MkdirAll(targetPath, 0o755); err != nil {
		t.Fatalf("failed creating target path: %v", err)
	}

	err = CleanupRepository(config.Repository{Path: ""}, "to-delete")
	if err != nil {
		t.Fatalf("expected cleanup to remove relative target, got %v", err)
	}

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		t.Fatalf("expected relative cleanup target to be removed, stat err: %v", err)
	}
}

func createCommittedRepo(t *testing.T, parentDir string) string {
	t.Helper()

	repoPath := filepath.Join(parentDir, "source-repo")
	if err := os.MkdirAll(repoPath, 0o755); err != nil {
		t.Fatalf("failed creating source repository directory: %v", err)
	}

	runCommand(t, parentDir, "git", "init", repoPath)

	if err := os.WriteFile(filepath.Join(repoPath, "README.md"), []byte("hello\n"), 0o644); err != nil {
		t.Fatalf("failed writing repository file: %v", err)
	}

	runCommand(t, repoPath, "git", "add", ".")
	runCommand(t, repoPath, "git", "commit", "-m", "initial commit")

	return repoPath
}

func runCommand(t *testing.T, dir, bin string, args ...string) {
	t.Helper()

	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Env = append(
		os.Environ(),
		"GIT_AUTHOR_NAME=Test User",
		"GIT_AUTHOR_EMAIL=test@example.com",
		"GIT_COMMITTER_NAME=Test User",
		"GIT_COMMITTER_EMAIL=test@example.com",
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("command %s %s failed: %v\n%s", bin, strings.Join(args, " "), err, string(output))
	}
}
