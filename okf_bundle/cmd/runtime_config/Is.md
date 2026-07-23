---
okf_version: "0.2"
type: Function
title: Is
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/Is
language: go
---

# Is

## Signature

```go
func (e *runtimeConfigStageError) Is(target error) bool
```

## Source
Lines 40–46 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [runtimeConfigErrorMessage](/cmd/runtime_config/runtimeConfigErrorMessage.md) |
| called_by | [TestRuntimeConfigStageError_NilReceiverAndNilInnerError](/cmd/runtime_config_additional_test/TestRuntimeConfigStageError_NilReceiverAndNilInnerError.md) |
| called_by | [TestBuildEffectiveConfig_RegressionCasesFromFixture](/cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture.md) |
| called_by | [TestBuildEffectiveConfig_WrapsLoadStageErrors](/cmd/runtime_config_test/TestBuildEffectiveConfig_WrapsLoadStageErrors.md) |
| called_by | [TestBuildEffectiveConfig_WrapsValidateStageErrors](/cmd/runtime_config_test/TestBuildEffectiveConfig_WrapsValidateStageErrors.md) |
