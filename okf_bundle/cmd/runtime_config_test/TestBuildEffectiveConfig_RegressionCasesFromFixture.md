---
okf_version: "0.2"
type: Function
title: TestBuildEffectiveConfig_RegressionCasesFromFixture
resource: cmd/runtime_config_test.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:cmd"
  - "domain:runtime_config_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T14:42:56Z"
concept_id: cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture
language: go
---

# TestBuildEffectiveConfig_RegressionCasesFromFixture

## Signature

```go
func TestBuildEffectiveConfig_RegressionCasesFromFixture(t *testing.T)
```

## Source
Lines 472–615 in `cmd/runtime_config_test.go`

## Relationships

| Type | Target |
|------|--------|
| related | [runtime_config_test](/cmd/runtime_config_test.md) |
| calls | [Is](/cmd/runtime_config/Is.md) |
| calls | [Error](/cmd/runtime_config/Error.md) |
| calls | [LoadConfig](/config/config/LoadConfig.md) |
| calls | [registerDynamicOverrideFlags](/cmd/runtime_config/registerDynamicOverrideFlags.md) |
| calls | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
