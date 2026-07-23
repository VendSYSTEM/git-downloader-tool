---
okf_version: "0.2"
type: Function
title: TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory
language: go
---

# TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory

## Signature

```go
func TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory(t *testing.T)
```

## Source
Lines 172–200 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [CleanupRepository](/git/git/CleanupRepository.md) |
