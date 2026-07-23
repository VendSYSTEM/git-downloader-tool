---
okf_version: "0.2"
type: Function
title: TestCloneRepository_SyncsWhenTargetExists
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/TestCloneRepository_SyncsWhenTargetExists
language: go
---

# TestCloneRepository_SyncsWhenTargetExists

## Signature

```go
func TestCloneRepository_SyncsWhenTargetExists(t *testing.T)
```

## Source
Lines 13–32 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [createCommittedRepo](/git/git_test/createCommittedRepo.md) |
| calls | [runCommand](/git/git_test/runCommand.md) |
| calls | [addCommittedFile](/git/git_test/addCommittedFile.md) |
| calls | [CloneRepository](/git/git/CloneRepository.md) |
