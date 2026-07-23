---
okf_version: "0.2"
type: Function
title: registerDynamicOverrideFlags
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/registerDynamicOverrideFlags
language: go
---

# registerDynamicOverrideFlags

## Signature

```go
func registerDynamicOverrideFlags(cmd *cobra.Command, loadedConfig *config.Config)
```

## Source
Lines 118–162 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| calls | [registerStringPersistentFlag](/cmd/runtime_config/registerStringPersistentFlag.md) |
| called_by | [TestInfoCommandUsesEffectiveConfig](/cmd/info_test/TestInfoCommandUsesEffectiveConfig.md) |
| called_by | [Execute](/cmd/root/Execute.md) |
| called_by | [TestRegisterDynamicOverrideFlags_NilInputsAreNoOps](/cmd/runtime_config_additional_test/TestRegisterDynamicOverrideFlags_NilInputsAreNoOps.md) |
| called_by | [TestApplyDynamicOverrides_OnlyChangedFlags](/cmd/runtime_config_test/TestApplyDynamicOverrides_OnlyChangedFlags.md) |
| called_by | [TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags](/cmd/runtime_config_test/TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags.md) |
| called_by | [TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression](/cmd/runtime_config_test/TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression.md) |
| called_by | [TestBuildEffectiveConfig_LoadApplyMergeValidate](/cmd/runtime_config_test/TestBuildEffectiveConfig_LoadApplyMergeValidate.md) |
| called_by | [TestBuildEffectiveConfig_RegressionCasesFromFixture](/cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture.md) |
| called_by | [TestRegisterDynamicOverrideFlags_CanBeCalledTwice](/cmd/runtime_config_test/TestRegisterDynamicOverrideFlags_CanBeCalledTwice.md) |
| called_by | [TestRegisterDynamicOverrideFlags_UsesReposNamespace](/cmd/runtime_config_test/TestRegisterDynamicOverrideFlags_UsesReposNamespace.md) |
