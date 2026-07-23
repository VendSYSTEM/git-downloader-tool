---
okf_version: "0.2"
type: Function
title: Unwrap
resource: cmd/runtime_config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:40Z"
concept_id: cmd/runtime_config/Unwrap
language: go
---

# Unwrap

## Signature

```go
func (e *runtimeConfigStageError) Unwrap() error
```

## Source
Lines 32–38 in `cmd/runtime_config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config](/cmd/runtime_config.md) |
| called_by | [TestRuntimeConfigStageError_NilReceiverAndNilInnerError](/cmd/runtime_config_additional_test/TestRuntimeConfigStageError_NilReceiverAndNilInnerError.md) |
