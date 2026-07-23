---
okf_version: "0.2"
type: Function
title: repositoryClonePath
resource: cmd/clone.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:clone.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/clone/repositoryClonePath
language: go
---

# repositoryClonePath

## Signature

```go
func repositoryClonePath(repo config.Repository, name string) string
```

## Source
Lines 55–57 in `cmd/clone.go`

## Relationships

| Type | Target |
|------|--------|
| related | [clone](/cmd/clone.md) |
| called_by | [cloneCmd](/cmd/clone/cloneCmd.md) |
| called_by | [TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName](/cmd/runtime_config_additional_test/TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName.md) |
