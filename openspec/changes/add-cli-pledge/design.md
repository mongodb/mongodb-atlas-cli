## Context

`atlas-cli` is a single Go binary that wraps the Atlas Admin API. Today its blast radius is bounded only by the credentials it loads (env vars, profile, or service-account key). When invoked by an AI coding agent, those credentials are typically the same human-Owner credentials the developer uses interactively — there is no narrower channel.

The proposal introduces `atlas pledge`, modeled loosely on OpenBSD's `pledge(2)`. The model is **in-process voluntary restriction**: the CLI promises not to use certain operations and refuses to issue them, regardless of what the caller asks for. This is the same shape as `sudo`'s timestamp / `git`'s `core.fsmonitor` — useful precisely because it cannot be circumvented by the workflow it protects, even though a determined operator with the same credentials can step around it (by calling the API directly, by clearing the state file, etc.).

Existing relevant surface:
- `internal/api/commands.go` is generated and already carries the HTTP verb per operation, which gives us the read/write split nearly for free.
- `internal/api/executor.go` is the single dispatch point for generated API commands.
- Hand-written commands under `internal/cli/**` each declare a Cobra `*Builder()` and an `*Opts.Run()`. There is no central executor — enforcement has to be added per command, or wedged into a Cobra `PreRunE` registered at root.
- The repo already supports `make gen-api-commands` and `make gen-mocks`, so adding a generated override table fits the existing pipeline.

## Goals / Non-Goals

**Goals:**
- Make it impossible for `atlas` (and any `atlas` child invocation in the same shell session) to execute a write or admin operation while a `readonly` pledge is active, without modifying the wrapping shell or the agent.
- Make pledge monotonic within a session and inheritable across `exec`/`fork` of the same `atlas` binary in the same SID.
- Give the developer an audit trail of what an agent attempted but was prevented from doing.
- Provide a clean way to grant a one-off exception from a different terminal, scoped to a single operation invocation.
- Make installing the integration into Claude Code (and equivalents) one command, idempotent, and removable.

**Non-Goals:**
- This is not a credential boundary. A user with the loaded API key can always call the Atlas API directly via `curl` or any other client. Document this explicitly.
- This is not a kernel sandbox. We do not call `seccomp`, `pledge(2)`, or `setrlimit`. Pledge is enforced inside the Go process at the dispatch boundary.
- Windows support is deferred. Windows has no POSIX session ID; emulating one across `cmd.exe`/PowerShell/conhost reliably is its own project. v1 prints a clear "not supported" message on Windows and exits non-zero from `atlas pledge`.
- No server-side change. Atlas RBAC, project roles, and API key scoping are unchanged.
- No daemon. Enforcement is purely file-based; no background process.

## Decisions

### D1. Session identity = POSIX SID, not env var, not PID

**Decision:** Anchor the pledge to `getsid(0)` (Linux/macOS) and store state at `$XDG_STATE_HOME/atlascli/pledge/<sid>.json`.

**Why:** Pledge must be shell-resistant — a child `atlas` must inherit, including under `bash -c`, `xargs atlas`, `&& atlas`, and `exec atlas`. Options considered:

- **Env var (`ATLAS_PLEDGE=...`)**: trivially defeated by `env -i atlas …`, by `unset ATLAS_PLEDGE`, or by an agent that builds its own environment for child processes. Rejected as primary mechanism.
- **PPID chain walk**: traverse parents looking for a marker file. Fragile because the parent shell may not be the SID leader, parents can re-parent to init, and the walk is racy under process death.
- **POSIX session ID (`getsid`)**: stable for the lifetime of the controlling terminal, inherited by every descendant of the session leader, survives `setsid`-free `fork`/`exec`, and cannot be reset without `setsid()` (which detaches from the terminal — visible to the user). New terminals / new logins get new SIDs, which matches the desired "fresh session resets pledge" behavior.

Trade-off: a user who really wants to escape can run `setsid atlas ...` to get a new SID. We accept this — it's loud, explicit, and not something an agent will stumble into. We log a one-line note at pledge load time when `getsid(0) != getsid(getppid())` so the audit trail flags the escape.

GUI agent launchers (Claude Desktop, IDE plugins) simply get their own SID like any other process. If the user has not run `atlas pledge` in that SID, `atlas` works as normal — same as it does today. If the user wants the GUI agent constrained, they pledge in that SID (e.g., via the agent's own SessionStart hook installed by `atlas hook install`).

### D2. Enforcement runs at exactly one chokepoint per command class

**Decision:**
- Generated API commands: enforcement runs in `internal/api/executor.go` *before* the HTTP request is built. Operation tier is read from the new field on `shared_api.Command` populated by `make gen-api-commands`.
- Hand-written commands: a root-level Cobra `PersistentPreRunE` resolves the leaf command's permission tier via an annotation on the Cobra command (`cmd.Annotations["atlas.permission"]`). Missing annotation == lint-fail; see D5.

**Why:** Two chokepoints (one per command class) instead of decorating every `Run()` method. Each chokepoint is the single source of truth for its class and shows up clearly in code review when changed.

**Alternative considered:** intercept at the HTTP transport layer (`http.RoundTripper`). Rejected because (a) some commands batch multiple HTTP calls and we want to fail before partial execution, and (b) it would not stop CLI-local destructive actions (e.g., a hypothetical future `atlas config wipe`).

### D3. Permission tier derivation

**Decision:** Default by HTTP verb — GET → `read`; POST/PUT/PATCH/DELETE → `write`. Maintain a hand-curated override file `tools/api-generator/permission-overrides.yaml` listing operationIDs that are semantically read despite using POST (search/aggregation endpoints). Generation merges the override list into `commands.go` as a `Permission` field. Admin-tier endpoints (e.g., anything touching billing, org-level role grants, key creation) are listed in the same file under an `admin` key.

**Why:** HTTP verb is right ~95% of the time and gets us a usable v1 without endpoint-by-endpoint review. The override file is small, reviewable, and grows over time. Storing it next to the generator keeps it in the existing regeneration loop rather than as a runtime config the user could tamper with.

### D4. Monotonic narrowing, with the `--allow` token as the only widening path

**Decision:** `atlas pledge` writes the new pledge only if it is a subset of the current one. To widen, the user must (a) end the session, or (b) use the token approval flow for a specific operation. The pledge file carries a version counter that increments on every narrowing.

**Why:** Mirrors OpenBSD pledge's monotonicity, which is the property that makes pledge trustworthy: once an agent observes a narrow pledge, it can't talk the user into widening it without an out-of-band action.

### D5. Build-time guarantee that every command has a tier

**Decision:** A unit test in `internal/cli/...` walks the entire Cobra tree at test time and asserts that every leaf command either has `Annotations["atlas.permission"]` set, or is one of a small allowlisted set of meta-commands (`help`, `completion`, `version`, `pledge`, `hook`). The generated API commands carry their permission inline so they are covered by construction.

**Why:** Without this, a single un-annotated command silently becomes a pledge bypass. Catching it in CI is cheap; catching it in production is a security incident.

### D6. Token approval flow

**Decision:**
- Blocked command (on a TTY) writes a request file to `<state>/pledge/<sid>/requests/<token>.json` containing `{operationID, paramsHash, requestingSID, requestedAt, ttl=5m, requestedTier}`, prints the token, and `inotify`-watches (Linux) / `kqueue`-watches (macOS) `<state>/pledge/<sid>/grants/` for a matching file.
- Approver runs `atlas pledge --allow <token>` in another session. The approver process verifies its own pledge permits the requested tier; if so, it writes `<state>/pledge/<sid>/grants/<token>.json` with `{approverSID, approvedAt, signature}` (signature = HMAC over the request contents with a key derived from a once-generated per-host secret in `<state>/pledge/.hmac`, mode 0600).
- Original process verifies the signature, marks the grant as consumed by renaming it to `<token>.consumed`, and proceeds. Audit log gets `allowed-by-token`.

**Why:** Filesystem rendezvous is simpler than a unix socket, works in containers without extra setup, and the file modes give us the security properties we need (only the same UID can write). HMAC binding to a per-host secret stops a malicious bystander from forging a grant by dropping a file into the directory if they're not the same UID — though in practice file mode 0600 plus the directory mode is the real boundary.

**Bound by paramsHash, not just operationID:** approving `deleteCluster` should not be a free pass to delete *all* clusters. The hash covers normalized argv + flags.

**TTY-only by default:** when running non-interactively (no controlling TTY), blocked commands fail immediately without printing a token. Agents running headlessly will fail fast rather than hang. A flag `--wait-for-approval` can opt in.

### D7. Hook installation: idempotent JSON merge with markers

**Decision:** For `claude-code` and `codex`, hook install reads the relevant settings.json, parses it as JSON5 (Claude Code allows comments) or strict JSON (Codex), inserts/updates a single entry tagged with a stable marker `"_atlas_managed": true` and a version field, writes the result to a temp file, and atomically renames. A `.bak` snapshot is taken first.

For `shell`, we print a sourceable snippet wrapped in `# >>> atlas pledge >>>` / `# <<< atlas pledge <<<` marker comments (same convention as `conda init`, `nvm`, etc.), and only write to a shell rc file with explicit `--write <path>`. Re-running with `--write` rewrites the block in place between the markers.

**Why:** Settings files are user-owned and may contain unrelated entries. Markers let `uninstall` find exactly what we wrote without diffing. Atomic rename + `.bak` prevents partial-write corruption.

### D8. State directory layout

```
$XDG_STATE_HOME/atlascli/pledge/
  .hmac                         # per-host HMAC secret, mode 0600
  audit.log                     # append-only JSON lines
  <sid>.json                    # active pledge for this session
  <sid>/requests/<token>.json   # pending approval requests
  <sid>/grants/<token>.json     # approver-signed grants (renamed .consumed on use)
```

Stale SID files (no process in that SID exists anymore) are garbage-collected lazily on the next pledge write — cheap, no daemon.

### D9. Non-API destructive operations

Some hand-written commands (`atlas config delete`, decryption operations, `atlas auth logout` in some flows) mutate local state, not Atlas. These are tagged `local-write` and treated as `write` under the readonly pledge so an agent can't, for example, log the user out from under them.

## Risks / Trade-offs

- **Risk:** A user can bypass pledge with `setsid atlas ...` or by `rm`-ing the pledge file.
  **Mitigation:** Document this clearly. Pledge is a guardrail, not a sandbox. The audit log records when SID lineage changes between pledge and use; agent workflows do not naturally produce `setsid`, so a `setsid` use in a transcript is itself a signal.

- **Risk:** Credentials remain in env / config files, so an agent that ignores `atlas-cli` and calls the Atlas REST API with `curl` is unrestricted.
  **Mitigation:** This is by design; pledge protects users *from `atlas-cli` itself*. The hook for Claude Code only constrains `atlas`-shaped commands. We mention in docs that for stronger isolation, users should provision a scoped service-account key.

- **Risk:** A blocking, hanging command (waiting for approval) inside an agent's tool-use loop will look like a timeout to the agent.
  **Mitigation:** Default `--wait-for-approval=false` when stdout is not a TTY. Agents see a fast non-zero exit with structured error output naming the operation, so they can report back to the human.

- **Risk:** SID is shared across all `atlas` invocations in a tmux window, including ones the user runs themselves. The user's own command gets blocked.
  **Mitigation:** That's the intended semantics — pledge is per-session, not per-agent. Users who want to escape briefly can use `setsid atlas …` or open a new terminal. We surface this in the error message when blocking.

- **Risk:** `make gen-api-commands` regeneration loses the override table.
  **Mitigation:** The override file is an input to the generator, not an output, and is checked into the repo. A unit test asserts every operationID in the override file exists in `commands.go` so removed endpoints don't silently drop their override.

- **Risk:** `~/.claude/settings.json` schema evolves and our merge breaks.
  **Mitigation:** Use a schema-aware merge keyed on the well-known `hooks.SessionStart` path documented at https://code.claude.com/docs/en/hooks.md, with a fallback to refuse-and-print if structure is unrecognized. Cover with golden-file tests.

## Migration Plan

This is additive. No existing behavior changes when no pledge is active. Rollout:

1. Land the permission tier on generated commands (no enforcement) and the override table. Verify CI green.
2. Land the hand-written command annotation requirement plus the CI test. Annotate existing commands as separate, reviewable PRs grouped by area.
3. Land the `pledge` command, enforcement chokepoints, audit log, and state machinery. Behind no flag — if no pledge file exists, everything works as before.
4. Land `hook install/uninstall` and docs.
5. Token approval flow last; it depends on (3) being stable.

Rollback for any step is a straight revert; no schema or persisted state matters until step 3, and step-3 state is per-session and self-cleaning.

## Open Questions

- Should `atlas hook install claude-code` default to `readonly` or to a `read-write` profile that excludes only destructive verbs? Leaning `readonly` (safer default; users opt up).
- Should the audit log rotate? For v1, no — it's local, append-only, and a human-readable safety net. Add rotation only if it becomes a problem.
- Codex's hook surface is less documented publicly than Claude Code's; confirm the exact JSON path before implementation. If unsupported in current Codex, ship `claude-code` and `shell` first and add `codex` in a follow-up.
- Do we want to expose the override table to power users as a runtime config? Leaning no — keeping it generator-side preserves the property that pledge classification is a build-time, reviewable artifact.
