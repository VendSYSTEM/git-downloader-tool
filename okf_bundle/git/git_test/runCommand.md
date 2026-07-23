---
okf_version: "0.2"
type: Function
title: runCommand
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/runCommand
language: go
---

# runCommand

## Signature

```go
func runCommand(t *testing.T, dir, bin string, args ...string)
```

## Source
Lines 241–257 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| called_by | [TestCloneRepository_SyncsWhenTargetExists](/git/git_test/TestCloneRepository_SyncsWhenTargetExists.md) |
| called_by | [addCommittedFile](/git/git_test/addCommittedFile.md) |
| called_by | [createBranchWithFile](/git/git_test/createBranchWithFile.md) |
| called_by | [createCommittedRepo](/git/git_test/createCommittedRepo.md) |
