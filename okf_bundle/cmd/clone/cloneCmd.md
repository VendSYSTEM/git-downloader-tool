---
okf_version: "0.2"
type: Variable
title: cloneCmd
resource: cmd/clone.go
tags:
  - "lang:go"
  - "type:Variable"
  - "module:cmd"
  - "domain:clone.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/clone/cloneCmd
language: go
---

# cloneCmd

## Signature

```go
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone repositories from config",
	Long:  `Clone repositories according to the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := buildEffectiveConfig(cmd)
		if err != nil {
			fmt.Printf("%s\n", runtimeConfigErrorMessage(err))
			return
		}

		// Process each repository
		for name, repo := range cfg.Repos {
			fmt.Printf("Cloning repository %s...\n", name)
			resolvedRepo := resolveRepositoryRemote(cfg, repo)

			// Clone the repository
			err := git.CloneRepository(resolvedRepo, name)
			if err != nil {
				fmt.Printf("Error cloning %s: %v\n", name, err)
				continue
			}

			fmt.Printf("Successfully cloned %s to %s\n", resolvedRepo.Remote+resolvedRepo.Repository, repositoryClonePath(repo, name))
		}
	},
}
```

## Source
Lines 13–39 in `cmd/clone.go`

## Relationships

| Type | Target |
|------|--------|
| related | [clone](/cmd/clone.md) |
| calls | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| calls | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
| calls | [resolveRepositoryRemote](/cmd/clone/resolveRepositoryRemote.md) |
| calls | [CloneRepository](/git/git/CloneRepository.md) |
| calls | [repositoryClonePath](/cmd/clone/repositoryClonePath.md) |
