# Git Downloader Tool Implementation Specification

## Overview
This specification outlines the implementation requirements for a git repository downloader tool called `git-downloader-tool` that supports generic git repositories with flexible configuration options. The implementation will be written in Go using Cobra for CLI handling and YAML parsing for configuration files.

## Configuration Structure

### Remote Definitions
The tool supports generic git repositories defined in a nested structure:
- `default`: Base URL for git repositories (e.g., https://git.example.com/)

Each remote can be overridden per repository or via command-line interface.
Remote URLs can be specified in three formats:
1. HTTPS: `https://git.example.com/`
2. SSH: `ssh://git@github.com/`
3. Local paths: `/path/to/repositories/`

Each remote can be overridden per repository or via command-line interface.

### Default Values
The tool provides default values for repositories:
- `remote`: Default remote platform (default: github)
- `revision`: Default branch/tag to clone (default: main)
- `path`: Default path for cloning (default: "")

### Repository Definitions
Each repository can be defined with key and value pairs:
- `key`: As repository name (e.g., `my-repo`)
- `value`: Object containing repository-specific configuration options:
  - `remote`: Override remote platform (optional, defaults to defaults.remote, if set to a custom URL, it will be used as-is)
  - `repository`: Custom repository name (required, can be set in CLI with `--repos.<name>.repository=...`)
  - `revision`: Override revision to clone (optional, defaults to defaults.revision)
  - `path`: Override path prefix for cloning (optional, defaults to defaults.path)

## Feature Requirements

### Configuration Handling
1. Support configuration from YAML files with nested structure using go-yaml library
2. Allow CLI overrides for all configuration options using Cobra
3. Handle inheritance from default values when not explicitly set
4. Support zero-value configuration (empty strings should use defaults)
5. Accept custom remote URLs not defined in the config
6. Support configuration validation and error reporting
7. Support local repository cloning
8. Handle different git URL formats (HTTPS, SSH, local paths)
9. Do not support environment variable interpolation in URLs
10. Validate and enforce working directory paths according to specification

### Repository Processing
1. Process repositories in the order they appear in configuration
2. Support special characters in repository names (dashes, underscores)
3. Handle complex path structures
4. Support URL overrides that differ from repository name

### Remote Platform Support
1. HTTPS git repositories with URLs like `https://git.example.com/`
2. SSH git repositories with URLs like `ssh://git@github.com/`
3. Local repository paths like `/path/to/repositories/`
4. Support for all git URL formats in configuration

### CLI Integration
1. Allow setting remote URLs via CLI: `--remote.<name>.url=...`
2. Allow overriding defaults via CLI: `--defaults.remote=...`, `--defaults.revision=...`, `--defaults.path=...`
3. Allow overriding repository settings via CLI: `--repos.<name>.<field>=...`

## Implementation Details

### Technology Stack
1. Language: Go (Golang)
2. CLI Framework: Cobra
3. YAML Parsing: go-yaml library
4. Git Operations: pure git commands executed via os/exec

### Command Structure
The tool will support three main commands:
1. `clone` - Clone repositories according to configuration file
2. `update` - Update existing repositories based on configuration file
3. `cleanup` - Remove repositories that are no longer defined in the configuration

All commands will operate in the same working directory, with the clone command creating new repositories and the update command ensuring existing repositories match the configuration.

### Configuration Merging Strategy
1. Start with default configuration values
2. Apply repository-specific overrides
3. Apply CLI overrides (highest priority)
4. Handle special cases:
   - Empty string for revision should fall back to default
   - Custom remote URLs should be used as-is
5. Validate configuration after merging

### Repository Processing Flow
1. Parse configuration file
2. Validate repository definitions
3. Merge with default values where needed
4. Apply CLI overrides if present
5. Determine working directory based on defaults.path:
   - If defaults.path is "." or empty, use current working directory
   - If defaults.path is a valid path (e.g., "services"), create the directory if it doesn't exist and use it as working directory
   - If defaults.path is ".." or any other invalid path, reject the configuration with error
6. Generate clone/update commands for each repository based on command mode (clone or update)

### Error Handling
1. Validate remote URLs are properly formatted
2. Handle invalid repository names gracefully
3. Provide clear error messages for configuration issues

## Examples from Configuration Files

### Example Configuration 1 (example-config.yaml)
- Repository: `https-repo` with GitHub remote, develop branch, https/ path prefix
- Repository: `ssh-repo` with BitBucket remote, v2.0 tag, ssh/ path prefix
- Repository: `local-repo` with local path, master branch, local/ path prefix
- Repository: `external-repo` with GitHub remote, main branch, external-org/external-repo URL
- Repository: `minimal-service` using all defaults (default remote, main revision, no path)
- Repository: `zero-config-service` with custom path override but default remote and revision
- Repository: `service-with-dashes_and_underscores` with release/v1.0 revision
- Repository: `independent-service` with GitLab remote, feature/new-feature branch, features/ path prefix
- Repository: `complex-path-service` with GitHub remote, main branch, components/services/ path prefix
- Repository: `custom-remote-service` with custom Git server URL, main branch, no path
- Repository: `fully-overridden-service` with GitLab remote, v2.0.0 tag, production/ path prefix, custom URL
- Repository: `ssh-remote-service` with SSH remote URL, develop branch, ssh-clones/ path prefix
- Repository: `relative-local-repo` with relative local path, v1.0 tag, rel/ path prefix
- Repository: `complex-url-repo` with complex URL structure, feature/branch-name branch, projects/ path prefix
- Repository: `default-remote-service` with no remote specified, release/1.0 revision, releases/ path prefix
- Repository: `http-remote-service` with HTTP remote URL, main branch, http/ path prefix
- Repository: `trailing-slash-service` with trailing slash in remote URL, develop branch, trailing/ path prefix
- Repository: `gitlab` with GitLab remote, main branch, custom/ path prefix
- Repository: `123-service` with numeric repository name, v1.2.3 revision, numeric/ path prefix
- Repository: `service-with-very-long-name-that-exceeds-normal-length-limitations` with long name, feature/very-long-branch-name branch, long-names/ path prefix
- Repository: `special-chars-repo` with special characters in URL, release/v2.0 revision, special/ path prefix
- Repository: `multi_dash_and_underscore-service` with multiple dashes and underscores, v1.0.0 revision, multi/ path prefix
- Repository: `port-service` with port in URL, main branch, ports/ path prefix

### Example Configuration 2 (requirements.md)
- Default remote: github
- Default revision: main
- Default path: "services/"
- Repository: `api-gateway` with GitLab remote, develop branch, api/ path prefix, custom URL

### New Example Configuration (generic git support)
- Default remote: https://git.example.com/
- Default revision: main
- Default path: "repos/"
- Repository: `my-project` with HTTPS git remote, v1.0 tag, projects/ path prefix
- Repository: `ssh-repo` with SSH git remote, develop branch, ssh/ path prefix
- Repository: `local-repo` with local path, main branch, local/ path prefix
- Repository: `custom-repo` with custom git server URL, feature branch, custom/ path prefix

## Testing Requirements
1. Validate configuration file parsing with various YAML structures
2. Test all git URL formats (HTTPS, SSH, local paths)
3. Verify default value inheritance
4. Test CLI override functionality with Cobra
5. Validate special character handling in repository names
6. Test concurrent repository cloning scenarios
7. Test error handling for invalid configurations and URLs
8. Test local repository cloning functionality
9. Test SSH key handling and authentication
10. Test edge cases with complex URLs, special characters, and numeric names
11. Test repository name validation and sanitization
12. Test path resolution for local repositories with relative and absolute paths
13. Test `clone` command functionality with various configurations
14. Test `update` command functionality with existing repositories
15. Test `cleanup` command functionality to remove obsolete repositories
16. Test command mode switching (clone vs update)

## Performance Considerations
1. Support concurrent repository cloning where appropriate
2. Efficient configuration parsing and validation
3. Minimal memory usage for large configurations

## Security Considerations
1. Validate remote URLs to prevent malicious redirects
2. Proper handling of SSH and HTTPS protocols
3. Secure storage and handling of credentials (if applicable)
4. Input validation for repository names and paths
5. Do not allow environment variable interpolation in configuration URLs
