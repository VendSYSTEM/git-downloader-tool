# Git Downloader Tool

A command-line utility for managing multiple Git repositories through centralized configuration.

## Overview

The git-downloader-tool is a Go-based CLI utility that consumes `git-downloader.yaml` configuration files to provide both listing and cloning capabilities for Git repositories. It simplifies managing multiple repositories by allowing you to define all your repository configurations in a single file and then list or clone them with consistent settings.

## Features

- **Centralized Configuration**: Define all repositories in a single `git-downloader.yaml` file
- **Repository Listing**: View configured remotes, defaults, and repositories with their settings
- **Git Cloning**: Clone repositories with proper inheritance handling
- **CLI Override Support**: Override any configuration values via command-line flags
- **Dry Run Mode**: Preview what would be cloned without actually performing the clone operation
- **Inheritance Logic**: Repositories can inherit default values from the configuration

## Installation

### Prerequisites

- Go 1.21 or later
- Git installed and available in PATH

### Build from Source

```bash
git clone <repository-url>
cd git-downloader-tool
go build -o git-downloader main.go
```

### Install with Go

```bash
go install <repository-url>/git-downloader-tool@latest
```

## Configuration

The tool expects a `git-downloader.yaml` file in the current directory with the following structure:

```yaml
remote:
  github: https://github.com/
  gitlab: https://gitlab.com/

defaults:
  remote: github
  revision: main
  path: ""

repos:
  - name: "my-repo"
    remote: github
    revision: main
    path: ""
    url: "my-repo"

  - name: "another-repo"
    remote: gitlab
    revision: develop
    path: ""
    url: "another-repo"
```

## Usage

### Listing Commands

```bash
# List all configuration
./git-downloader

# List remotes only
./git-downloader --list-remotes

# List defaults only
./git-downloader --list-defaults

# List repositories only
./git-downloader --list-repos
```

### Cloning Commands

```bash
# Clone all repositories (default behavior)
./git-downloader clone

# Dry run - show what would be cloned
./git-downloader clone --dry-run

# Clone with override for revision
./git-downloader clone --revision=develop

# Clone with override for remote
./git-downloader clone --remote=gitlab
```

## Command-Line Flags

| Flag              | Description                                        |
| ----------------- | -------------------------------------------------- |
| `--list-remotes`  | List configured remotes                            |
| `--list-defaults` | List default values                                |
| `--list-repos`    | List all repositories                              |
| `--dry-run`       | Show what would be cloned without actually cloning |
| `--remote`        | Override the remote URL                            |
| `--revision`      | Override the revision to clone                     |
| `--path`          | Override the path for cloning                      |

## Simple Example Configuration

```yaml
remote:
  github: https://github.com/
  gitlab: https://gitlab.com/

defaults:
  remote: github
  revision: main
  path: ""

repos:
  - name: "my-project"
    url: "my-organization/my-project"
    revision: main
    path: "projects/"

  - name: "another-repo"
    remote: gitlab
    url: "my-organization/another-repo"
    revision: develop
```

## Extensive Example Configuration

[Example Configuration](./example-config.yaml) is an extensive example configuration file demonstrating various use cases, including multiple remotes, default values, and repository-specific overrides.

### Configuration Notes

1. **Remote Definitions**: Remotes can be simple string URLs or nested structures with URL fields
2. **Inheritance**: Repositories inherit default values unless explicitly overridden
3. **Path Handling**: Paths are relative to the current working directory, with trailing slashes
4. **Revision Override**: Can specify branches, tags, or commit hashes
5. **URL Override**: Explicit URL can be provided instead of deriving from name
6. **Empty Values**: Empty strings for revision should fall back to defaults
7. **Special Characters**: Names can include dashes, underscores, and other characters
8. **Custom Remotes**: Can use custom remote URLs not defined in the remote section
9. **Flexible Path Prefixes**: Paths can be nested with multiple directory levels

### Usage Examples

```bash
# List all repositories with their configurations
./git-downloader --list-repos

# Clone all repositories with defaults
./git-downloader clone

# Clone with revision override for all repositories
./git-downloader clone --revision=develop

# Clone with remote override for all repositories
./git-downloader clone --remote=gitlab

# Dry run to see what would be cloned
./git-downloader clone --dry-run

# Clone specific repository with custom configuration
./git-downloader clone --revision=feature/new-feature --path="features/"
```

### Edge Cases Covered

1. **Empty repository definitions**: Repositories defined without any configuration fields
2. **Path prefix handling**: Different path configurations including nested paths
3. **Remote override flexibility**: Custom remotes and default remotes
4. **Revision specification**: Branch names, tags, and commit hashes
5. **URL override scenarios**: Explicit URLs that differ from derived names
6. **Zero-value configurations**: Empty strings and special cases
7. **Special character handling**: Names with various characters
8. **Custom remote URLs**: Non-standard remote definitions
9. **Complete override scenarios**: All fields overridden for maximum flexibility

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

This tool was created to simplify managing multiple Git repositories through centralized configuration.
