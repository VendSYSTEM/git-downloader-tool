package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/VendSYSTEM/git-downloader-tool/config"
)

func TestCloneRepository_SyncsWhenTargetExists(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)
	existingPath := filepath.Join(tempDir, "existing")
	runCommand(t, tempDir, "git", "clone", sourceRepo, existingPath)
	addCommittedFile(t, sourceRepo, "SYNCED.md", "synced\n")

	err := CloneRepository(config.Repository{
		Remote:   sourceRepo,
		Path:     tempDir,
		Revision: "main",
	}, "existing")
	if err != nil {
		t.Fatalf("expected existing target path to sync, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(existingPath, "SYNCED.md")); err != nil {
		t.Fatalf("expected existing checkout to receive synced file: %v", err)
	}
}

func TestCloneRepository_SucceedsWithLocalRepository(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)
	createBranchWithFile(t, sourceRepo, "feature", "FEATURE.md", "feature\n")
	clonePath := filepath.Join(tempDir, "cloned")

	err := CloneRepository(config.Repository{
		Remote:   sourceRepo,
		Path:     tempDir,
		Revision: "feature",
	}, "cloned")
	if err != nil {
		t.Fatalf("expected clone from local repository to succeed, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(clonePath, "README.md")); err != nil {
		t.Fatalf("expected cloned repository to contain committed file: %v", err)
	}
	if _, err := os.Stat(filepath.Join(clonePath, "FEATURE.md")); err != nil {
		t.Fatalf("expected cloned repository to switch to requested revision: %v", err)
	}
}

func TestCloneRepository_ReturnsErrorWhenCloneFails(t *testing.T) {
	tempDir := t.TempDir()
	err := CloneRepository(config.Repository{
		Remote:   filepath.Join(tempDir, "missing-source"),
		Path:     tempDir,
		Revision: "main",
	}, "dest")
	if err == nil {
		t.Fatalf("expected clone failure for missing source path, got nil")
	}

	if !strings.Contains(err.Error(), "failed to clone repository") {
		t.Fatalf("expected clone error prefix, got %q", err.Error())
	}
}

func TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)

	err := CloneRepository(config.Repository{
		Remote:   sourceRepo,
		Path:     tempDir,
		Revision: "missing-branch",
	}, "cloned")
	if err == nil {
		t.Fatalf("expected clone failure for missing revision, got nil")
	}

	if !strings.Contains(err.Error(), "failed to switch revision") {
		t.Fatalf("expected switch error prefix, got %q", err.Error())
	}
}

func TestUpdateRepository_SucceedsForValidGitRepository(t *testing.T) {
	tempDir := t.TempDir()
	sourceRepo := createCommittedRepo(t, tempDir)
	createBranchWithFile(t, sourceRepo, "feature", "FEATURE.md", "feature\n")

	if err := CloneRepository(config.Repository{
		Remote:   sourceRepo,
		Path:     tempDir,
		Revision: "main",
	}, "cloned"); err != nil {
		t.Fatalf("failed cloning local fixture repository: %v", err)
	}

	err := UpdateRepository(config.Repository{Path: tempDir, Revision: "feature"}, "cloned")
	if err != nil {
		t.Fatalf("expected git pull to succeed for valid clone, got %v", err)
	}

	if _, err := os.Stat(filepath.Join(tempDir, "cloned", "FEATURE.md")); err != nil {
		t.Fatalf("expected updated repository to switch to requested revision: %v", err)
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

	if !strings.Contains(err.Error(), "failed to switch revision") {
		t.Fatalf("expected switch error prefix, got %q", err.Error())
	}
}

func TestUpdateRepository_ReturnsErrorWhenPullFails(t *testing.T) {
	tempDir := t.TempDir()
	repoPath := createCommittedRepo(t, tempDir)

	err := UpdateRepository(config.Repository{Path: tempDir, Revision: "main"}, filepath.Base(repoPath))
	if err == nil {
		t.Fatalf("expected update failure when pull has no tracking branch, got nil")
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

	runCommand(t, parentDir, "git", "init", "-b", "main", repoPath)

	if err := os.WriteFile(filepath.Join(repoPath, "README.md"), []byte("hello\n"), 0o644); err != nil {
		t.Fatalf("failed writing repository file: %v", err)
	}

	runCommand(t, repoPath, "git", "add", ".")
	runCommand(t, repoPath, "git", "commit", "-m", "initial commit")

	return repoPath
}

func addCommittedFile(t *testing.T, repoPath, name, content string) {
	t.Helper()

	if err := os.WriteFile(filepath.Join(repoPath, name), []byte(content), 0o644); err != nil {
		t.Fatalf("failed writing repository file %s: %v", name, err)
	}

	runCommand(t, repoPath, "git", "add", ".")
	runCommand(t, repoPath, "git", "commit", "-m", "add "+name)
}

func createBranchWithFile(t *testing.T, repoPath, branch, name, content string) {
	t.Helper()

	runCommand(t, repoPath, "git", "switch", "-c", branch)
	addCommittedFile(t, repoPath, name, content)
	runCommand(t, repoPath, "git", "switch", "main")
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
