# Proposal: Auto-generation Framework for the MongoDB MCP Server

**Status:** Draft  
**Scope:** Atlas Admin API tool surface for the MongoDB MCP Server  
**References:** `docs/autogen-framework-design.md`, `docs/mcp-tool-generation-alternatives.md`, `docs/mcp-tool-design-research.md`

---

## Problem

The MongoDB MCP Server needs to expose Atlas Admin API operations as MCP tools. The Atlas Admin API has roughly 700 operations. Without a generation pipeline, that surface requires hand-maintained tool definitions that will drift the moment the API changes.

The Atlas CLI already solved this problem for CLI commands. The proposal is to extend that solution to the MCP server — reusing the spec, the overlay mechanism, and the executor, while adding a new generator output and a tool surface designed for agents rather than humans.

This is not a port of the CLI to MCP. An agent cannot browse a hierarchy and tab-complete through 700 operations. The tool surface must be designed around user intents, not API coverage, and the generation pipeline must enforce that.

---

## Proposed architecture

```
tools/internal/specs/spec.yaml          (upstream Atlas API spec, never edited)
          │
          ▼
    apply-overlay                        (existing, unchanged)
          │
          ▼
tools/internal/specs/spec-with-overlays.yaml
          │
          ├──► api-generator --output-type commands
          │         └──► internal/api/commands.go          (CLI, unchanged)
          │
          └──► api-generator --output-type mcp-tools
                    └──► internal/api/mcp_tools_gen.go     (MCP, new)
                              │
                              ▼
                    MongoDB MCP Server
                    (imports mcp_tools_gen.go or reads equivalent)
```

The generator gains a new output mode. Everything upstream of it — the spec, the overlays, the `apply-overlay` binary, the `Command` data model, the watcher validation, the executor — is unchanged. The MCP server imports the generated output and wires it into tool registration at startup.

If the MCP server is in a separate repository, the generated file is published as a Go module or copied in CI, the same way generated protobuf files are handled in practice.

---

## Tool surface design

The server exposes three layers. Each layer has a different purpose and a different generation strategy.

### Layer 1 — Workflow tools (primary surface)

Workflow tools are the tools an agent reaches for first. Each one maps to a user intent, not an API operation. The input schema is small (4–8 parameters), the description is written for agent consumption, and the implementation sequences multiple API calls internally.

Target: **25–35 workflow tools** at launch, covering the top intents across Atlas domains.

Initial workflow set, organized by domain:

**Clusters**
* `atlas_create_cluster` — Create a new cluster and wait until it is ready. Returns connection strings. Internally: `createCluster` → poll `getGroupCluster` until `stateName == IDLE`.
* `atlas_delete_cluster` — Delete a cluster and wait until it is removed. Internally: `deleteCluster` → poll `getGroupCluster` until 404.
* `atlas_pause_cluster` — Pause a running cluster. Internally: `updateCluster` with `paused: true` → poll until `stateName == IDLE`.
* `atlas_resume_cluster` — Resume a paused cluster. Internally: `updateCluster` with `paused: false` → poll until `stateName == IDLE`.
* `atlas_scale_cluster` — Change instance size. Internally: `updateCluster` with new `instanceSize` → poll until `stateName == IDLE`.
* `atlas_get_cluster_status` — Return current state, tier, and connection strings for a cluster. Internally: `getGroupCluster`.

**Projects**
* `atlas_create_project` — Create a new project. Returns project ID. Internally: `createProject`.
* `atlas_list_projects` — List all projects the caller can access. Internally: `listProjects`.
* `atlas_get_project` — Get details for a specific project by ID or name. Internally: `getProject` or `getProjectByName`.

**Database Users**
* `atlas_create_database_user` — Create a database user with specified roles. Internally: `createDatabaseUser`.
* `atlas_delete_database_user` — Remove a database user. Internally: `deleteDatabaseUser`.
* `atlas_list_database_users` — List all database users in a project. Internally: `listDatabaseUsers`.

**Network Access**
* `atlas_allow_ip_access` — Add an IP address or CIDR block to the project IP access list. Internally: `createProjectIpAccessList`.
* `atlas_remove_ip_access` — Remove an IP address from the access list. Internally: `deleteProjectIpAccessList`.
* `atlas_list_ip_access` — List current IP access list entries. Internally: `listProjectIpAccessLists`.

**Backups**
* `atlas_list_snapshots` — List available snapshots for a cluster. Internally: `listReplicaSetBackups`.
* `atlas_create_snapshot` — Trigger an on-demand snapshot. Internally: `takeSnapshot` → poll until snapshot state is `COMPLETED`.
* `atlas_restore_snapshot` — Restore a snapshot to a cluster. Internally: `createBackupRestoreJob` → poll until job state is `COMPLETED`.

**Search**
* `atlas_create_search_index` — Create an Atlas Search index on a collection. Internally: `createAtlasSearchIndex` → poll until index status is `READY`.
* `atlas_list_search_indexes` — List all search indexes on a cluster. Internally: `listAtlasSearchIndexes`.
* `atlas_delete_search_index` — Delete a search index. Internally: `deleteAtlasSearchIndex`.

**Streams (Atlas Stream Processing)**
* `atlas_create_stream_instance` — Create a stream processing instance. Internally: `createStreamInstance`.
* `atlas_list_stream_instances` — List stream instances in a project. Internally: `listStreamInstances`.

**Serverless**
* `atlas_create_serverless_instance` — Create a serverless instance. Internally: `createServerlessInstance` → poll until state is `IDLE`.
* `atlas_delete_serverless_instance` — Delete a serverless instance. Internally: `deleteServerlessInstance`.

These are the starting 25. More are added over time based on observed agent usage patterns (see the evolution strategy section).

### Layer 2 — Curated raw tools (fallback surface)

Raw tools expose individual API operations for cases no workflow covers. They use API-shaped descriptions but have description overrides in overlays to add agent-centric guidance.

Target: **50–70 raw tools** covering the operations agents actually need.

Selection criteria for inclusion:
* Mutations (POST/PATCH/PUT/DELETE) that are not covered by a workflow and have clear standalone value.
* Read operations that are parameterized enough that they cannot be meaningfully expressed as MCP Resources (search queries, filtered lists with many parameters).
* Operations frequently requested in Atlas CLI telemetry.

Explicit exclusions:
* Pure audit/compliance endpoints (billing, invoice details, audit log configuration).
* Admin-only operations requiring Organization Owner or higher that an agent is unlikely to have.
* Deprecated operations and operations with only sunset versions.
* Any operation that returns binary data (gzip dumps, etc.).
* Metrics and monitoring endpoints — these are best served via Resources or a dedicated observability tool.

Tool names for raw tools follow `{service}_{verb}_{noun}` using the operation's `x-xgen-atlascli: operationId` override where it exists, falling back to the camelCase `operationId`. The `x-mcp: tool-name` overlay key overrides both for agent-specific clarity.

### Layer 3 — Read resources

GET operations that return a single resource by ID or a simple list with minimal parameters become MCP Resources. They are accessible via URI template and do not consume tool slots.

URI scheme: `atlas://{resourceType}/{...pathParams}`

Examples:
```
atlas://clusters/{groupId}/{clusterName}
atlas://projects/{groupId}
atlas://database-users/{groupId}/{databaseName}/{username}
atlas://search-indexes/{groupId}/{clusterName}/{indexId}
```

The generator emits resource templates for all GET operations marked with `x-mcp: expose-as: resource` in the overlay. The default is to not emit resources — resource designation is opt-in because GET operations with many query parameters are poor resources and are better excluded or expressed as raw tools.

---

## Progressive tool discovery

With 25–35 workflow tools and 50–70 raw tools, the total is 75–105 tools. That exceeds the ~30-tool threshold where model performance begins to degrade noticeably.

The solution is session-scoped tool discovery, following Amazon Prime Video's pattern.

**Default visible tools** when a session starts:
* All 25–35 workflow tools (always visible — these are curated and small enough).
* One meta-tool: `atlas_find_tools`.

`atlas_find_tools` accepts a `domain` parameter (an enum) and returns the raw tools for that domain. After the call, the server sends a tool list change notification. The client re-fetches the tool list, and the session now has workflow tools + the requested domain's raw tools in context.

```
domain options:
  clusters     → raw cluster operations (resize, upgrade, modify, etc.)
  projects     → project settings, teams, alerts
  users        → database users, custom roles, LDAP
  networking   → VPC peering, private endpoints, network containers
  backups      → backup schedules, export, compliance policies
  search       → search index management, analyzers
  monitoring   → metrics, logs, performance advisor
  advanced     → everything else
```

An agent focused on cluster management sees workflow tools + ~10 cluster raw tools. It never loads monitoring tools, project audit tools, or backup tools unless explicitly switching domains. Context stays under 40 tools in all practical cases.

**Implementation notes:**
* `atlas_find_tools` is always included in generation output.
* Domain → tool mapping is static (generated), not dynamic. The server does not need to query anything to answer `atlas_find_tools`.
* Session state for "which domains are loaded" lives in the MCP session. The 2026 MCP roadmap is evolving stateful session semantics — the implementation should be designed to work statelessly if session tracking becomes unreliable (i.e., always-available workflow tools must function without having called `atlas_find_tools` first).

---

## Generator changes

### New `--output-type mcp-tools` flag

The `api-generator` binary gains a third output type. The existing `commands` and `metadata` modes are unchanged.

In `mcp-tools` mode, the generator produces a Go source file containing:
1. A `MCPWorkflowTools` slice — workflow tool definitions with step sequences.
2. A `MCPRawTools` slice — single-operation tool definitions, filtered by `x-mcp: skip`.
3. A `MCPResources` slice — resource templates, populated by `x-mcp: expose-as: resource`.
4. A `MCPToolDomains` map — the domain → raw tool name mapping for `atlas_find_tools`.

The file structure mirrors `commands.go`: a static Go variable, committed to version control, formatted with `go/format`, with a `// Code generated` header.

### New template: `mcp_tools.go.tmpl`

Added alongside `commands.go.tmpl`. Receives the same `GroupedAndSortedCommands` data plus workflow definitions parsed from `x-mcp-workflows`. Emits MCP tool registration structs rather than Cobra commands.

### Workflow resolver

The generator gains a workflow resolution pass that runs before template execution:

1. Reads `x-mcp-workflows` from the overlaid spec.
2. For each workflow step that references an `operationId`, resolves it to the corresponding `Command` struct.
3. Validates that all referenced operation IDs exist, all referenced parameters exist on those operations, and no required parameters are unaccounted for in the step's `params` mapping.
4. Emits validation errors that fail the build — misconfigured workflow definitions are caught at `make gen-api-commands` time, not at runtime.

This validation follows the same pattern as watcher validation in the existing generator.

### Tool name generation

MCP tool names are generated from the `operationId` (or `x-xgen-atlascli: operationId` override if set) using this priority:

1. `x-mcp: tool-name` (explicit override in overlay)
2. `x-xgen-atlascli: short-operation-id` prefixed with `atlas_`
3. `operationId` in `snake_case` prefixed with `atlas_`

All generated tool names are validated for uniqueness across the full tool set. Collisions fail the build.

---

## New overlay extensions

### Per-operation extensions (`x-mcp`)

```yaml
x-mcp:
  skip: true                    # exclude from MCP entirely
  expose-as: resource           # emit as MCP Resource, not Tool
  tool-name: "atlas_my_name"    # override generated tool name
  tool-description: |           # agent-centric description (overrides spec description)
    Use this to ... Requires ...
  domain: networking            # assign to a discovery domain (raw tools only)
```

### Top-level workflow definitions (`x-mcp-workflows`)

These live at the document root, not on individual operations:

```yaml
x-mcp-workflows:
  <workflowId>:
    tool-name: "atlas_create_cluster"
    description: |
      Creates a new Atlas cluster and waits until it is ready to accept connections.
      Use this when the user wants to provision a new database cluster in a project.
      Requires: project ID, cluster name, cloud provider, and region.
    input:
      groupId:
        type: string
        description: "The project ID. Use atlas_list_projects to find it."
        required: true
      clusterName:
        type: string
        description: "Name for the new cluster. Must be unique within the project."
        required: true
      provider:
        type: string
        enum: [AWS, GCP, AZURE]
        description: "Cloud provider to deploy on."
        required: true
      region:
        type: string
        description: "Provider region code (e.g. US_EAST_1 for AWS, us-east1 for GCP)."
        required: true
      tier:
        type: string
        description: "Instance size tier. Defaults to M10."
        default: "M10"
    steps:
      - id: create
        operationId: createCluster
        version: "2024-08-05"
        body:
          name: "{{ .Input.clusterName }}"
          clusterType: REPLICASET
          replicationSpecs:
            - zoneName: Zone 1
              regionConfigs:
                - providerName: "{{ .Input.provider }}"
                  regionName: "{{ .Input.region }}"
                  electableSpecs:
                    instanceSize: "{{ .Input.tier }}"
                    nodeCount: 3
      - id: wait
        type: poll
        operationId: getGroupCluster
        version: "2024-08-05"
        params:
          groupId: "{{ .Input.groupId }}"
          clusterName: "{{ .Steps.create.response.name }}"
        until:
          path: $.stateName
          values: [IDLE]
        timeout: 1800
        interval: 10
    output:
      clusterName: "{{ .Steps.wait.response.name }}"
      connectionString: "{{ .Steps.wait.response.connectionStrings.standardSrv }}"
      state: "{{ .Steps.wait.response.stateName }}"
    on-error:
      partial-completion: |
        The cluster was created but did not reach IDLE state within the timeout.
        The cluster may still be provisioning. Use atlas_get_cluster_status to check.
```

The template DSL uses Go template syntax (`{{ .Input.x }}`, `{{ .Steps.id.response.x }}`). The generator validates that all referenced fields exist in the input schema or step outputs at generation time.

---

## Overlay file strategy

New overlays live in `tools/internal/specs/overlays/` alongside the existing CLI overlays, following the same naming convention and sort order.

Files to create:

```
overlays/
  (existing CLI overlays)
  mcp_skip_audit.yaml           — skip audit/billing/compliance endpoints
  mcp_skip_deprecated.yaml      — skip endpoints with only deprecated versions
  mcp_skip_binary.yaml          — skip endpoints returning non-JSON content
  mcp_domains.yaml              — assign domain tags to raw operations
  mcp_tool_descriptions.yaml    — agent-centric description overrides
  mcp_resources.yaml            — mark GET operations as Resources
  mcp_workflows.yaml            — workflow definitions (the primary authorship surface)
```

The `mcp_workflows.yaml` overlay is the most important file and the primary artifact that requires human authorship. Everything else is mechanical. This file is what the engineering team maintains as the Atlas API evolves.

---

## Runtime tool handler

In the MCP server, tool invocation works as follows:

**For raw tools:**
1. Receive tool call with arguments.
2. Look up the `Command` definition from the generated `MCPRawTools` slice.
3. Map tool arguments to `CommandRequest.Parameters` (flat key → `[]string`).
4. Select the pinned version (newest stable, determined at generation time).
5. Call `executor.ExecuteCommand()`.
6. Return JSON response body or structured error.

**For workflow tools:**
1. Receive tool call with arguments.
2. Look up the workflow definition from `MCPWorkflowTools`.
3. Execute steps sequentially, passing data between steps via the template DSL resolver.
4. For `poll` steps, loop with configurable interval and timeout.
5. Return the output projection from the final step.
6. On mid-workflow error, return the structured `on-error` message from the workflow definition.

**Authentication** is handled at the executor level, identical to how the CLI does it. The MCP server gets an HTTP client wired with Atlas credentials (API key or OAuth token). The tool handler never handles credentials directly.

**Version pinning** — MCP tools do not expose `--version`. The generator pins each raw tool to its newest non-deprecated stable version at generation time. Workflow tools pin versions per-step in the workflow definition. Version selection is a generation-time decision, not a runtime decision.

---

## Implementation phases

### Phase 1 — Generator foundation (2 weeks)

Deliverables:
* `--output-type mcp-tools` flag on `api-generator`.
* New `mcp_tools.go.tmpl` template that emits raw tool definitions (no workflows yet).
* `x-mcp: skip`, `x-mcp: domain`, and `x-mcp: tool-name` extension support.
* Initial skip overlays (`mcp_skip_audit.yaml`, `mcp_skip_deprecated.yaml`, `mcp_skip_binary.yaml`).
* Domain assignment overlay (`mcp_domains.yaml`).
* Generated `internal/api/mcp_tools_gen.go` committed.
* Unit tests for tool name generation, skip logic, and domain assignment.

At the end of Phase 1, the MCP server has ~300–400 auto-generated raw tool definitions (before any description curation). These are not ready for agent use but establish the generation pipeline.

### Phase 2 — Description curation and tool reduction (2 weeks)

Deliverables:
* `x-mcp: tool-description` overlay support.
* `mcp_tool_descriptions.yaml` overlay with agent-centric descriptions for the 50–70 retained raw tools.
* All operations not in the retained set marked `x-mcp: skip: true`.
* `atlas_find_tools` meta-tool, statically generated from the domain map.
* Progressive discovery in the MCP server (session → domain mapping, tool list change notifications).
* Integration test: agent can call `atlas_find_tools`, receive domain tools, and successfully use them.

At the end of Phase 2, the raw tool surface is agent-usable and the discovery mechanism is working.

### Phase 3 — Workflow tools (3 weeks)

Deliverables:
* `x-mcp-workflows` overlay parsing and resolution in the generator.
* Workflow step validation (operation ID exists, parameters exist, required params covered).
* Workflow tool emission in `mcp_tools.go.tmpl`.
* Workflow executor in the MCP server (step sequencing, poll loops, data flow resolution, timeout/error handling).
* `mcp_workflows.yaml` with the initial 25-tool workflow set (see Layer 1 above).
* Integration tests per workflow: happy path + poll timeout + mid-workflow error.

At the end of Phase 3, the server has a complete three-layer tool surface.

### Phase 4 — Resources and polish (1 week)

Deliverables:
* `x-mcp: expose-as: resource` support in generator.
* `mcp_resources.yaml` overlay with initial resource designations.
* Resource handler in the MCP server.
* End-to-end agent evaluation: run a representative set of Atlas management tasks and measure tool selection accuracy and task completion rate.
* Documentation for overlay authors (how to add a new workflow, how to adjust domain assignments).

---

## What this proposal explicitly does not cover

**Telemetry and usage instrumentation.** The MCP server needs to track which tools get called, which fail, and which get retried, in order to drive the workflow evolution strategy. That is a separate design.

**Authentication flows.** The proposal assumes the MCP server already handles credential acquisition. It does not propose changing how the server authenticates to Atlas.

**Multi-project context.** The current design requires `groupId` on most calls. A more ergonomic design might infer the project from a session-level default. That is a session management question outside the scope of generation.

**Tool evolution governance.** As agents use the server in production, the observable signals (retry rates, tool sequences) should feed back into overlay changes. Who owns `mcp_workflows.yaml`, what the review process looks like, and how often overlays are updated — that needs a process design that is separate from the technical design.

**Streaming responses.** Workflow tools currently poll synchronously. For very long operations (cluster creation can take 15+ minutes), polling in the tool handler may time out before the client does. Streaming via MCP's SSE transport is the right long-term answer but requires MCP server infrastructure changes outside this proposal's scope.

---

## Open questions

**Should the generator live in the Atlas CLI repository or in the MCP server repository?**

The generator should stay in the Atlas CLI repository, next to the spec and overlays it reads. The MCP server imports the generated Go file. Keeping the generator with the spec prevents the overlay and spec from diverging across two repositories. The MCP server's only build-time dependency is the generated file, not the generator or the spec.

**Should raw tools pin to newest stable or allow version selection?**

The proposal recommends pinning at generation time to the newest non-deprecated stable version. Exposing version selection to agents adds complexity with little benefit — agents do not know Atlas API version semantics and will not make good version choices. If a user explicitly needs a specific version, they should use the CLI. If this assumption is wrong for a specific use case, it can be revisited.

**How are breaking changes in workflows handled?**

If a workflow's underlying `operationId` changes or is removed from the spec, the generator will fail the build — the same way a missing watcher `operationId` fails today. That is intentional. Breaking API changes require an explicit overlay update before the build passes. This ensures breaking changes are surfaced and handled, rather than silently producing a broken tool.

**What is the right timeout for poll steps?**

The proposal sets 1800 seconds (30 minutes) as the default for cluster operations. That covers normal cluster creation (typically 7–12 minutes) with headroom. The timeout is configurable per workflow step in the overlay definition. The MCP client's own timeout may be shorter than the workflow timeout, in which case the tool call will fail with a partial completion message. This is a known limitation until streaming is implemented.
