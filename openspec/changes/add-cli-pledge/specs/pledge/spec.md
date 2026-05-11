## ADDED Requirements

### Requirement: Pledge profiles and operation taxonomy
The CLI SHALL provide three named pledge profiles — `readonly`, `read-write`, and `admin` — plus a `--allow` flag accepting one or more operation identifiers or operation groups. Every API command in `internal/api/commands.go` and every hand-written command MUST be tagged with exactly one permission tier (`read`, `write`, `admin`) derived primarily from its HTTP verb, with a maintained override table for endpoints that are semantically read despite using POST.

#### Scenario: Readonly profile allows GET endpoints
- **WHEN** a session is pledged with `atlas pledge readonly` and the user runs an API command whose underlying HTTP verb is GET
- **THEN** the command executes normally and is not logged as a block

#### Scenario: Readonly profile blocks mutating endpoints
- **WHEN** a session is pledged `readonly` and the user runs `atlas clusters delete myCluster`
- **THEN** the CLI exits non-zero before any network call, prints a message naming the operation and the active pledge, and appends a block entry to the audit log

#### Scenario: Admin profile requires explicit confirmation
- **WHEN** a user runs `atlas pledge admin` without `--yes`
- **THEN** the CLI prompts with an explicit warning that admin pledge provides no restriction and requires confirmation before persisting the pledge

#### Scenario: Override table reclassifies semantic-read POSTs
- **WHEN** an operation in `internal/api/commands.go` uses HTTP POST but is listed in the read-only override table (for example, a search endpoint)
- **THEN** that operation is classified `read` and is permitted under `readonly`

### Requirement: Session anchoring and inheritance
Pledge state SHALL be anchored to the POSIX session ID (`getsid(2)`) of the invoking process and persisted to a file under `$XDG_STATE_HOME/atlascli/pledge/<sid>.json` with mode 0600. Every `atlas-cli` invocation MUST read the pledge file matching its own SID at startup, before dispatch. Child `atlas-cli` invocations within the same shell session inherit the pledge automatically because they share the SID.

#### Scenario: Child invocation inherits pledge
- **GIVEN** the user has run `atlas pledge readonly` in their shell
- **WHEN** the user runs `bash -c 'atlas clusters delete x'` in the same terminal
- **THEN** the nested `atlas` invocation reads the same pledge file and blocks the delete

#### Scenario: New login session starts unpledged
- **WHEN** the user opens a new terminal that produces a fresh session ID
- **THEN** there is no pledge file for that SID and `atlas` runs with no pledge restrictions

#### Scenario: Windows is unsupported in v1
- **WHEN** `atlas pledge` is run on Windows
- **THEN** the command exits non-zero with a message stating that pledge requires POSIX sessions and is not supported on Windows

### Requirement: Monotonic narrowing
Within a single session, pledges SHALL only narrow. The CLI MUST reject any `atlas pledge` invocation that would widen the active permission set (move to a more permissive profile, or add operations not in the current allowlist) and refer the user to the `--allow` token flow.

#### Scenario: Re-pledging to a stricter profile succeeds
- **GIVEN** the session is pledged `read-write`
- **WHEN** the user runs `atlas pledge readonly`
- **THEN** the pledge is replaced with the stricter `readonly` profile

#### Scenario: Attempting to widen fails
- **GIVEN** the session is pledged `readonly`
- **WHEN** the user runs `atlas pledge read-write`
- **THEN** the command exits non-zero with a message explaining pledges may only be narrowed and pointing at `pledge --allow`

### Requirement: Out-of-band token approval
When a pledged command is blocked, the CLI SHALL optionally (gated by `--interactive-approval`, default on for TTYs) print a one-time request token and wait up to a configurable timeout. A second terminal MUST be able to grant approval by running `atlas pledge --allow <token>`. An approval grants the exact operation identifier captured at request time for a 5-minute TTL, scoped to the requesting SID.

#### Scenario: Approval grants the blocked operation
- **GIVEN** session A is pledged `readonly` and runs `atlas clusters delete x`
- **WHEN** the CLI prints token `T` and the user runs `atlas pledge --allow T` in session B
- **THEN** the original command proceeds, the approval is consumed (single-use), and an audit entry records SID-A, SID-B, operation, and timestamp

#### Scenario: Approving terminal must itself permit the operation
- **WHEN** session B is itself pledged below the operation's required tier and the user runs `atlas pledge --allow T` there
- **THEN** the approval is refused with a message stating the approver must have at least the requested tier

#### Scenario: Tokens expire and are single-use
- **WHEN** an approval token is older than its 5-minute TTL, or has already been consumed
- **THEN** `atlas pledge --allow T` refuses with an explicit reason and the blocked command remains blocked

#### Scenario: Approval binds to operation, not arbitrary widening
- **WHEN** session A requested approval for `deleteCluster(name=x)` and session B approves
- **THEN** only that specific operationID + parameters-hash is allowed for the 5-minute window; an unrelated `deleteCluster(name=y)` is still blocked

### Requirement: Enforcement at a single chokepoint
Pledge enforcement SHALL run at exactly one place in the dispatch path so that no command can be added that bypasses it. For generated API commands this is the API executor; for hand-written commands this is a `PreRunE` interceptor registered by `internal/cli/root/builder.go` that resolves the command's declared permission tier.

#### Scenario: Hand-written command without permission tag fails CI
- **WHEN** a developer adds a new command under `internal/cli/**` without declaring its permission tier
- **THEN** a lint or unit test fails with a message naming the command and the missing annotation

### Requirement: Audit log
Every block and every approved-by-token execution SHALL append a structured JSON line to `$XDG_STATE_HOME/atlascli/pledge/audit.log` (mode 0600). Each entry MUST include timestamp, SID, operationID, parameters hash, outcome (`blocked` | `allowed-by-token`), and approving SID when applicable.

#### Scenario: Blocked attempt is recorded
- **WHEN** any pledged session blocks a command
- **THEN** a JSON line with the fields above is appended to the audit log
