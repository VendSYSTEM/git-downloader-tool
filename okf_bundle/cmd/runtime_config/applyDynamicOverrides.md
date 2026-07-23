---
okf_version: "0.2"
type: Function
title: applyDynamicOverrides
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/applyDynamicOverrides
language: go
---

# applyDynamicOverrides

## Signature

```go
func applyDynamicOverrides(cmd *cobra.Command, loadedConfig *config.Config)
```

## Source
Lines 170–221 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| calls | [getChangedStringFlag](/cmd/runtime_config/getChangedStringFlag.md) |
| called_by | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| called_by | [TestApplyDynamicOverrides_NilInputsAreNoOps](/cmd/runtime_config_additional_test/TestApplyDynamicOverrides_NilInputsAreNoOps.md) |
| called_by | [TestApplyDynamicOverrides_OnlyChangedFlags](/cmd/runtime_config_test/TestApplyDynamicOverrides_OnlyChangedFlags.md) |
| called_by | [TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags](/cmd/runtime_config_test/TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags.md) |
