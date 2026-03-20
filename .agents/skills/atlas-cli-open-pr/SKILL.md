---
name: atlas-cli-open-pr
description: Pre-PR validation and pull request creation workflow for the mongodb-atlas-cli repository. Use when the user asks to open a PR, create a pull request, submit changes for review, or push code to the repository.
---

# Atlas CLI - Open Pull Request

## Step 1: Verify branch naming

The branch name **must** be a JIRA ticket ID (e.g. `CLOUDP-12345`).

```bash
git rev-parse --abbrev-ref HEAD
```

If it doesn't match `CLOUDP-NNNNN`:
1. Ask the user for the JIRA ticket ID.
2. Rename the branch: `git branch -m CLOUDP-XXXXX`

## Step 2: Run pre-PR checks

Run the validation script. It formats, lints, builds, tests, and regenerates docs — printing only errors:

```bash
./scripts/agent-pr-hook.sh
```

If any check fails, fix the reported issues and re-run until all checks pass. Stage and commit any files modified by formatting or doc generation.

## Step 3: Verify unit test coverage

For each changed `.go` file (excluding `_test.go` and `mock`), check for a corresponding `_test.go`. Warn about missing test files but let the user decide whether to proceed.

## Step 4: Push and create PR (optional)

**Ask the user before proceeding with this step.** Only push and create a PR if the user explicitly confirms.

```bash
git push -u origin HEAD
```

Then create the PR using the repo's [PR template](.github/pull_request_template.md):

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
