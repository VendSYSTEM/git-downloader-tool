---
okf_version: "0.2"
type: Function
title: getChangedStringFlag
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/getChangedStringFlag
language: go
---

# getChangedStringFlag

## Signature

```go
func getChangedStringFlag(cmd *cobra.Command, name string) (string, bool)
```

## Source
Lines 267–293 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [applyDynamicOverrides](/cmd/runtime_config/applyDynamicOverrides.md) |
| called_by | [TestGetChangedStringFlag_SupportsCommandFlagsAndAbsentFlags](/cmd/runtime_config_additional_test/TestGetChangedStringFlag_SupportsCommandFlagsAndAbsentFlags.md) |
