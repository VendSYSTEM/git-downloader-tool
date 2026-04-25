# Git Downloader Tool

A command-line tool for managing multiple git repositories through YAML configuration files.

## Features

- Clone repositories from configuration
- Update repositories to latest commits
- Cleanup repositories with optional filtering
- Support for HTTPS, SSH, and local git repositories
- Generic git remote support with custom URLs
- Configuration inheritance from defaults
- CLI override capabilities for all configuration options
- Support for special characters in repository names

## Installation

```bash
go install git-downloader-tool
```

## Usage

```bash
git-downloader-tool [command] [flags]
```

## Project Structure

The tool follows a modular structure:

- `cmd/` - Command implementations (clone, update, cleanup)
- `config/` - Configuration parsing and handling
- `git/` - Git operations (clone, update, cleanup)
- `main.go` - Entry point

## Configuration Format

Configuration is managed through YAML files with the following structure:

### Root Elements

#### `remotes`

Defines named git remote endpoints with URLs. Can be overridden per repository or via CLI.

```yaml
remotes:
  github:
    url: https://github.com/
  gitlab:
    url: https://gitlab.com/
  bitbucket:
    url: ssh://git@bitbucket.org/
```

#### `defaults`

Default values for repository configurations:

```yaml
defaults:
  remote: github        # Default remote platform (default: github)
  revision: main        # Default branch/tag to clone (default: main)
  path: ""              # Default path for cloning (default: "")
```

#### `repos`

Repository definitions with various configurations:

```yaml
repos:
  my-repo:
    remote: github        # Override remote platform (optional)
    repository: "user/my-repo"  # Required: Repository name
    revision: develop     # Override revision to clone (optional)
    path: "my-repo/"      # Override path prefix for cloning (optional)
```

### Configuration Examples

#### Basic Configuration

```yaml
remotes:
  github:
    url: https://github.com/

defaults:
  remote: github
  revision: main
  path: ""

repos:
  my-repo:
    remote: github
    repository: "user/my-repo"
    revision: develop
    path: "my-repo/"
```

#### Advanced Configuration with Multiple Remotes

```yaml
remotes:
  default:
    url: https://git.example.com/
  github:
    url: https://github.com/
  gitlab:
    url: https://gitlab.com/
  bitbucket:
    url: ssh://git@bitbucket.org/

defaults:
  remote: default
  revision: main
  path: ""

repos:
  https-repo:
    remote: github
    revision: develop
    path: "https/"
    repository: "user/https-repo"

  ssh-repo:
    remote: bitbucket
    revision: v2.0
    path: "ssh/"
    repository: "user/ssh-repo"

  local-repo:
    remote: "/local/path/to/repositories/"
    revision: master
    path: "local/"
    repository: "user/local-repo"
```

### Feature Support

#### Git URL Formats

The tool supports various git URL formats:

- HTTPS: `https://github.com/user/repo.git`
- SSH: `ssh://git@github.com/user/repo.git`
- Local paths: `/local/path/to/repositories/`

#### Remote Handling

1. Named remotes defined in `remotes` section can be referenced by name
2. Direct URL overrides are supported for custom git servers
3. Custom remote URLs can be specified directly in repository configuration

#### Default Value Inheritance

- Repositories inherit default values when not explicitly set
- Empty strings for revision fall back to default values
- Path handling follows specification requirements

#### CLI Override Support

All configuration options can be overridden via command line:

- `--remote.<name>.url=...` - Set remote URLs
- `--defaults.remote=...` - Override default remote
- `--defaults.revision=...` - Override default revision
- `--defaults.path=...` - Override default path
- `--repos.<name>.<field>=...` - Override repository fields

## Commands

### Clone

Clone repositories according to configuration file:

```bash
git-downloader-tool clone
```

### Update

Update existing repositories based on configuration file:

```bash
git-downloader-tool update
```

### Cleanup

Remove repositories that are no longer defined in the configuration:

```bash
git-downloader-tool cleanup
```

## Flags

- `--config, -c`: Path to config file (default: ~/.config/git-downloader/config.yaml)
- `--verbose, -v`: Enable verbose output
- `--remote.<name>.url=...`: Override remote URL via CLI
- `--defaults.remote=...`: Override default remote via CLI
- `--defaults.revision=...`: Override default revision via CLI
- `--defaults.path=...`: Override default path via CLI
- `--repos.<name>.<field>=...`: Override repository field via CLI

## Testing Approach

The tool includes comprehensive tests covering:

1. Configuration file parsing with various YAML structures
2. All git URL formats (HTTPS, SSH, local paths)
3. Default value inheritance
4. CLI override functionality with Cobra
5. Special character handling in repository names
6. Concurrent repository cloning scenarios
7. Error handling for invalid configurations and URLs
8. Local repository cloning functionality
9. Edge cases with complex URLs, special characters, and numeric names
10. Command mode switching (clone vs update)
11. Path resolution for local repositories with relative and absolute paths

### Test Coverage

- Unit tests for configuration parsing and validation
- Integration tests for git operations (clone, update)
- End-to-end tests for complete workflow scenarios
- Edge case testing with various configuration combinations

### Test Configuration Examples

The project includes test configurations:

- `example-config.yaml` - Example configuration demonstrating features
- `test-full-config.yaml` - Comprehensive test configuration with edge cases
- `test-config.yaml` - Minimal test configuration

## Security Considerations

1. Validate remote URLs to prevent malicious redirects
2. Proper handling of SSH and HTTPS protocols
3. Input validation for repository names and paths
4. Configuration file security (no environment variable interpolation)
5. Secure handling of credentials if needed

## Performance Considerations

1. Support for concurrent repository cloning where appropriate
2. Efficient configuration parsing and validation
3. Minimal memory usage for large configurations

## Development

### Taskfile workflow (`Taskfile.yml`)

This repository includes a development taskfile at `Taskfile.yml`.

Prerequisites:

- Go `1.25+`
- Task (go-task) `3.x`
- System tools available in `PATH`: `git`, `tar`, `zip`, `shasum`

Use it with:

```bash
task <task>
```

Common tasks:

- `build` - Build a local binary to `dist/bin/`
- `build:multiarch` - Build linux/darwin/windows binaries for amd64 and arm64 to `dist/multiarch/`
- `test` - Run `go test ./...`
- `coverage` - Generate `dist/coverage.out` and `dist/coverage.html`
- `verify` - Run gofmt check, `go vet`, and tests
- `govuln` - Run `govulncheck ./...` (auto-installs if missing)
- `release` - Run clean + verify + coverage + govuln + multiarch builds, package archives and checksums in `dist/release/`

Examples:

```bash
task verify
task coverage
task release
```

### Building (manual)

```bash
go build -o git-downloader-tool main.go
```

### Running Tests (manual)

```bash
go test ./...
```

### Documentation Generation

Documentation is maintained in:

- `README.md` - Main documentation
- `specification.md` - Implementation specification
- `requirements.md` - Requirements specification

## Contributing

1. Fork the repository
2. Create a feature branch
3. Implement your changes
4. Add tests for new functionality
5. Submit a pull request

## License

MIT License - see LICENSE file for details.
