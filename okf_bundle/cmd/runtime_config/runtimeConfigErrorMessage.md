---
okf_version: "0.2"
type: Function
title: runtimeConfigErrorMessage
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/runtimeConfigErrorMessage
language: go
---

# runtimeConfigErrorMessage

## Signature

```go
func runtimeConfigErrorMessage(err error) string
```

## Source
Lines 252–265 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| calls | [Is](/cmd/runtime_config/Is.md) |
| called_by | [cleanupCmd](/cmd/cleanup/cleanupCmd.md) |
| called_by | [cloneCmd](/cmd/clone/cloneCmd.md) |
| called_by | [infoCmd](/cmd/info/infoCmd.md) |
| called_by | [TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix](/cmd/runtime_config_test/TestRuntimeConfigErrorMessage_UsesStageSpecificPrefix.md) |
| called_by | [updateCmd](/cmd/update/updateCmd.md) |
