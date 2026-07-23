---
okf_version: "0.2"
type: Function
title: registerStringPersistentFlag
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/registerStringPersistentFlag
language: go
---

# registerStringPersistentFlag

## Signature

```go
func registerStringPersistentFlag(cmd *cobra.Command, name, value, usage string)
```

## Source
Lines 164–168 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [registerDynamicOverrideFlags](/cmd/runtime_config/registerDynamicOverrideFlags.md) |
| called_by | [registerOverrideLikeFlagsFromArgs](/cmd/runtime_config/registerOverrideLikeFlagsFromArgs.md) |
| called_by | [TestRegisterStringPersistentFlag_ReRegisterDoesNotMarkChanged](/cmd/runtime_config_test/TestRegisterStringPersistentFlag_ReRegisterDoesNotMarkChanged.md) |
