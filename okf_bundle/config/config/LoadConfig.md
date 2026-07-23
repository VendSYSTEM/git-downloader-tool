---
okf_version: "0.2"
type: Function
title: LoadConfig
description: LoadConfig loads configuration from a file path
resource: config/config.go
tags:
  - "lang:go"
  - "type:Function"
  - "module:config"
  - "domain:config.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-25T12:19:15Z"
concept_id: config/config/LoadConfig
language: go
---

# LoadConfig

LoadConfig loads configuration from a file path

## Signature

```go
func LoadConfig(path string) (*Config, error)
```

## Docstring

LoadConfig loads configuration from a file path

## Source
Lines 11–24 in `config/config.go`

## Relationships

| Type | Target |
|------|--------|
| related | [config](/config/config.md) |
| called_by | [TestInfoCommandUsesEffectiveConfig](/cmd/info_test/TestInfoCommandUsesEffectiveConfig.md) |
| called_by | [Execute](/cmd/root/Execute.md) |
| called_by | [buildEffectiveConfig](/cmd/runtime_config/buildEffectiveConfig.md) |
| called_by | [TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression](/cmd/runtime_config_test/TestBuildEffectiveConfig_DynamicRegistrationPrecedenceRegression.md) |
| called_by | [TestBuildEffectiveConfig_LoadApplyMergeValidate](/cmd/runtime_config_test/TestBuildEffectiveConfig_LoadApplyMergeValidate.md) |
| called_by | [TestBuildEffectiveConfig_RegressionCasesFromFixture](/cmd/runtime_config_test/TestBuildEffectiveConfig_RegressionCasesFromFixture.md) |
| called_by | [TestLoadConfig_ReturnsReadErrorForMissingFile](/config/config_test/TestLoadConfig_ReturnsReadErrorForMissingFile.md) |
| called_by | [TestLoadConfig_ReturnsYAMLErrorForInvalidData](/config/config_test/TestLoadConfig_ReturnsYAMLErrorForInvalidData.md) |
| called_by | [TestLoadConfig_Success](/config/config_test/TestLoadConfig_Success.md) |
