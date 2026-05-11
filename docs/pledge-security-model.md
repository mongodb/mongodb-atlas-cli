# Atlas CLI Pledge — Security Model

## What is a pledge?

`atlas pledge set <profile>` voluntarily restricts the current shell session to a subset of Atlas CLI operations. It is inspired by OpenBSD's `pledge(2)` syscall: a process declares what it no longer needs, and the kernel enforces it. Here, the CLI enforces it instead.

Three built-in profiles exist:

| Profile     | Allowed operations       |
|-------------|--------------------------|
| `readonly`  | GET (read-only) endpoints only |
| `read-write`| GET + mutation endpoints |
| `admin`     | All endpoints including org-level destructive operations |

## Trust boundary

The pledge is a **client-side control**. It does not talk to Atlas or change any server-side permissions. A malicious actor with access to your API key can still call the Atlas API directly.

The pledge protects you from accidental or automated mistakes — it does not protect you from adversarial actors who already have your credentials.

## Session anchoring

Pledge state is keyed to the POSIX session ID (`getsid(2)`). All processes that share a session — including child shells (`bash -c '...'`), `xargs`, and subprocesses — inherit the restriction automatically. This is why pledge is "shell-resistant": you cannot escape it by wrapping `atlas` in another shell.

A user can escape by running `setsid atlas ...`, which creates a new session with no pledge. This escape is intentional (security, not safety) but is **logged to the audit log** so it is detectable.

## Monotonic narrowing

Pledges can only be made more restrictive within a session, never more permissive. To reset to full permissions, open a new terminal.

This monotonic property prevents an escalation pattern where a script silently widens its own permissions.

## Out-of-band approval

When a pledged session attempts a blocked operation, it:

1. Writes a request file under `<state>/pledge/<sid>/requests/<token>.json`
2. Prints the token and instructions to stderr
3. If stdin is a TTY, waits up to 5 minutes for a grant

In a second terminal (which may have a different or no pledge), the approver runs:

```
atlas pledge allow <token>
```

This writes an HMAC-signed grant. The waiting process verifies the HMAC, consumes the grant, and proceeds.

**The approver's own pledge must permit the requested operation.** An approver under a readonly pledge cannot approve a write operation.

Grants are valid for 5 minutes and can only be consumed once. A forged or tampered grant is rejected via HMAC verification.

## Audit log

All blocked operations and out-of-band approvals are appended to `<state>/audit.log` as JSON lines. The log is append-only (mode 0600). Errors writing to the audit log are silently ignored so they do not block the CLI.

## Recommended workflow with Claude Code

```bash
# Install the pledge hook — runs at the start of every Claude Code session
atlas hook install claude-code --profile readonly

# When Claude needs to perform a write, it will be blocked and print a token.
# In a separate terminal, the developer approves:
atlas pledge allow <token>
```

This gives the developer explicit control over every destructive operation while letting read operations proceed automatically.

## What pledge does NOT protect against

- Direct Atlas API calls bypassing the CLI
- Another process with access to the state directory forging a grant (mitigated by HMAC)
- A user running `setsid atlas ...` (logged but not blocked)
- Windows (pledge is a no-op on Windows due to the absence of POSIX session IDs)
