---
okf_version: "0.2"
type: Function
title: TestInfoCommandUsesEffectiveConfig
resource: cmd/info_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:info_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:51Z"
concept_id: cmd/info_test/TestInfoCommandUsesEffectiveConfig
language: go
---

# TestInfoCommandUsesEffectiveConfig

## Signature

```go
func TestInfoCommandUsesEffectiveConfig(t *testing.T)
```

## Source
Lines 16–64 in `cmd/info_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [info_test](/cmd/info_test.md) |
| calls | [LoadConfig](/config/config/LoadConfig.md) |
| calls | [registerDynamicOverrideFlags](/cmd/runtime_config/registerDynamicOverrideFlags.md) |
| calls | [captureStdout](/cmd/info_test/captureStdout.md) |
