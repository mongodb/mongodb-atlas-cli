---
name: atlas-cli-open-pr
description: Pre-PR validation and pull request creation workflow for the mongodb-atlas-cli repository. Use when the user asks to open a PR, create a pull request, submit changes for review, or push code to the repository.
---

# Atlas CLI - Open Pull Request

## Prerequisites

Ensure the working directory is the atlas-cli repo root.

## Pre-PR Validation Checklist

Run these checks **in order** before creating the PR. Stop and fix issues at each step before proceeding.

```
Task Progress:
- [ ] Step 1: Verify branch and commit conventions
- [ ] Step 2: Format code
- [ ] Step 3: Fix lint issues
- [ ] Step 4: Build the project
- [ ] Step 5: Run unit tests
- [ ] Step 6: Regenerate docs (if commands changed)
- [ ] Step 7: Verify unit test coverage for changes
- [ ] Step 8: Create the pull request
```

### Step 1: Verify branch and commit conventions

**Branch naming**: Must be a JIRA ticket ID (e.g. `CLOUDP-12345`).

```bash
git rev-parse --abbrev-ref HEAD
```

If the branch name does not match `CLOUDP-NNNNN`, warn the user and ask if they want to rename it.

**Commit messages**: Every commit must start with the ticket ID:

```
CLOUDP-12345: Short description of the change
```

Verify with:

```bash
git log main..HEAD --oneline
```

Warn the user about any non-conforming commits.

### Step 2: Format code

```bash
make fmt
```

If files were modified, stage and commit them using the branch's ticket ID:

```bash
git add -A && git commit -m "CLOUDP-XXXXX: Run make fmt"
```

### Step 3: Fix lint issues

```bash
make fix-lint
```

If auto-fixed, stage and commit. If unfixable errors remain, stop and report to user.

### Step 4: Build the project

```bash
make build
```

If the build fails, stop and report errors.

### Step 5: Run unit tests

```bash
make unit-test
```

If tests fail, stop and report. Do not proceed with failing tests.

### Step 6: Regenerate docs (if commands changed)

If files under `internal/cli/` were modified (`git diff main..HEAD --name-only`), run:

```bash
make gen-docs
```

Stage and commit any doc changes.

### Step 7: Verify unit test coverage

For each changed `.go` file (excluding `_test.go` and `mock`), check for a corresponding `_test.go`. Warn about missing test files but let the user decide whether to proceed.

### Step 8: Create the pull request

```bash
git push -u origin HEAD
```

Then create the PR using the repo's template:

```bash
gh pr create --title "CLOUDP-XXXXX: <description>" --body "$(cat <<'EOF'
## Proposed changes

_Jira ticket:_ CLOUDP-XXXXX

## Checklist

- [x] I have added tests that prove my fix is effective or that my feature works
- [x] I have added any necessary documentation (if appropriate)
- [x] I have run `make fmt` and formatted my code

## Further comments

<context>
EOF
)"
```

Fill in the actual ticket ID and adjust the checklist to reflect what was done.

## Important Notes

- **Never skip lint or test steps** -- they match CI checks.
- Commit any files modified by `make fmt` or `make fix-lint` before opening the PR.
- `make unit-test` may take several minutes; set an appropriate timeout.
- If the user explicitly asks to skip a step, allow it but warn about potential CI failures.
