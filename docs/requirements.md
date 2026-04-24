# requirements

```yaml
---
remote:
  # Name of the urls to use for cloning repositories (can be overridden per repo and in cli `--remote.<name>.url=...`); new repo can be created from CLI with `--remote.<name>.url=...`.
  github:
    # default URL for GitHub (can be overridden per repository)
    url: http://url.com/
  gitlab:
    url: https://url.com/
  bitbucket:
    url: ssh://git@url.com/

# Default values for repositories (can be overridden per repo)
defaults:
  # Default remote to use; can be overridden per repository and in CLI with `--defaults.remote=...`
  remote: github
  # Default branch/tag to clone; can be overridden per repository and in CLI with `--defaults.revision=...`
  revision: main
  # Default path for cloning; can be overridden per repository and in CLI with `--defaults.path=...`
  path: "services/"

# Repository definitions with various configurations; can be modified via CLI with `--repos.<name>.<field>=...`
repos:
  # index name for the repository, do not use it as a reference in the code, use url instead
  api-gateway:
    # Remote can be overridden per repository to use different URL and credentials; can also be overridden in CLI with `--repos.api-gateway.remote=...`
    remote: gitlab
    # Name of the repository (mandatory); can be set in CLI with `--repos.api-gateway.repository=...`
    repository: "api/gateway"
    # Revision can be overridden per repository; can also be overridden in CLI with `--repos.api-gateway.revision=...`
    revision: develop
    # Path can be overridden per repository to specify different cloning location; can also be overridden in CLI with `--repos.api-gateway.path=...`
    path: "api/"
```
