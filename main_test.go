package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain_ExecutesInfoCommandWithoutFatal(t *testing.T) {
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

	originalArgs := os.Args
	os.Args = []string{"git-downloader-tool", "info", "--config", configPath}
	t.Cleanup(func() {
		os.Args = originalArgs
	})

	main()
}
