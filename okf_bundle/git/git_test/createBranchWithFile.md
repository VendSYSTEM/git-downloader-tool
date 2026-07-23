---
okf_version: "0.2"
type: Function
title: createBranchWithFile
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/createBranchWithFile
language: go
---

# createBranchWithFile

## Signature

```go
func createBranchWithFile(t *testing.T, repoPath, branch, name, content string)
```

## Source
Lines 233–239 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [runCommand](/git/git_test/runCommand.md) |
| calls | [addCommittedFile](/git/git_test/addCommittedFile.md) |
| called_by | [TestCloneRepository_SucceedsWithLocalRepository](/git/git_test/TestCloneRepository_SucceedsWithLocalRepository.md) |
| called_by | [TestUpdateRepository_SucceedsForValidGitRepository](/git/git_test/TestUpdateRepository_SucceedsForValidGitRepository.md) |
