---
okf_version: "0.2"
type: Function
title: repositoryPath
resource: git/git.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git/repositoryPath
language: go
---

# repositoryPath

## Signature

```go
func repositoryPath(repo config.Repository, name string) string
```

## Source
Lines 72–74 in `git/git.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git](/git/git.md) |
| called_by | [CleanupRepository](/git/git/CleanupRepository.md) |
| called_by | [CloneRepository](/git/git/CloneRepository.md) |
| called_by | [UpdateRepository](/git/git/UpdateRepository.md) |
