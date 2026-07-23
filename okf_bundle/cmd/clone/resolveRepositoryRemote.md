---
okf_version: "0.2"
type: Function
title: resolveRepositoryRemote
resource: cmd/clone.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:clone.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/clone/resolveRepositoryRemote
language: go
---

# resolveRepositoryRemote

## Signature

```go
func resolveRepositoryRemote(cfg *config.Config, repo config.Repository) config.Repository
```

## Source
Lines 41–53 in `cmd/clone.go`

## Relationships

| Type | Target |
|------|--------|
| related | [clone](/cmd/clone.md) |
| called_by | [cloneCmd](/cmd/clone/cloneCmd.md) |
| called_by | [TestResolveRepositoryRemote_PreservesDirectRemoteURL](/cmd/runtime_config_additional_test/TestResolveRepositoryRemote_PreservesDirectRemoteURL.md) |
| called_by | [TestResolveRepositoryRemote_PreservesRepoWhenConfigNil](/cmd/runtime_config_additional_test/TestResolveRepositoryRemote_PreservesRepoWhenConfigNil.md) |
| called_by | [TestResolveRepositoryRemote_ReplacesNamedRemoteWithURL](/cmd/runtime_config_additional_test/TestResolveRepositoryRemote_ReplacesNamedRemoteWithURL.md) |
