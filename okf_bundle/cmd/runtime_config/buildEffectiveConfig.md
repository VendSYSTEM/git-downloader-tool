---
okf_version: "0.2"
type: Function
title: buildEffectiveConfig
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/buildEffectiveConfig
language: go
---

# buildEffectiveConfig

## Signature

```go
func buildEffectiveConfig(cmd *cobra.Command) (*config.Config, error)
```

## Source
Lines 223–242 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| calls | [LoadConfig](/config/config/LoadConfig.md) |
| calls | [wrapRuntimeConfigError](/cmd/runtime_config/wrapRuntimeConfigError.md) |
| calls | [applyDynamicOverrides](/cmd/runtime_config/applyDynamicOverrides.md) |
| calls | [MergeWithDefaults](/config/config/MergeWithDefaults.md) |
| calls | [Validate](/config/config/Validate.md) |
| called_by | [cleanupCmd](/cmd/cleanup/cleanupCmd.md) |
| called_by | [cloneCmd](/cmd/clone/cloneCmd.md) |
| called_by | [infoCmd](/cmd/info/infoCmd.md) |
| called_by | [TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression](/cmd/runtime_config_test/TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression.md) |
| called_by | [TestBuildEffectiveConfig_LoadApplyMergeValidate](/cmd/runtime_config_test/TestBuildEffectiveConfig_LoadApplyMergeValidate.md) |
| called_by | [TestBuildEffectiveConfig_RegressionCasesFromFixture](/cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture.md) |
| called_by | [TestBuildEffectiveConfig_WrapsLoadStageErrors](/cmd/runtime_config_test/TestBuildEffectiveConfig_WrapsLoadStageErrors.md) |
| called_by | [TestBuildEffectiveConfig_WrapsValidateStageErrors](/cmd/runtime_config_test/TestBuildEffectiveConfig_WrapsValidateStageErrors.md) |
| called_by | [updateCmd](/cmd/update/updateCmd.md) |
