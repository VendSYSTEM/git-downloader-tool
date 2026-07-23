---
okf_version: "0.2"
type: Function
title: captureStdout
resource: cmd/info_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:info_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:51Z"
concept_id: cmd/info_test/captureStdout
language: go
---

# captureStdout

## Signature

```go
func captureStdout(t *testing.T, run func) string
```

## Source
Lines 66–98 in `cmd/info_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [info_test](/cmd/info_test.md) |
| called_by | [TestInfoCommandUsesEffectiveConfig](/cmd/info_test/TestInfoCommandUsesEffectiveConfig.md) |
| called_by | [TestCommands_PrintRuntimeConfigLoadErrorWhenConfigMissing](/cmd/runtime_config_additional_test/TestCommands_PrintRuntimeConfigLoadErrorWhenConfigMissing.md) |
| called_by | [TestExecute_AllowsOverrideFlagsWhenConfigLoadFails](/cmd/runtime_config_test/TestExecute_AllowsOverrideFlagsWhenConfigLoadFails.md) |
