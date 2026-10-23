# Atlas CLI Agent Skills PoC results

Date: 2026-10-22

## Command surface

```sh
atlas skills list
atlas skills install mongodb-schema-design --agent claude
atlas skills install --all --agent claude
atlas skills install --plugin mongodb-atlas --agent claude
atlas skills install mongodb-schema-design --agent codex
atlas skills install --all --agent codex
atlas skills install --plugin mongodb-atlas --agent codex
```

## Demo results

- `atlas skills list --output json` returned the eight standalone skills from
  `mongodb/agent-skills` at the resolved `main` snapshot and the supported
  `mongodb-atlas` plugin entry.
- Named install placed `mongodb-schema-design` under the isolated Claude home
  at `.claude/skills/mongodb-schema-design`.
- `--all` installed all eight standalone skills under a second isolated Claude
  home.
- Plugin installation was run using
  `@anthropic-ai/claude-code@2.1.281` with an isolated `HOME` and
  `CLAUDE_CONFIG_DIR`. Claude reported the `mongodb-plugins` marketplace as
  GitHub source `mongodb/agent-skills`, and `claude plugin list --json` reported
  `mongodb-atlas@mongodb-plugins`, enabled at user scope, version `1.2.1`, at
  `.claude/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1`. The plugin's MCP
  metadata was type `http` at `https://mcp.mongodb.com/`.
- With Codex CLI 0.163.0 and an isolated `HOME`/`CODEX_HOME`, `--all --agent
  codex` installed all eight standalone skills under `$CODEX_HOME/skills/`.
- A named Codex install placed `mongodb-schema-design` under
  `$CODEX_HOME/skills/mongodb-schema-design`.
- `--plugin mongodb-atlas --agent codex` added and upgraded the native
  `mongodb-plugins` marketplace, then installed `mongodb-atlas@mongodb-plugins`.
  `codex plugin list --json` reported version `1.2.1`, `installed: true`, and
  `enabled: true`; the cache was under
  `$CODEX_HOME/plugins/cache/mongodb-plugins/mongodb-atlas/1.2.1`.
- Re-running the Codex plugin install with `GIT_SSH_COMMAND=/usr/bin/false`
  returned `alreadyInstalled: true` without a marketplace refresh.
- The Codex plugin manifest declares its native app and bundled skills. It does
  not add an entry to the separate `codex mcp` registry; verify installation
  with `codex plugin list --json` or the `/plugins` browser. Standalone skills
  are installed individually under `$CODEX_HOME/skills/`.
- The stale-clone fix explicitly runs `claude plugin marketplace update
  mongodb-plugins` before looking up a missing plugin. A fresh isolated home
  with a matching marketplace pre-registered refreshed and installed the plugin
  successfully. An already-installed user plugin is idempotent after local
  source verification and does not require a remote marketplace refresh.
- The first unauthenticated GitHub API request hit the shared egress rate
  limit (`403`); listing and installs succeeded after the rate-limit reset.

All install demos used temporary homes; active Claude/Codex configuration was
not modified. OAuth sign-in and a live MCP request were not tested.

Skills commands intentionally skip Atlas telemetry because they do not load an
Atlas profile/service. The installer still uses a command-local signal context
for cancellation and rollback.

## AI discoverability

The normal successful-command post-run hook uses the Atlas SDK's AI-agent
detector. If it detects an AI caller and finds neither a standalone MongoDB
skill nor the Atlas plugin in the active Claude/Codex config, it writes a stderr
hint with standalone and plugin install commands for both agents. A per-user
cache marker suppresses later warnings; concurrent first invocations can race.
Removing the marker shows the hint again. Detection honors `CLAUDE_CONFIG_DIR`
and `CODEX_HOME`. The check also runs after `atlas skills list` and is
independent of Atlas profile/telemetry initialization.

## Validation

- Focused tests passed for skills, root, usage, and `cmd/atlas`; focused lint
  reported zero issues.
- `make gen-docs` and `make check` passed: 1,889 tests, 107 expected skips, and
  0 lint issues.
- Agent CLIs run with standard `exec.CommandContext` and installer cancellation
  uses `signal.NotifyContext`. Windows skills/root/Atlas cross-builds passed;
  `.cmd`/`.bat` launching and descendant-process cleanup are deferred.
