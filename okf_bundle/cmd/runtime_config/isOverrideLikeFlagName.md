---
okf_version: "0.2"
type: Function
title: isOverrideLikeFlagName
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/isOverrideLikeFlagName
language: go
---

# isOverrideLikeFlagName

## Signature

```go
func isOverrideLikeFlagName(flagName string) bool
```

## Source
Lines 112–116 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [registerOverrideLikeFlagsFromArgs](/cmd/runtime_config/registerOverrideLikeFlagsFromArgs.md) |
