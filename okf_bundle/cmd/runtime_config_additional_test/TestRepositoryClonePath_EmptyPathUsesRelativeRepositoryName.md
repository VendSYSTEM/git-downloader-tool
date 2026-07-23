---
okf_version: "0.2"
type: Function
title: TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName
resource: cmd/runtime_config_additional_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config_additional_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: cmd/runtime_config_additional_test/TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName
language: go
---

# TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName

## Signature

```go
func TestRepositoryClonePath_EmptyPathUsesRelativeRepositoryName(t *testing.T)
```

## Source
Lines 116–121 in `cmd/runtime_config_additional_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config_additional_test](/cmd/runtime_config_additional_test.md) |
| calls | [repositoryClonePath](/cmd/clone/repositoryClonePath.md) |
