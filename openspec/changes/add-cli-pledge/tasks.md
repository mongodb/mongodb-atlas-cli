## 1. Permission tier on generated API commands

- [x] 1.1 Add `Permission` field (`read` | `write` | `admin`) to `shared_api.Command` in `tools/shared/api`.
- [x] 1.2 Create `tools/api-generator/permission-overrides.yaml` listing operationIDs that override the HTTP-verb default (plus the `admin` set).
- [x] 1.3 Update `tools/api-generator` to derive the tier from HTTP verb, apply overrides, and emit the field into `internal/api/commands.go`.
- [x] 1.4 Add a generator unit test asserting every operationID in `permission-overrides.yaml` is present in the generated `commands.go` (catches drift when endpoints are removed).
- [x] 1.5 Run `make gen-api-commands` and commit the regenerated `internal/api/commands.go`.

## 2. Permission tier on hand-written commands

- [x] 2.1 Define annotation keys (`atlas.permission` = `read|write|admin|local-write`) and a small helper that sets them on a Cobra command.
- [x] 2.2 Add a unit test that walks the full Cobra tree from `internal/cli/root/builder.go` and fails if any leaf command lacks an `atlas.permission` annotation, excluding an explicit meta-command allowlist (`help`, `completion`, `version`, `pledge`, `hook`).
- [x] 2.3 Annotate existing hand-written commands area by area (clusters, dbusers, accesslists, …); land as small reviewable PRs once the CI test is in place.

## 3. Pledge core (`internal/pledge`)

- [x] 3.1 Implement `Session()` returning the POSIX SID on Linux/macOS via `syscall.Getsid(0)`; return a sentinel error on Windows.
- [x] 3.2 Implement state directory resolution honoring `XDG_STATE_HOME` with documented per-OS fallbacks; ensure `0700` on dir create and `0600` on file create.
- [x] 3.3 Define the on-disk `PledgeFile` struct: `{Version, Profile, AllowedOps []string, NarrowedAt, NarrowedBy}` plus JSON marshal/unmarshal.
- [x] 3.4 Implement `Load(sid)`, `Save(sid, p)`, and `Narrow(sid, next)` enforcing subset semantics (reject widening).
- [x] 3.5 Implement `Check(p, opPermission, opID)` returning `Allow | Block` with a structured error carrying enough info for the error message and audit log.
- [x] 3.6 Implement append-only JSON-lines audit log writer with file locking.
- [x] 3.7 Implement HMAC secret bootstrap at `<state>/.hmac` (generate 32 random bytes on first use, mode 0600).
- [x] 3.8 Unit tests for narrow/widen, load/save, subset checks, audit-log append, and HMAC bootstrap idempotence.

## 4. Enforcement chokepoints

- [x] 4.1 In `internal/api/executor.go`, call `pledge.Check(...)` before building the HTTP request; on `Block`, return the structured error and append to audit log.
- [x] 4.2 In `internal/cli/root/builder.go`, register a `PersistentPreRunE` that reads `cmd.Annotations["atlas.permission"]` and calls `pledge.Check`; bypass for meta-commands.
- [x] 4.3 Add integration tests: a generated command and a hand-written command, both blocked under `readonly`, both allowed under `admin`.

## 5. `atlas pledge` command (`internal/cli/pledge`)

- [x] 5.1 Implement `atlas pledge <profile> [--allow op,op,…]` with subset-only narrowing.
- [x] 5.2 Implement `atlas pledge admin` requiring `--yes` or interactive confirmation.
- [x] 5.3 Implement `atlas pledge show` printing the active pledge for the current SID.
- [x] 5.4 Wire `pledge` into `internal/cli/root/builder.go`; tag it as meta in the annotation allowlist.
- [ ] 5.5 Docs: regenerate via `make gen-docs`.

## 6. Out-of-band `--allow <token>` approval flow

- [x] 6.1 On block, write a request file under `<state>/pledge/<sid>/requests/<token>.json` (random 16-byte token, hex-encoded) and print token + instructions to stderr.
- [x] 6.2 Implement filesystem watcher (inotify on Linux, kqueue on macOS) on `<state>/pledge/<sid>/grants/` with a polling fallback every 500 ms.
- [x] 6.3 Implement `atlas pledge --allow <token>`: read the matching request file, verify the approver's pledge permits the requested tier, write an HMAC-signed grant file, refuse if expired or already consumed.
- [x] 6.4 On the blocked side, verify HMAC signature, consume the grant (rename to `.consumed`), record `allowed-by-token` audit entry, and proceed with the original command.
- [x] 6.5 Default `--wait-for-approval=false` when stdin is not a TTY; document the flag.
- [x] 6.6 Tests covering: happy path, expired token, already-consumed token, approver under-pledged, paramsHash mismatch, forged-grant rejection.

## 7. `atlas hook install/uninstall` (`internal/cli/hook`)

- [x] 7.1 Define an `Agent` interface (`Install(opts) error`, `Uninstall() error`, `Status() State`) and register `claude-code`, `codex`, `shell` implementations.
- [x] 7.2 Implement Claude Code adapter: locate `~/.claude/settings.json` (or `./.claude/settings.json` with `--project`), merge a `SessionStart` hook entry tagged with `"_atlas_managed": true`, write atomically with `.bak`.
- [x] 7.3 Implement Codex adapter (gated behind a feature flag pending Codex hook documentation; if blocked, leave a stub returning a clear "not yet supported, see issue X" error).
- [x] 7.4 Implement shell adapter: print sourceable snippet wrapped in `# >>> atlas pledge >>>` markers; write to a file only with `--write <path>`; in-place idempotent block rewrite between markers.
- [x] 7.5 Implement `atlas hook uninstall <agent>` removing only the marker-tagged entries.
- [x] 7.6 Golden-file tests for each adapter: install on empty file, install on populated file, idempotent re-install, uninstall preserves unrelated entries.

## 8. Edge cases and hardening

- [x] 8.1 Lazily garbage-collect stale `<sid>.json` files on pledge write (skip files for SIDs with no live process).
- [x] 8.2 Detect SID lineage break (`getsid(0) != getsid(getppid())`) at startup and log a one-line note to the audit log.
- [x] 8.3 Clear, actionable error messages on block: name the operation, the active pledge, and how to approve via `--allow`.

## 9. Documentation and release

- [x] 9.1 Add a security-model page under `docs/` covering the trust boundary, what pledge does and does not protect against, and the recommended workflow with Claude Code.
- [x] 9.2 Update `CONTRIBUTING.md` with the "every command must carry a permission tier" rule and how to add overrides.
- [x] 9.3 Run `make check` clean; capture screenshots / asciinema for the PR.
- [ ] 9.4 Open PR per the `.agents/skills/atlas-cli-open-pr/` workflow.
