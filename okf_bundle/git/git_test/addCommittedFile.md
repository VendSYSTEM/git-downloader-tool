---
okf_version: "0.2"
type: Function
title: addCommittedFile
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/addCommittedFile
language: go
---

# addCommittedFile

## Signature

```go
func addCommittedFile(t *testing.T, repoPath, name, content string)
```

## Source
Lines 222–231 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [runCommand](/git/git_test/runCommand.md) |
| called_by | [TestCloneRepository_SyncsWhenTargetExists](/git/git_test/TestCloneRepository_SyncsWhenTargetExists.md) |
| called_by | [createBranchWithFile](/git/git_test/createBranchWithFile.md) |
