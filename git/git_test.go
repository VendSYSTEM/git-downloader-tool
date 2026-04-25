package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCloneRepository_ReturnsNilWhenTargetExists(t *testing.T) {
	existingPath := filepath.Join(t.TempDir(), "existing")
	if err := os.MkdirAll(existingPath, 0o755); err != nil {
		t.Fatalf("failed creating existing target path: %v", err)
	}

	err := CloneRepository("https://example.com/does-not-matter.git", existingPath)
	if err != nil {
		t.Fatalf("expected existing target path to short-circuit clone, got %v", err)
	}
}

func TestCloneRepository_SucceedsWithLocalRepository(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)
	clonePath := filepath.Join(tempDir, "cloned")

	err := CloneRepository(sourceRepo, clonePath)
	if err != nil {
		t.Fatalf("expected clone from local repository to succeed, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(clonePath, "README.md")); err != nil {
		t.Fatalf("expected cloned repository to contain committed file: %v", err)
	}
}

func TestCloneRepository_ReturnsErrorWhenCloneFails(t *testing.T) {
	tempDir := t.TempDir()
	err := CloneRepository(filepath.Join(tempDir, "missing-source"), filepath.Join(tempDir, "dest"))
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
	clonePath := filepath.Join(tempDir, "cloned")

	if err := CloneRepository(sourceRepo, clonePath); err != nil {
		t.Fatalf("failed cloning local fixture repository: %v", err)
	}

	err := UpdateRepository(clonePath)
	if err != nil {
		t.Fatalf("expected git pull to succeed for valid clone, got %v", err)
	}
}

func TestUpdateRepository_ReturnsErrorForNonRepositoryPath(t *testing.T) {
	nonRepoPath := filepath.Join(t.TempDir(), "not-a-repo")
	if err := os.MkdirAll(nonRepoPath, 0o755); err != nil {
		t.Fatalf("failed creating non-repo path: %v", err)
	}

	err := UpdateRepository(nonRepoPath)
	if err == nil {
		t.Fatalf("expected update failure for non-repository path, got nil")
	}

	if !strings.Contains(err.Error(), "failed to update repository") {
		t.Fatalf("expected update error prefix, got %q", err.Error())
	}
}

func TestCleanupRepository_ReturnsNilWhenPathDoesNotExist(t *testing.T) {
	err := CleanupRepository(filepath.Join(t.TempDir(), "missing"))
	if err != nil {
		t.Fatalf("expected cleanup of missing path to be a no-op, got %v", err)
	}
}

func TestCleanupRepository_RemovesExistingDirectory(t *testing.T) {
	targetPath := filepath.Join(t.TempDir(), "to-delete")
	if err := os.MkdirAll(targetPath, 0o755); err != nil {
		t.Fatalf("failed creating target path: %v", err)
	}

	if err := os.WriteFile(filepath.Join(targetPath, "artifact.txt"), []byte("content"), 0o644); err != nil {
		t.Fatalf("failed writing artifact file: %v", err)
	}

	err := CleanupRepository(targetPath)
	if err != nil {
		t.Fatalf("expected cleanup to remove existing directory, got %v", err)
	}

	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		t.Fatalf("expected cleanup target to be removed, stat err: %v", err)
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
