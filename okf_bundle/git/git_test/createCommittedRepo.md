---
okf_version: "0.2"
type: Function
title: createCommittedRepo
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/createCommittedRepo
language: go
---

# createCommittedRepo

## Signature

```go
func createCommittedRepo(t *testing.T, parentDir string) string
```

## Source
Lines 202–220 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [runCommand](/git/git_test/runCommand.md) |
| called_by | [TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails](/git/git_test/TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails.md) |
| called_by | [TestCloneRepository_SucceedsWithLocalRepository](/git/git_test/TestCloneRepository_SucceedsWithLocalRepository.md) |
| called_by | [TestCloneRepository_SyncsWhenTargetExists](/git/git_test/TestCloneRepository_SyncsWhenTargetExists.md) |
| called_by | [TestUpdateRepository_ReturnsErrorWhenPullFails](/git/git_test/TestUpdateRepository_ReturnsErrorWhenPullFails.md) |
| called_by | [TestUpdateRepository_SucceedsForValidGitRepository](/git/git_test/TestUpdateRepository_SucceedsForValidGitRepository.md) |
