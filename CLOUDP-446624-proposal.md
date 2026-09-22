# CLOUDP-446624: Safe auto-approval and auto-merge for Atlas CLI PRs

## Goal

Enable approved Atlas CLI automation to approve and auto-merge eligible bot PRs without allowing a PR to merge while required GitHub or Evergreen checks are failing.

## Context

Atlas CLI already has several automation paths that create PRs and request auto-merge:

* `.github/workflows/dependabot-jira.yaml` runs when Dependabot opens a PR. It delegates to `mongodb/apix-action/.github/workflows/_dependabot-jira.yaml@f7d2400a387dea2661651b14bbaf8c8f46c8bf66`.
* The reusable Dependabot workflow creates a Jira ticket, comments on the PR, adds the `auto_close_jira` label, approves the PR with `gh pr review --approve`, and runs `gh pr merge --auto --squash`.
* `.github/workflows/autoupdate-sdk.yaml`, `.github/workflows/autoupdate-spec.yaml`, and `.github/workflows/update-e2e-tests.yml` create Atlas CLI update PRs with the APIx bot token and run `gh pr merge --auto --squash`. They do not currently approve their PRs.

The active GitHub ruleset for the default branch is named `master`. It applies to `~DEFAULT_BRANCH` and requires:

* one approving PR review;
* code-owner review;
* stale review dismissal on new pushes;
* required linear history;
* no deletion or non-fast-forward updates;
* these required status checks: `evergreen/code_health`, `docs`, `mocks`, `tidy`, `lint`, `shellcheck`, `unit-tests (macos-latest)`, `unit-tests (ubuntu-latest)`, `unit-tests (windows-latest)`, `libraryOwners`, and `licensecheck`.

The classic branch-protection API reports `master` as unprotected, so checks against branch protection need to read GitHub rulesets instead.

`CODEOWNERS` currently assigns all files to `@mongodb/apix-devtools`, with Docs ownership for `/docs/` and generated API docs assigned back to `@mongodb/apix-devtools`.

The referenced PR, https://github.com/mongodb/mongodb-atlas-cli/pull/4769, shows the current concern. `apix-bot` approved and enabled auto-merge soon after Dependabot opened the PR. A human code-owner approval came later before the PR merged.

## Proposal

Keep the merge gate in GitHub rulesets and change the automation so approval happens only after GitHub reports every required check as successful.

The Dependabot path should do this in the reusable `mongodb/apix-action` workflow, because this repo only calls that workflow:

1. Keep Jira creation, the tracking comment, and the `auto_close_jira` label at PR-open time.
2. Add an eligibility gate before any approval or auto-merge call. For Atlas CLI, eligible PRs should be limited to Dependabot PRs where `dependabot/fetch-metadata` reports a patch update, a minor update, or a security update. Major updates should still create Jira tracking but should not receive bot approval or auto-merge.
3. Before approving, wait for required checks to finish and fail closed if any required check fails or times out. The implementation can use `gh pr checks "$PR_URL" --required --watch` with a bounded timeout, or an equivalent GraphQL check against the repository ruleset and PR check rollup.
4. After the required checks pass, run `gh pr review "$PR_URL" --approve` with the APIx bot token.
5. Then run `gh pr merge "$PR_URL" --auto --squash` with the same token. This schedules the merge through GitHub rather than merging directly.

For Atlas CLI owned automation PRs from `autoupdate-sdk.yaml`, `autoupdate-spec.yaml`, and `update-e2e-tests.yml`, use the same green-first rule if we decide these PRs are allowed to receive bot approval. These workflows can either call a shared APIx action step or keep the approval logic local. The behavior should be identical either way: no approval before required checks pass, and no direct merge.

Do not weaken CODEOWNERS for normal code. The current ruleset already requires code-owner review. If `apix-bot` approval does not satisfy that requirement, GitHub should leave auto-merge pending until a human code owner approves. If we want these PRs to merge with no human review, first verify whether GitHub can treat the bot as a valid code owner for the relevant generated or dependency files. If it cannot, we should not use ruleset bypass as a shortcut because it risks bypassing more than the intended review requirement.

## Expected behavior

Failing PR:

* The workflow creates or updates Jira tracking.
* The workflow does not approve the PR.
* The workflow does not enable auto-merge, or it leaves auto-merge disabled after the check wait fails.
* Jira remains open for follow-up.

Partially green PR:

* The workflow waits until required checks finish.
* If the wait times out, the workflow exits without approving or enabling auto-merge.
* A retry through `workflow_dispatch` or a later scheduled scan can approve the PR after checks become green.

Fully green eligible PR:

* The workflow approves the PR.
* The workflow enables squash auto-merge.
* GitHub merges only when the ruleset is satisfied, including any code-owner-review requirement that still applies.
* The `auto_close_jira` label closes the Jira ticket after merge.

Ineligible PR:

* The workflow creates Jira tracking.
* The workflow does not approve or enable auto-merge.
* A human reviewer handles the PR.

## Implementation steps

1. Update `mongodb/apix-action/.github/workflows/_dependabot-jira.yaml` so `Approve PR` and `Auto merge PR` run only after an eligibility check and a required-checks wait succeed.
2. Add a bounded timeout around the wait. A timeout should fail closed and leave a clear log message with the pending or failed checks.
3. Add `workflow_dispatch` support, or a small scheduled retry workflow, so an operator can re-run the approval path after transient CI or Evergreen delays without reopening the PR.
4. Decide whether Atlas CLI's generated-update workflows should also auto-approve their PRs. If yes, share the same check-waiting logic instead of copying ad hoc shell into each workflow.
5. Verify whether `apix-bot` approval counts as a code-owner review for files owned by `@mongodb/apix-devtools`. Keep CODEOWNERS unchanged until this is proven.
6. Document the behavior in the workflow comments or repo maintainer docs so future changes do not reintroduce approval-before-green behavior.

## Validation plan

Use test PRs against `mongodb/mongodb-atlas-cli` or a temporary fork with the same ruleset shape.

1. Open a Dependabot-style PR with a deliberately failing required check. Confirm the workflow creates Jira tracking but does not approve or enable auto-merge.
2. Open a PR where GitHub checks pass but `evergreen/code_health` fails. Confirm the workflow does not approve or enable auto-merge.
3. Open a PR where all required checks pass. Confirm the workflow approves the PR and enables auto-merge.
4. Push a new commit after bot approval. Confirm stale-review dismissal removes the approval and the workflow must wait for the new checks before approving again.
5. Confirm whether a bot approval alone satisfies the ruleset's code-owner-review requirement. If it does not, confirm auto-merge remains pending until a human code owner approves.

## Open questions

* Should Atlas CLI generated-update PRs receive bot approval, or should this task cover only Dependabot PRs?
* Should major Dependabot updates ever be eligible, or should they always require human review?
* If a required check is skipped because the workflow path does not apply, does the ruleset treat that check as satisfied for every eligible automation PR type?
* Where should the long-term behavior documentation live? This repo has generated `docs/`, so a maintainer-facing note may belong in `CONTRIBUTING.md` or in the shared `apix-action` repository instead.
