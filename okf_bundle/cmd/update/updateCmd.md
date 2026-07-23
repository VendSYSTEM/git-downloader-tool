---
okf_version: "0.2"
type: Variable
title: updateCmd
resource: cmd/update.go
tags:
  - "lang:go"
  - "type:Variable"
  - "module:cmd"
  - "domain:update.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/update/updateCmd
language: go
---

# updateCmd

## Signature

```go
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update repositories from config",
	Long:  `Update existing repositories according to the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		// Process each repository
		for name, repo := range cfg.Repos {
			fmt.Printf("Updating repository %s...\n", name)

			// Update the repository
			err := git.UpdateRepository(repo, name)
			if err != nil {
				fmt.Printf("Error updating %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Successfully updated %s\n", name)
		}
	},
}
```

## Source
Lines 11–36 in `cmd/update.go`

## Relationships

| Type | Target |
|------|--------|
| related | [update](/cmd/update.md) |
| calls | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| calls | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
| calls | [UpdateRepository](/git/git/UpdateRepository.md) |
