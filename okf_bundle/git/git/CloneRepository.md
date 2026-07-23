---
okf_version: "0.2"
type: Function
title: CloneRepository
description: CloneRepository clones a repository at the specified path. repo.Remote must be a resolved URL or base path.
resource: git/git.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git/CloneRepository
language: go
---

# CloneRepository

CloneRepository clones a repository at the specified path. repo.Remote must be a resolved URL or base path.

## Signature

```go
func CloneRepository(repo config.Repository, name string) error
```

## Docstring

CloneRepository clones a repository at the specified path. repo.Remote must be a resolved URL or base path.

## Source
Lines 13–35 in `git/git.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git](/git/git.md) |
| calls | [repositoryPath](/git/git/repositoryPath.md) |
| calls | [UpdateRepository](/git/git/UpdateRepository.md) |
| called_by | [cloneCmd](/cmd/clone/cloneCmd.md) |
| called_by | [TestCloneRepository_ReturnsErrorWhenCloneFails](/git/git_test/TestCloneRepository_ReturnsErrorWhenCloneFails.md) |
| called_by | [TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails](/git/git_test/TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails.md) |
| called_by | [TestCloneRepository_SucceedsWithLocalRepository](/git/git_test/TestCloneRepository_SucceedsWithLocalRepository.md) |
| called_by | [TestCloneRepository_SyncsWhenTargetExists](/git/git_test/TestCloneRepository_SyncsWhenTargetExists.md) |
| called_by | [TestUpdateRepository_SucceedsForValidGitRepository](/git/git_test/TestUpdateRepository_SucceedsForValidGitRepository.md) |
