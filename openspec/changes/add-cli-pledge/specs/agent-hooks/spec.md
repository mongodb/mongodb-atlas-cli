## ADDED Requirements

### Requirement: `atlas hook install <agent>` command
The CLI SHALL provide an `atlas hook install <agent>` command supporting at minimum `claude-code`, `codex`, and `shell` as agent identifiers. The command MUST idempotently configure the named agent to invoke `atlas pledge readonly` (or a user-supplied profile via `--profile`) at the start of every session it owns. Running `install` twice with the same arguments SHALL be a no-op rather than producing duplicate entries.

#### Scenario: Claude Code install writes a SessionStart hook
- **WHEN** the user runs `atlas hook install claude-code`
- **THEN** the CLI merges a `SessionStart` hook entry into `~/.claude/settings.json` that runs `atlas pledge readonly`, preserves any pre-existing hooks and settings, and prints the resulting hook path

#### Scenario: Project-scoped install
- **WHEN** the user runs `atlas hook install claude-code --project`
- **THEN** the hook is written to `./.claude/settings.json` instead of the user-global file, and the command refuses with a clear error if not run inside a directory containing `.claude/` or with `--create-dir`

#### Scenario: Custom profile and allowlist
- **WHEN** the user runs `atlas hook install claude-code --profile read-write --allow listClusters,getProject`
- **THEN** the installed hook invokes `atlas pledge read-write --allow listClusters,getProject` instead of bare `readonly`

#### Scenario: Re-install is idempotent
- **WHEN** the user runs `atlas hook install claude-code` twice with identical arguments
- **THEN** the second run reports "already installed" and exits zero without modifying the settings file

### Requirement: `atlas hook uninstall <agent>` command
The CLI SHALL provide an `atlas hook uninstall <agent>` command that removes only the hook entries it previously installed, identified by a stable marker, and leaves all other user-managed settings untouched.

#### Scenario: Uninstall removes only Atlas-installed hooks
- **GIVEN** `~/.claude/settings.json` contains both an Atlas-installed `SessionStart` hook and a user-authored unrelated hook
- **WHEN** the user runs `atlas hook uninstall claude-code`
- **THEN** only the Atlas hook is removed; the unrelated hook and any other settings are preserved byte-for-byte

#### Scenario: Uninstall when not installed
- **WHEN** `atlas hook uninstall claude-code` is run and no Atlas hook is present
- **THEN** the command exits zero with a "nothing to remove" message

### Requirement: Generic shell hook
For `atlas hook install shell`, the CLI SHALL print a snippet the user can source from their shell rc (bash/zsh/fish), and SHALL refuse to modify shell rc files automatically without an explicit `--write <path>` flag.

#### Scenario: Shell hook print-only by default
- **WHEN** the user runs `atlas hook install shell`
- **THEN** the CLI prints a sourceable snippet to stdout and instructions for which file to add it to, but does not modify the filesystem

#### Scenario: Shell hook with explicit --write
- **WHEN** the user runs `atlas hook install shell --write ~/.zshrc`
- **THEN** the snippet is appended to `~/.zshrc` between Atlas-managed marker comments, and re-running with `--write` to the same file is a no-op

### Requirement: Settings-file merge safety
Hook installation SHALL never overwrite a settings file in place without first writing a `.bak` snapshot, and SHALL preserve unknown JSON keys, key ordering where the format permits, and trailing newlines.

#### Scenario: Backup before modify
- **WHEN** `atlas hook install claude-code` modifies an existing `~/.claude/settings.json`
- **THEN** the prior file contents are copied to `~/.claude/settings.json.bak` (mode 0600) before the new contents are written
