package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"git-downloader-tool/config"

	"github.com/spf13/cobra"
)

func TestInfoCommandUsesEffectiveConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	configContents := []byte(`remote:
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

	if err := os.WriteFile(configPath, configContents, 0o644); err != nil {
		t.Fatalf("failed writing test config: %v", err)
	}

	previousCfgFile := cfgFile
	cfgFile = configPath
	t.Cleanup(func() {
		cfgFile = previousCfgFile
	})

	testCmd := &cobra.Command{Use: "info"}
	loadedConfig, err := config.LoadConfig(configPath)
	if err != nil {
		t.Fatalf("failed loading config fixture: %v", err)
	}

	registerDynamicOverrideFlags(testCmd, loadedConfig)
	if err := testCmd.PersistentFlags().Set("defaults.revision", "release"); err != nil {
		t.Fatalf("setting changed defaults flag failed: %v", err)
	}

	output := captureStdout(t, func() {
		infoCmd.Run(testCmd, nil)
	})

	if !strings.Contains(output, "Defaults: remote=origin, revision=release, path=repos") {
		t.Fatalf("expected info output to include overridden defaults.revision, got:\n%s", output)
	}

	if !strings.Contains(output, "- tool: remote=https://example.com/, revision=release, path=repos") {
		t.Fatalf("expected info output to include merged repo revision from override, got:\n%s", output)
	}
}

func captureStdout(t *testing.T, run func()) string {
	t.Helper()

	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed creating stdout pipe: %v", err)
	}

	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()
	defer func() {
		_ = r.Close()
	}()
	defer func() {
		_ = w.Close()
	}()

	run()

	if err := w.Close(); err != nil {
		t.Fatalf("failed closing writer end of stdout pipe: %v", err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed reading captured stdout: %v", err)
	}

	return buf.String()
}
