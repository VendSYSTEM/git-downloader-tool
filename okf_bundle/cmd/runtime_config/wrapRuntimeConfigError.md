---
okf_version: "0.2"
type: Function
title: wrapRuntimeConfigError
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/wrapRuntimeConfigError
language: go
---

# wrapRuntimeConfigError

## Signature

```go
func wrapRuntimeConfigError(stage error, err error) error
```

## Source
Lines 244–250 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| called_by | [TestWrapRuntimeConfigError_ReturnsNilWhenInnerErrorIsNil](/cmd/runtime_config_additional_test/TestWrapRuntimeConfigError_ReturnsNilWhenInnerErrorIsNil.md) |
| called_by | [TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix](/cmd/runtime_config_test/TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix.md) |
