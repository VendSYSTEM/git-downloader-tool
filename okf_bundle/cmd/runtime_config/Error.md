---
okf_version: "0.2"
type: Function
title: Error
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/Error
language: go
---

# Error

## Signature

```go
func (e *runtimeConfigStageError) Error() string
```

## Source
Lines 24–30 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [TestRuntimeConfigStageError_NilReceiverAndNilInnerError](/cmd/runtime_config_additional_test/TestRuntimeConfigStageError_NilReceiverAndNilInnerError.md) |
| called_by | [TestBuildEffectiveConfig_RegressionCasesFromFixture](/cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture.md) |
