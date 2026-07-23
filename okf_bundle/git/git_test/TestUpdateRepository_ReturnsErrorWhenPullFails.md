---
okf_version: "0.2"
type: Function
title: TestUpdateRepository_ReturnsErrorWhenPullFails
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test/TestUpdateRepository_ReturnsErrorWhenPullFails
language: go
---

# TestUpdateRepository_ReturnsErrorWhenPullFails

## Signature

```go
func TestUpdateRepository_ReturnsErrorWhenPullFails(t *testing.T)
```

## Source
Lines 130–142 in `git/git_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [git_test](/git/git_test.md) |
| calls | [createCommittedRepo](/git/git_test/createCommittedRepo.md) |
| calls | [UpdateRepository](/git/git/UpdateRepository.md) |
| calls | [Error](/cmd/runtime_config/Error.md) |
