---
okf_version: "0.2"
type: Function
title: TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix
resource: cmd/runtime_config_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:56Z"
concept_id: cmd/runtime_config_test/TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix
language: go
---

# TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix

## Signature

```go
func TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix(t *testing.T)
```

## Source
Lines 665–683 in `cmd/runtime_config_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config_test](/cmd/runtime_config_test.md) |
| calls | [wrapRuntimeConfigError](/cmd/runtime_config/wrapRuntimeConfigError.md) |
| calls | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
