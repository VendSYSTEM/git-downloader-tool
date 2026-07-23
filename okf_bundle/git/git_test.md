---
okf_version: "0.2"
type: Module
title: git_test
resource: git/git_test.go
tags:
  - "lang:go"
  - "type:Module"
  - "module:git"
  - "domain:git_test.go"
  - "git:branch:main"
  - "git:repo:git-downloader-tool"
timestamp: "2026-04-27T14:10:08Z"
concept_id: git/git_test
language: go
---

# git_test

## Relationships

| Type | Target |
|------|--------|
| related | [TestCloneRepository_SyncsWhenTargetExists](/git/git_test/TestCloneRepository_SyncsWhenTargetExists.md) |
| related | [TestCloneRepository_SucceedsWithLocalRepository](/git/git_test/TestCloneRepository_SucceedsWithLocalRepository.md) |
| related | [TestCloneRepository_ReturnsErrorWhenCloneFails](/git/git_test/TestCloneRepository_ReturnsErrorWhenCloneFails.md) |
| related | [TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails](/git/git_test/TestCloneRepository_ReturnsErrorWhenRevisionSwitchFails.md) |
| related | [TestUpdateRepository_SucceedsForValidGitRepository](/git/git_test/TestUpdateRepository_SucceedsForValidGitRepository.md) |
| related | [TestUpdateRepository_ReturnsErrorForNonRepositoryPath](/git/git_test/TestUpdateRepository_ReturnsErrorForNonRepositoryPath.md) |
| related | [TestUpdateRepository_ReturnsErrorWhenPullFails](/git/git_test/TestUpdateRepository_ReturnsErrorWhenPullFails.md) |
| related | [TestCleanupRepository_ReturnsNilWhenPathDoesNotExist](/git/git_test/TestCleanupRepository_ReturnsNilWhenPathDoesNotExist.md) |
| related | [TestCleanupRepository_RemovesExistingDirectory](/git/git_test/TestCleanupRepository_RemovesExistingDirectory.md) |
| related | [TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory](/git/git_test/TestCleanupRepository_EmptyPathUsesCurrentWorkingDirectory.md) |
| related | [createCommittedRepo](/git/git_test/createCommittedRepo.md) |
| related | [addCommittedFile](/git/git_test/addCommittedFile.md) |
| related | [createBranchWithFile](/git/git_test/createBranchWithFile.md) |
| related | [runCommand](/git/git_test/runCommand.md) |
