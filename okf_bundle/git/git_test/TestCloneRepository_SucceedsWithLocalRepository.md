---
okf_version: "0.2"
type: Function
title: TestCloneRepository_SucceedsWithLocalRepository
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/TestCloneRepository_SucceedsWithLocalRepository
language: go
---

# TestCloneRepository_SucceedsWithLocalRepository

## Signature

```go
func TestCloneRepository_SucceedsWithLocalRepository(t *testing.T)
```

## Source
Lines 34–55 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [createCommittedRepo](/git/git_test/createCommittedRepo.md) |
| calls | [createBranchWithFile](/git/git_test/createBranchWithFile.md) |
| calls | [CloneRepository](/git/git/CloneRepository.md) |
