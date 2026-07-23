---
okf_version: "0.2"
type: Variable
title: infoCmd
resource: cmd/info.go
tags:
  - "lang:go"
  - "type:Variable"
  - "module:cmd"
  - "domain:info.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T12:20:06Z"
concept_id: cmd/info/infoCmd
language: go
---

# infoCmd

## Signature

```go
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about repositories",
	Long:  `Show details about the repositories defined in the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		fmt.Printf("Config loaded from %s\n", cfgFile)
		fmt.Printf("Defaults: remote=%s, revision=%s, path=%s\n", cfg.Defaults.Remote, cfg.Defaults.Revision, cfg.Defaults.Path)
		fmt.Printf("Repos: %d found\n", len(cfg.Repos))
		for name, repo := range cfg.Repos {
			fmt.Printf("- %s: remote=%s, revision=%s, path=%s\n", name, repo.Remote, repo.Revision, repo.Path)
		}
	},
}
```

## Source
Lines 9–27 in `cmd/info.go`

## Relationships

| Type | Target |
|------|--------|
| related | [info](/cmd/info.md) |
| calls | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| calls | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
