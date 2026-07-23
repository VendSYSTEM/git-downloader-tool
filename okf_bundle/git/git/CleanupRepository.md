---
okf_version: "0.2"
type: Function
title: CleanupRepository
description: CleanupRepository removes a repository directory
resource: git/git.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git/CleanupRepository
language: go
---

# CleanupRepository

CleanupRepository removes a repository directory

## Signature

```go
func CleanupRepository(repo config.Repository, name string) error
```

## Docstring

CleanupRepository removes a repository directory

## Source
Lines 56–70 in `git/git.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git](/git/git.md) |
| calls | [repositoryPath](/git/git/repositoryPath.md) |
| called_by | [cleanupCmd](/cmd/cleanup/cleanupCmd.md) |
| called_by | [TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory](/git/git_test/TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory.md) |
| called_by | [TestCleanupRepository_RemovesExistingDirectory](/git/git_test/TestCleanupRepository_RemovesExistingDirectory.md) |
| called_by | [TestCleanupRepository_ReturnsNilWhenPathDoesNotExist](/git/git_test/TestCleanupRepository_ReturnsNilWhenPathDoesNotExist.md) |
