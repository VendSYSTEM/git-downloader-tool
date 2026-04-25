package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfig_Success(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "config.yaml")

	contents := []byte(`remotes:
  origin:
    url: https://example.com/
defaults:
  remote: origin
  revision: main
  path: repos
repos:
  tool:
    repository: team/tool.git
`)

	if err := os.WriteFile(configPath, contents, 0o644); err != nil {
		t.Fatalf("failed writing config fixture: %v", err)
	}

	loaded, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("expected LoadConfig to succeed, got: %v", err)
	}

	if got := loaded.Remotes["origin"].URL; got != "https://example.com/" {
		t.Fatalf("expected remotes.origin.url to be parsed, got %q", got)
	}

	if got := loaded.Defaults.Remote; got != "origin" {
		t.Fatalf("expected defaults.remote to be parsed, got %q", got)
	}

	if got := loaded.Repos["tool"].Repository; got != "team/tool.git" {
		t.Fatalf("expected repos.tool.repository to be parsed, got %q", got)
	}
}

func TestLoadConfig_ReturnsReadErrorForMissingFile(t *testing.T) {
	_, err := LoadConfig(filepath.Join(t.TempDir(), "missing.yaml"))
	if err == nil {
		t.Fatalf("expected missing file read error, got nil")
	}
}

func TestLoadConfig_ReturnsYAMLErrorForInvalidData(t *testing.T) {
	configPath := filepath.Join(t.TempDir(), "invalid.yaml")
	if err := os.WriteFile(configPath, []byte("defaults: ["), 0o644); err != nil {
		t.Fatalf("failed writing invalid yaml fixture: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Fatalf("expected yaml parse error, got nil")
	}
}

func TestMergeWithDefaults_FillsMissingRepoFields(t *testing.T) {
	cfg := &Config{
		Remotes: map[string]Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]Repository{
			"tool": {
				Repository: "team/tool.git",
			},
		},
	}

	if err := cfg.MergeWithDefaults(); err != nil {
		t.Fatalf("expected MergeWithDefaults to succeed, got %v", err)
	}

	repo := cfg.Repos["tool"]
	if repo.Remote != "https://example.com/" {
		t.Fatalf("expected repo remote to inherit remote URL, got %q", repo.Remote)
	}
	if repo.Revision != "main" {
		t.Fatalf("expected repo revision to inherit defaults.revision, got %q", repo.Revision)
	}
	if repo.Path != "repos" {
		t.Fatalf("expected repo path to inherit defaults.path, got %q", repo.Path)
	}
}

func TestMergeWithDefaults_PreservesExplicitRepoFields(t *testing.T) {
	cfg := &Config{
		Remotes: map[string]Remote{
			"origin": {URL: "https://example.com/"},
		},
		Defaults: Defaults{
			Remote:   "origin",
			Revision: "main",
			Path:     "repos",
		},
		Repos: map[string]Repository{
			"tool": {
				Remote:     "ssh://git@example.com/",
				Repository: "team/tool.git",
				Revision:   "release",
				Path:       "workspace/tool",
			},
		},
	}

	if err := cfg.MergeWithDefaults(); err != nil {
		t.Fatalf("expected MergeWithDefaults to succeed, got %v", err)
	}

	repo := cfg.Repos["tool"]
	if repo.Remote != "ssh://git@example.com/" {
		t.Fatalf("expected explicit remote value to remain unchanged, got %q", repo.Remote)
	}
	if repo.Revision != "release" {
		t.Fatalf("expected explicit revision to remain unchanged, got %q", repo.Revision)
	}
	if repo.Path != "workspace/tool" {
		t.Fatalf("expected explicit path to remain unchanged, got %q", repo.Path)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		cfg         *Config
		expectError string
	}{
		{
			name: "requires defaults.remote",
			cfg: &Config{
				Remotes:  map[string]Remote{"origin": {URL: "https://example.com/"}},
				Defaults: Defaults{},
				Repos:    map[string]Repository{"tool": {Repository: "team/tool.git"}},
			},
			expectError: "defaults.remote is required",
		},
		{
			name: "requires defaults.remote to point to known remote",
			cfg: &Config{
				Remotes:  map[string]Remote{"origin": {URL: "https://example.com/"}},
				Defaults: Defaults{Remote: "missing"},
				Repos:    map[string]Repository{"tool": {Repository: "team/tool.git"}},
			},
			expectError: "default remote 'missing' is not defined",
		},
		{
			name: "requires non-empty repository name",
			cfg: &Config{
				Remotes:  map[string]Remote{"origin": {URL: "https://example.com/"}},
				Defaults: Defaults{Remote: "origin"},
				Repos:    map[string]Repository{"tool": {Repository: ""}},
			},
			expectError: "repository name is required",
		},
		{
			name: "valid config passes",
			cfg: &Config{
				Remotes: map[string]Remote{
					"origin": {URL: "https://example.com/"},
				},
				Defaults: Defaults{Remote: "origin", Revision: "main", Path: "repos"},
				Repos: map[string]Repository{
					"tool": {Repository: "team/tool.git"},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cfg.Validate()
			if tc.expectError == "" {
				if err != nil {
					t.Fatalf("expected Validate to succeed, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("expected Validate to fail with %q, got nil", tc.expectError)
			}

			if !strings.Contains(err.Error(), tc.expectError) {
				t.Fatalf("expected error containing %q, got %q", tc.expectError, err.Error())
			}
		})
	}
}
