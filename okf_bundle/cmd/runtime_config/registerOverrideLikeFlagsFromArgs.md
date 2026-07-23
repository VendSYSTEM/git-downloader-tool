---
okf_version: "0.2"
type: Function
title: registerOverrideLikeFlagsFromArgs
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/registerOverrideLikeFlagsFromArgs
language: go
---

# registerOverrideLikeFlagsFromArgs

## Signature

```go
func registerOverrideLikeFlagsFromArgs(cmd *cobra.Command, args []string)
```

## Source
Lines 80–110 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| calls | [isOverrideLikeFlagName](/cmd/runtime_config/isOverrideLikeFlagName.md) |
| calls | [registerStringPersistentFlag](/cmd/runtime_config/registerStringPersistentFlag.md) |
| called_by | [Execute](/cmd/root/Execute.md) |
| called_by | [TestRegisterOverrideLikeFlagsFromArgs_NilCommandIsNoOp](/cmd/runtime_config_additional_test/TestRegisterOverrideLikeFlagsFromArgs_NilCommandIsNoOp.md) |
| called_by | [TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression](/cmd/runtime_config_test/TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression.md) |
| called_by | [TestPreregisterOverrideFlagsFromArgs_RegistersMatchingLongFlags](/cmd/runtime_config_test/TestPreregisterOverrideFlagsFromArgs_RegistersMatchingLongFlags.md) |
| called_by | [TestPreregisterOverrideFlagsFromArgs_StopsAtEndOfOptionsMarker](/cmd/runtime_config_test/TestPreregisterOverrideFlagsFromArgs_StopsAtEndOfOptionsMarker.md) |
