# MongoDB MCP Server — Tool Catalog

This document shows the full proposed tool list across all layers, with MCP safety annotations and implementation status against the current codebase.

**Status key:**
- ✅ **EXISTS** — tool already exists in the codebase with this name
- 🔄 **RENAME** — equivalent tool exists under a different name; needs renaming
- ⚙️ **EXTEND** — related tool exists but needs significant extension (e.g. polling, richer output)
- 🆕 **NEW** — no equivalent; needs to be built

---

## How the codebase derives annotations today

The base `ToolBase` class in `src/tools/tool.ts` maps `operationType` to `ToolAnnotations` automatically:

```
"read" | "metadata" | "connect"  →  readOnlyHint: true,  destructiveHint: false
"delete"                          →  readOnlyHint: false, destructiveHint: true
"create" | "update"               →  readOnlyHint: false, destructiveHint: false
```

`idempotentHint` and `openWorldHint` are not set today — they take spec defaults (`false` and `true` respectively). The proposal is to add explicit `idempotentHint` support to `ToolBase` so workflow tools and PATCH-backed raw tools can declare it.

---

## Layer 1 — Workflow tools (always visible, 25 tools)

### Clusters

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-cluster` | ⚙️ EXTEND | `atlas-create-free-cluster` | create | — | — | — |
| `atlas-delete-cluster` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-pause-cluster` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-resume-cluster` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-scale-cluster` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-get-cluster-status` | ⚙️ EXTEND | `atlas-inspect-cluster` | read | ✓ | — | ✓ |

Notes:
- `atlas-create-free-cluster` only creates M0 free tier clusters. `atlas-create-cluster` needs to support paid tiers with provider/region/tier params and poll until `stateName == IDLE`.
- `atlas-inspect-cluster` exists but returns raw cluster metadata. `atlas-get-cluster-status` should return a simplified, agent-friendly shape (state, connection string, tier) and is the rename target.

### Projects

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-project` | ✅ EXISTS | `atlas-create-project` | create | — | — | — |
| `atlas-list-projects` | ✅ EXISTS | `atlas-list-projects` | read | ✓ | — | ✓ |
| `atlas-get-project` | 🆕 NEW | — | read | ✓ | — | ✓ |

### Database Users

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-database-user` | 🔄 RENAME | `atlas-create-db-user` | create | — | — | — |
| `atlas-delete-database-user` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-database-users` | 🔄 RENAME | `atlas-list-db-users` | read | ✓ | — | ✓ |

Note: `atlas-create-db-user` and `atlas-list-db-users` use abbreviated names. Proposal is to rename to the full `database-user` form for consistency.

### Network Access

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-allow-ip-access` | ⚙️ EXTEND | `atlas-create-access-list` | create | — | — | ✓ |
| `atlas-remove-ip-access` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-ip-access` | 🔄 RENAME | `atlas-inspect-access-list` | read | ✓ | — | ✓ |

Note: `atlas-create-access-list` is the closest existing tool. `atlas-allow-ip-access` is a more agent-friendly name (the verb describes intent). `atlas-inspect-access-list` → `atlas-list-ip-access` for naming consistency.

### Backups

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-list-snapshots` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-snapshot` | 🆕 NEW | — | create | — | — | — |
| `atlas-restore-snapshot` | 🆕 NEW | — | delete* | — | ✓ | — |

*`atlas-restore-snapshot` uses `operationType: "delete"` so it gets `destructiveHint: true` automatically, which is correct — restoring overwrites live data.

### Search

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-search-index` | 🆕 NEW | — | create | — | — | — |
| `atlas-list-search-indexes` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-delete-search-index` | 🆕 NEW | — | delete | — | ✓ | — |

### Streams

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-stream-instance` | ⚙️ EXTEND | `atlas-streams-build` | create | — | — | — |
| `atlas-list-stream-instances` | ⚙️ EXTEND | `atlas-streams-discover` | read | ✓ | — | ✓ |

Note: `atlas-streams-build` and `atlas-streams-discover` cover streams creation and discovery. The rename/extension aligns them with the `atlas-{verb}-{resource}` naming pattern.

### Serverless

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-create-serverless-instance` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-serverless-instance` | 🆕 NEW | — | delete | — | ✓ | — |

### Discovery meta-tool

| Tool | Status | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|
| `atlas-find-tools` | 🆕 NEW | read | ✓ | — | ✓ |

---

## Layer 2 — Raw tools by domain (loaded via `atlas-find-tools`)

### Domain: `clusters` (12 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-list-clusters` | ✅ EXISTS | `atlas-list-clusters` | read | ✓ | — | ✓ |
| `atlas-get-cluster` | ⚙️ EXTEND | `atlas-inspect-cluster` | read | ✓ | — | ✓ |
| `atlas-update-cluster` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-get-cluster-advanced-settings` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-cluster-advanced-settings` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-enable-termination-protection` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-disable-termination-protection` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-get-connection-strings` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-upgrade-cluster-version` | 🆕 NEW | — | update | — | — | — |
| `atlas-test-failover` | 🆕 NEW | — | update | — | — | — |
| `atlas-list-cloud-provider-regions` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-list-cluster-instance-sizes` | 🆕 NEW | — | read | ✓ | — | ✓ |

### Domain: `projects` (10 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-delete-project` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-get-project-settings` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-project-settings` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-list-project-teams` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-add-team-to-project` | 🆕 NEW | — | create | — | — | — |
| `atlas-remove-team-from-project` | 🆕 NEW | — | delete | — | — | — |
| `atlas-list-alert-configurations` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-alert-configuration` | 🆕 NEW | — | create | — | — | — |
| `atlas-update-alert-configuration` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-delete-alert-configuration` | 🆕 NEW | — | delete | — | ✓ | — |

### Domain: `users` (8 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-get-database-user` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-database-user` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-list-custom-db-roles` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-custom-db-role` | 🆕 NEW | — | create | — | — | — |
| `atlas-update-custom-db-role` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-delete-custom-db-role` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-project-users` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-remove-project-user` | 🆕 NEW | — | delete | — | — | — |

Note: `atlas-remove-project-user` removes access from a project, not the Atlas account itself, so `destructiveHint` should be `false` despite using `delete` operationType. This is an overlay override case.

### Domain: `networking` (10 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-list-private-endpoint-services` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-get-private-endpoint-service` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-private-endpoint-service` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-private-endpoint-service` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-network-peerings` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-network-peering` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-network-peering` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-network-containers` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-network-container` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-network-container` | 🆕 NEW | — | delete | — | ✓ | — |

### Domain: `backups` (10 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-get-backup-schedule` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-backup-schedule` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-get-snapshot` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-delete-snapshot` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-restore-jobs` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-get-restore-job` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-list-export-buckets` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-export-bucket` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-export-bucket` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-get-backup-compliance-policy` | 🆕 NEW | — | read | ✓ | — | ✓ |

### Domain: `search` (7 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-get-search-index` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-search-index` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-get-search-deployment` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-search-deployment` | 🆕 NEW | — | create | — | — | — |
| `atlas-update-search-deployment` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-delete-search-deployment` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-list-search-analyzers` | 🆕 NEW | — | read | ✓ | — | ✓ |

### Domain: `monitoring` (8 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-list-alerts` | ✅ EXISTS | `atlas-list-alerts` | read | ✓ | — | ✓ |
| `atlas-get-alert` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-acknowledge-alert` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-list-alert-configurations` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-get-measurement` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-list-namespace-metrics` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-get-slow-query-log` | ⚙️ EXTEND | `atlas-get-performance-advisor` | read | ✓ | — | ✓ |
| `atlas-get-suggested-indexes` | ⚙️ EXTEND | `atlas-get-performance-advisor` | read | ✓ | — | ✓ |

Note: The existing `atlas-get-performance-advisor` bundles both slow query log and index suggestions into one tool. The proposal splits this into two focused tools for clearer agent invocation, each targeting the relevant API endpoint.

### Domain: `advanced` (8 tools)

| Tool | Status | Existing name | operationType | readOnly | destructive | idempotent |
|---|---|---|---|---|---|---|
| `atlas-list-organizations` | 🔄 RENAME | `atlas-list-orgs` | read | ✓ | — | ✓ |
| `atlas-get-organization` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-list-project-api-keys` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-create-project-api-key` | 🆕 NEW | — | create | — | — | — |
| `atlas-delete-project-api-key` | 🆕 NEW | — | delete | — | ✓ | — |
| `atlas-get-maintenance-window` | 🆕 NEW | — | read | ✓ | — | ✓ |
| `atlas-update-maintenance-window` | 🆕 NEW | — | update | — | — | ✓ |
| `atlas-defer-maintenance-window` | 🆕 NEW | — | update | — | — | — |

---

## Existing tools not in this catalog

These tools exist today but sit outside the proposed auto-gen scope — either because they provide local dev infrastructure or connect the MCP session to an actual cluster.

| Existing tool | Category | Notes |
|---|---|---|
| `atlas-connect-cluster` | atlas | Connects the MCP session to an Atlas cluster for MongoDB tools. Retained as-is; not part of the auto-gen framework. |
| `atlas-streams-manage` | atlas | Manages running stream processor state (start/stop/drop). Retained; streams tooling may be partially auto-gen'd in future. |
| `atlas-streams-teardown` | atlas | Tears down a stream processing pipeline. Retained. |
| `atlas-local-create-deployment` | atlas-local | Local Atlas via CLI. Out of scope for API auto-gen. |
| `atlas-local-connect-deployment` | atlas-local | Out of scope. |
| `atlas-local-list-deployments` | atlas-local | Out of scope. |
| `atlas-local-delete-deployment` | atlas-local | Out of scope. |
| `connect` | mongodb | Direct MongoDB connection. Out of scope. |
| `switch-connection` | mongodb | Out of scope. |
| All `mongodb/*` tools | mongodb | CRUD on MongoDB instances, not Atlas API. Out of scope for auto-gen. |
| `search-knowledge` | assistant | MongoDB docs search. Out of scope. |
| `list-knowledge-sources` | assistant | MongoDB docs. Out of scope. |

---

## Summary

| Status | Count |
|---|---|
| ✅ EXISTS (exact match) | 5 |
| 🔄 RENAME (same logic, different name) | 5 |
| ⚙️ EXTEND (exists but needs significant work) | 7 |
| 🆕 NEW | 82 |
| **Total proposed** | **99** |

The 5 exact matches: `atlas-create-project`, `atlas-list-projects`, `atlas-list-clusters`, `atlas-list-alerts`, and `atlas-list-clusters` (raw domain duplicate of the workflow-layer read).

The 5 renames touch `db-user` → `database-user`, `orgs` → `organizations`, and `inspect-*` → `list-*`/`get-*`. These are non-breaking if handled with a deprecation alias.

The 7 extensions are all workflow tools that need polling added on top of existing single-operation tools (`atlas-create-free-cluster` → paid tiers + wait, `atlas-inspect-cluster` → simplified output shape, access list + performance advisor splits).

---

## Annotation override cases

The `operationType` → `ToolAnnotations` mapping in `ToolBase` handles the common cases. Two tools need explicit overrides because the automatic mapping is wrong:

**`atlas-remove-project-user` (operationType: `delete`)** — automatically gets `destructiveHint: true`, but removing a user from a project does not delete their Atlas account or any data. Should be `destructiveHint: false`. Requires either a new `operationType` value (`"revoke"`) or an explicit annotation override in the tool class.

**`atlas-restore-snapshot` (operationType: `delete`\*)** — correctly gets `destructiveHint: true`. The asterisk is because this operation is modeled as `delete`-class purely for annotation purposes; internally it calls a POST endpoint. The operationType should be chosen for what it means to the user, not what HTTP verb is used.

A small addition to `ToolBase` that lets individual tool classes override specific annotation fields would handle both cases without changing the automatic mapping for everything else:

```typescript
// In ToolBase — optional override per tool
protected annotationOverrides(): Partial<ToolAnnotations> {
    return {};
}

public get annotations(): ToolAnnotations {
    const base = this.baseAnnotations();  // existing switch block
    return { ...base, ...this.annotationOverrides() };
}
```

```typescript
// In RemoveProjectUserTool
protected annotationOverrides(): Partial<ToolAnnotations> {
    return { destructiveHint: false };
}
```
