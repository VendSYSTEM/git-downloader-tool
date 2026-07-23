---
okf_version: "0.2"
type: Function
title: TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags
resource: cmd/runtime_config_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:56Z"
concept_id: cmd/runtime_config_test/TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags
language: go
---

# TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags

## Signature

```go
func TestApplyDynamicOverrides_UsesChangedInheritedPersistentFlags(t *testing.T)
```

## Source
Lines 321–358 in `cmd/runtime_config_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config_test](/cmd/runtime_config_test.md) |
| calls | [registerDynamicOverrideFlags](/cmd/runtime_config/registerDynamicOverrideFlags.md) |
| calls | [applyDynamicOverrides](/cmd/runtime_config/applyDynamicOverrides.md) |
