---
okf_version: "0.2"
type: Function
title: UpdateRepository
description: UpdateRepository updates an existing repository to the latest commit
resource: git/git.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git/UpdateRepository
language: go
---

# UpdateRepository

UpdateRepository updates an existing repository to the latest commit

## Signature

```go
func UpdateRepository(repo config.Repository, name string) error
```

## Docstring

UpdateRepository updates an existing repository to the latest commit

## Source
Lines 38–53 in `git/git.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git](/git/git.md) |
| calls | [repositoryPath](/git/git/repositoryPath.md) |
| called_by | [updateCmd](/cmd/update/updateCmd.md) |
| called_by | [CloneRepository](/git/git/CloneRepository.md) |
| called_by | [TestUpdateRepository_ReturnsErrorForNonRepositoryPath](/git/git_test/TestUpdateRepository_ReturnsErrorForNonRepositoryPath.md) |
| called_by | [TestUpdateRepository_ReturnsErrorWhenPullFails](/git/git_test/TestUpdateRepository_ReturnsErrorWhenPullFails.md) |
| called_by | [TestUpdateRepository_SucceedsForValidGitRepository](/git/git_test/TestUpdateRepository_SucceedsForValidGitRepository.md) |
