---
okf_version: "0.2"
type: Variable
title: cleanupCmd
resource: cmd/cleanup.go
tags:
  - "lang:go"
  - "type:Variable"
  - "module:cmd"
  - "domain:cleanup.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/cleanup/cleanupCmd
language: go
---

# cleanupCmd

## Signature

```go
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup repositories that are no longer defined in config",
	Long:  `Remove repositories that are no longer defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		// Process each repository
		for name, repo := range cfg.Repos {
			fmt.Printf("Cleaning repository %s...\n", name)

			err := git.CleanupRepository(repo, name)
			if err != nil {
				fmt.Printf("Error cleaning %s: %v\n", name, err)
				continue
			}
		}
	},
}
```

## Source
Lines 11–33 in `cmd/cleanup.go`

## Relationships

| Type | Target |
|------|--------|
| related | [cleanup](/cmd/cleanup.md) |
| calls | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| calls | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
| calls | [CleanupRepository](/git/git/CleanupRepository.md) |
