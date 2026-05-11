## Why

AI coding agents (Claude Code, Codex, etc.) increasingly invoke `atlas-cli` on the user's behalf. The CLI runs with the user's full Atlas credentials, so a confused or compromised agent can delete clusters, rotate keys, or change billing settings just as easily as it can list projects. Today the only mitigation is "use a less-privileged API key", which is coarse, requires pre-provisioning, and doesn't compose with a long-running interactive shell where the human and the agent share the same terminal.

OpenBSD's `pledge(2)` solves a similar problem at the syscall layer: a process voluntarily drops capabilities it promises not to need, and any later attempt to use them is fatal. We want the same shape for `atlas-cli` at the API-surface layer: a session-scoped, monotonically-shrinking allowlist of operations the CLI is permitted to perform, enforced inside the binary itself, that child invocations of `atlas-cli` inherit automatically.

## What Changes

- Add a new top-level command `atlas pledge` that restricts the current shell session's `atlas-cli` invocations to a subset of operations. Profiles: `readonly` (HTTP GET only), `read-write` (no destructive verbs; warned as elevated), `admin` (no restriction; explicit opt-in with warning), plus an explicit `--allow <operationID|group>` list for finer-grained scoping.
- Pledge state is anchored to the OS session (POSIX session ID on Linux/macOS) and persisted to `$XDG_STATE_HOME/atlascli/pledge/<sid>.json`. Every `atlas-cli` invocation checks this file at startup before dispatching the requested command. Child invocations inherit by sharing the SID; new login sessions start unpledged.
- Pledges are monotonic within a session: a session can only narrow its pledge, never widen it. Loosening requires either ending the session or the `--allow` token flow below.
- Add `atlas hook install <agent>` and `atlas hook uninstall <agent>` for `claude-code`, `codex`, and a generic `shell` target. For Claude Code, the installer writes a `SessionStart` hook entry to `~/.claude/settings.json` (or project `.claude/settings.json` with `--project`) that runs `atlas pledge readonly` before the agent runs commands. Existing settings are merged, not clobbered.
- Add a deferred-but-specified out-of-band approval flow: when a pledged session attempts a blocked operation, the CLI prints a one-time request token and blocks for up to N seconds. In a separate terminal the user runs `atlas pledge --allow <token>`, which grants that exact operation (bound by operationID + parameters hash) for a 5-minute window. The approving terminal must itself not be pledged below `read-write`.
- Mark every entry in `internal/api/commands.go` with an effective permission tier derived from the HTTP verb (GET → read; POST/PUT/PATCH/DELETE → write), and add a manual override table for endpoints that are semantically read-only despite using POST (search/aggregation style). Non-generated commands in `internal/cli/**` get an explicit `Permissions` annotation on their `*Builder()` registration.
- Every blocked attempt is logged to `$XDG_STATE_HOME/atlascli/pledge/audit.log` with timestamp, SID, operationID, and outcome (blocked / allowed-by-token).

This is **not** a server-side authorization mechanism: any user with valid credentials can bypass pledge by calling the Atlas API directly. Pledge is a guardrail against confused-deputy mistakes inside `atlas-cli`, comparable to `sudo`'s NOPASSWD timeouts — useful precisely because the credential boundary is wider than the trust boundary in agent workflows.

## Capabilities

### New Capabilities
- `pledge`: Session-scoped, monotonically-shrinking allowlist of `atlas-cli` operations enforced inside the binary, with out-of-band token-based widening for individual blocked calls.
- `agent-hooks`: `atlas hook install/uninstall` integration for AI coding agents (Claude Code, Codex) and generic shells, with idempotent settings-file merging.

### Modified Capabilities
<!-- None: no existing OpenSpec specs in this repo. -->

## Impact

- **New code**: `internal/cli/pledge/`, `internal/cli/hook/`, `internal/pledge/` (enforcement core + session resolver + audit log + token broker).
- **Generated code**: `tools/api-generator/` and `internal/api/commands.go` must surface the HTTP verb (already present) plus the read-only override table; `make gen-api-commands` must round-trip the override list.
- **Wired into**: `internal/cli/root/builder.go` (register `pledge` and `hook`), and the API command executor in `internal/api/executor.go` (enforcement hook before dispatch). Hand-written commands gain a `Permissions()` method on their `*Opts` struct or a registration-time annotation.
- **Filesystem**: new state directory under `$XDG_STATE_HOME/atlascli/pledge/` (fallback `~/.local/state/atlascli/pledge/` on Linux, `~/Library/Application Support/atlascli/pledge/` on macOS). Files are mode 0600.
- **No new dependencies** expected; uses stdlib `syscall` for `getsid(2)` on Unix. Windows is explicitly out of scope for v1 (no POSIX session ID); the command exits with a clear "not supported" message.
- **Telemetry**: pledge enforcement events feed the existing telemetry pipeline behind the standard opt-out, mirroring how other CLI events are reported.
- **Docs**: regenerated via `make gen-docs`. New section in `docs/` describing the security model and explicit non-goals (not a credential boundary, not a sandbox).
