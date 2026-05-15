# MCP Tool Generation: Design Alternatives

This document explores the design space for auto-generating MCP tools from an OpenAPI spec, using the Atlas CLI framework as a starting point. The core tension is that the Atlas CLI framework does a 1:1 mapping (one operation → one command), which works for a CLI but breaks down in MCP. This document works through the alternatives and their trade-offs.

---

## Why 1:1 mapping does not work for MCP

The Atlas Admin API has roughly 700 operations. A CLI can expose all of them as leaf commands because humans navigate hierarchically (`atlas api clusters createOne`) and tab-complete their way through. A CLI user can also read a man page.

An LLM cannot. Tool selection in MCP happens by the model reading all tool names and descriptions simultaneously and picking the best match. Two things degrade that process fast:

**Too many tools.** Most model providers start seeing quality degradation somewhere between 20 and 60 tools in context. At 700, tool selection is effectively random. Even at 100, a model looking for "create a cluster" will be distracted by `listClusterAuditingExpressions` and `getClusterAdvancedSettings`.

**API-shaped descriptions.** Descriptions written for API reference docs describe what the endpoint does mechanically. They do not describe when to use it, what preconditions it requires, or what the user's goal is. An LLM agent needs the latter. `Creates, edits, verifies, and removes IP access list entries for the specified project` is the kind of text that confuses tool selection. A useful MCP tool description answers the question: "what is the user trying to accomplish when they invoke this?"

**Multi-step operations.** Many useful tasks in Atlas require a sequence of API calls. Creating a cluster means calling `createCluster`, then polling `getCluster` until `stateName == IDLE`, then possibly calling `createDatabaseUser` and `createProjectIpAccessList`. A 1:1 mapping forces the agent to discover and sequence these calls itself, which is error-prone and produces bad user experiences. The CLI partially solves this with watchers, but even watcher-aware CLI commands are still single-operation.

The conclusion is: the right design for MCP tool generation is different from the right design for CLI command generation, even when both draw from the same OpenAPI spec.

---

## The design space

Five meaningfully different approaches exist. They are not mutually exclusive — the right answer for a real MCP server is likely a combination of two or three.

---

### Alternative 1: Filtered 1:1 mapping

The simplest adaptation: generate one MCP tool per operation, but use `skip` overlays aggressively to reduce coverage to the 50–80 operations that actually matter. Everything that's read-only and rarely needed by an agent, all admin/internal endpoints, all endpoints that require context an agent can't have — those get skipped.

**How it works.** This reuses the Atlas CLI pipeline almost unchanged. The generator emits MCP tool registrations instead of Cobra commands. The `--output-type tools` flag produces TypeScript/Python/Go tool definitions rather than `commands.go`. The filter logic lives in overlays: `x-mcp: skip: true` on everything you do not want exposed.

**What the overlay looks like:**
```yaml
overlay: 1.0.0
info:
  title: MCP tool filter - skip audit and metrics endpoints
  version: 1.0.0
actions:
  - target: $.paths['/api/atlas/v2/groups/{groupId}/auditLog'].get
    update:
      x-mcp:
        skip: true
  - target: $.paths['/api/atlas/v2/groups/{groupId}/metrics/**']
    update:
      x-mcp:
        skip: true
```

**Trade-offs.**
- Generating this is easy — the pipeline already exists.
- Tool descriptions are still API-shaped, not agent-shaped. You can mitigate this with description overrides in overlays, but you are fighting the grain of API documentation.
- No workflow composition. The agent still has to sequence multi-step operations.
- Good as a foundation layer but insufficient as a complete MCP tool design.

**When to use it.** As the baseline. You want this layer to exist because it gives the agent escape hatches for operations that no one thought to make into a workflow tool. The agent should be able to fall back to `raw_api_call` when no higher-level tool fits.

---

### Alternative 2: Resource/tool split following MCP semantics

MCP distinguishes between Resources (data that can be read) and Tools (actions that have side effects). GET operations are natural Resources. POST/PUT/PATCH/DELETE are natural Tools. This alternative exploits that distinction to remove all read operations from the tool list.

**How it works.** The generator applies a rule: operations with `GET` verb become MCP Resources; all other verbs become MCP Tools. Resources are accessed by the model via `read_resource` with a URI, not via tool call. This immediately halves the tool count.

Resources would use URI templates based on the operation's URL path:
```
atlas://groups/{groupId}/clusters
atlas://groups/{groupId}/clusters/{clusterName}
```

**What this changes in the pipeline.** The generator needs to emit two output types: a resource registry and a tool registry. The resource registry maps URI templates to operations. The tool registry maps tool names to operations. The overlay can add `x-mcp: expose-as: resource` or `x-mcp: expose-as: tool` to override the default verb-based rule.

**Trade-offs.**
- Reduces tool count significantly.
- Aligns with MCP's semantic model. Resources represent state; tools represent actions.
- Resources are less ergonomic for agents than tools in current MCP implementations. Most LLMs are better at tool-calling than resource-reading, and resources do not appear in the tool list that models use for planning.
- GET operations that are "search" operations (with many query parameters) are awkward as resources. A resource URI for `listClusters?pageNum=2&includeCount=true` is not a natural URI.
- Pagination is unsolved — MCP Resources do not have a standard pagination model.

**When to use it.** When the MCP server is consumed by a model that handles resources well, or when the primary goal is reducing tool count. Not recommended as the sole approach because of the ergonomics gap.

---

### Alternative 3: Tag-grouped dispatch tools

Instead of one tool per operation, expose one tool per OpenAPI tag — one tool for all of Clusters, one for Projects, one for Database Users. The tool accepts an `operation` parameter that specifies which API operation to call, and a `parameters` object for that operation's inputs.

**What a tool call looks like:**
```json
{
  "tool": "atlas_clusters",
  "input": {
    "operation": "createCluster",
    "version": "2024-08-05",
    "parameters": {
      "groupId": "abc123",
      "body": { "name": "my-cluster", "clusterType": "REPLICASET" }
    }
  }
}
```

**How it works.** The generator emits one tool per tag. Each tool's description covers the whole tag and lists available operations. The input schema has a discriminated union: `operation` is an enum of all operation IDs in the tag, and `parameters` is typed per-operation (or typed as a free-form object if the implementation accepts the loss of static typing).

**Trade-offs.**
- Reduces tool count to the number of tags (~30 for Atlas), which is manageable.
- The model must know the `operationId` to use the tool. This reintroduces the problem of API-shaped mental models. An agent that doesn't know `getGroupCluster` is a thing cannot easily discover it by looking at the `atlas_clusters` tool's description.
- Schemas become very large (a union of all operation schemas per tag). Models do not handle large schemas well.
- Static type checking is largely lost — you're back to a generic dispatch pattern.
- Discovery inside a tag requires either a long description or a separate `list_operations` call.

**When to use it.** Not recommended as the primary design. This solves the count problem but makes the model's job harder, trading one failure mode for another. Could work as a "power user" escape hatch alongside higher-level tools.

---

### Alternative 4: Workflow tools with declarative composition

This is the most significant departure from the CLI model. Instead of exposing API operations directly, you define workflows — multi-step sequences with a single user intent — and generate tools from those workflow definitions.

A workflow is not a single API operation. It is a named sequence of operations with data flow between steps, a coherent description written for agent consumption, and a simplified input schema that hides intermediate parameters.

**What a workflow definition looks like in an overlay:**

```yaml
overlay: 1.0.0
info:
  title: MCP workflow definitions
  version: 1.0.0
actions:
  - target: $.x-mcp-workflows
    update:
      createCluster:
        description: |
          Creates a new Atlas cluster and waits until it is ready to accept connections.
          Use this when the user wants to provision a new database cluster in a project.
          Requires: project ID, cluster name, and cloud provider/region.
        input:
          groupId:
            from: parameter
            operation: createCluster
            name: groupId
          clusterName:
            type: string
            description: Name for the new cluster
          provider:
            type: string
            enum: [AWS, GCP, AZURE]
            description: Cloud provider to deploy on
          region:
            type: string
            description: Provider-specific region code (e.g. US_EAST_1)
        steps:
          - id: create
            operationId: createCluster
            version: "2024-08-05"
            body:
              name: "{{ input.clusterName }}"
              clusterType: REPLICASET
              replicationSpecs:
                - zoneName: Zone 1
                  regionConfigs:
                    - providerName: "{{ input.provider }}"
                      regionName: "{{ input.region }}"
                      electableSpecs:
                        instanceSize: M10
                        nodeCount: 3
          - id: wait
            type: poll
            operationId: getGroupCluster
            version: "2024-08-05"
            params:
              groupId: "{{ input.groupId }}"
              clusterName: "{{ steps.create.response.name }}"
            until:
              path: $.stateName
              values: [IDLE]
        output:
          clusterName: "{{ steps.create.response.name }}"
          connectionString: "{{ steps.wait.response.connectionStrings.standardSrv }}"
```

**How the generator handles this.** The generator reads the `x-mcp-workflows` extension block, resolves all referenced `operationId`s to their `Command` definitions (for type validation and header construction), and emits workflow tool handlers in the target language. Each step in the sequence uses the same `executor.ExecuteCommand()` machinery as raw operation tools — the workflow is just orchestration on top.

**What this produces.** The tool `createCluster` takes four parameters (`groupId`, `clusterName`, `provider`, `region`), handles the full create+poll sequence internally, and returns connection strings. The agent never sees the intermediate `getGroupCluster` polling calls.

**Trade-offs.**
- This is the most useful design for agents. The tool descriptions map to user intents. The schemas are small. The agent does not need to know Atlas API internals.
- Workflow definitions require human authorship. You cannot auto-generate workflows from an OpenAPI spec alone — you need someone who understands what users are trying to accomplish. This is a one-time cost per workflow, but it is real cost.
- The template DSL for data flow (`{{ steps.create.response.name }}`) adds complexity to the generator.
- Error handling mid-workflow is more complex than single-operation errors. If `createCluster` succeeds but the poll times out, the cluster exists but is not returned — the agent is in a confused state. Rollback behavior needs a design decision.
- Keeping workflow definitions in sync with spec changes requires attention. If `createCluster`'s request body schema changes, the workflow body template may need updating.

**When to use it.** As the primary tool surface. Every tool the agent reaches for first should be a workflow tool. Raw operation tools are the fallback.

---

### Alternative 5: Layered tool sets with dynamic exposure

The model context limit problem is not just about count — it is about relevance. If an agent is helping a user manage clusters, it does not need database user tools in context. This alternative registers a large tool set but only injects relevant tools based on conversation context.

**How it works.** The MCP server implements tool filtering at the session or message level. Tools are grouped into feature sets (clusters, networking, users, backups, etc.). The server starts with a minimal bootstrap set — maybe 5–10 tools — and exposes a `get_available_tools` meta-tool that returns tools relevant to a given intent or resource type.

```json
{
  "tool": "get_available_tools",
  "input": {
    "intent": "I want to set up a new Atlas project with a cluster"
  }
}
```

The meta-tool returns a filtered list of tool names and descriptions. The client then injects those tools into context for subsequent turns.

This approach is implemented in some MCP servers today as "tool pagination" and is likely to become a more standard pattern as the MCP specification matures.

**Trade-offs.**
- Solves the context limit problem most cleanly.
- The meta-tool layer adds a round-trip. Agents that do not know to call `get_available_tools` first will work with a degraded tool set.
- Requires session state or conversation context to determine which tools are relevant.
- Does not solve the API-shaped description problem — tools still need agent-friendly descriptions.
- The MCP spec does not yet standardize this pattern, so implementation is bespoke.

**When to use it.** As a scaling mechanism on top of Alternative 1 or Alternative 4. Start with a curated tool set; add dynamic exposure when the tool count becomes a problem.

---

## Recommendation

A production MCP server for Atlas should combine three of these alternatives:

**Layer 1 — Workflow tools (Alternative 4) as the primary surface.** Define 20–40 workflow tools that map to the most common user intents: create cluster, delete cluster, add database user, configure IP access, create backup policy, and so on. These are the tools the agent reaches for 90% of the time. Write their descriptions for agents, not API readers.

**Layer 2 — Filtered raw operation tools (Alternative 1) as the fallback.** Expose the 50–80 most useful individual operations directly. An agent that needs to do something no workflow covers should be able to call the raw operation. These tools have API-shaped descriptions and full parameter sets, but they exist. Skip everything that is read-heavy, audit-related, internal, or requires complex setup.

**Layer 3 — Resource/tool split (Alternative 2) for reads.** GET operations that return single resources by ID or list resources become MCP Resources rather than tools. This keeps them accessible without polluting the tool list.

Together this gives the agent roughly 25–45 tools in the default configuration — inside the performance window — plus a resource layer for reads and an escape hatch for edge cases.

---

## How the generator pipeline adapts

The overlay pipeline already supports everything this architecture needs. The changes are in what the generator reads and what it emits.

**New extension keys in overlays:**

```
x-mcp:
  skip: true                     # exclude from MCP entirely
  expose-as: resource            # emit as MCP Resource, not Tool
  tool-description: "..."        # override tool description for agent use
  workflow: <workflow-id>        # this operation belongs to workflow X
```

**New top-level overlay section:**
```
x-mcp-workflows:
  <workflowId>:
    description: "..."
    input: { ... }
    steps: [ ... ]
    output: { ... }
```

**Generator changes:**

The generator gains a new `--output-type mcp-tools` mode. In this mode:

1. It reads `x-mcp-workflows` and resolves each step's `operationId` to a `Command` struct — same lookup it already does for watcher validation.
2. It emits workflow tool handlers that internally sequence `executor.ExecuteCommand()` calls.
3. It reads `x-mcp: expose-as: resource` and emits resource definitions for those operations.
4. It reads the filtered raw operations and emits single-operation tools for them.
5. It validates at generation time that all workflow `operationId` references exist and that all referenced parameters exist on those operations.

The core parsing (`specToCommands()`), the `Command` data model, the `executor`, and the HTTP request builder are unchanged. The generator gains a new output branch; everything upstream of it stays the same.

---

## The key shift from CLI to MCP

The CLI framework optimized for completeness: expose every operation, let the user find what they need via tab-completion and help text. An agent cannot browse. It must be steered.

MCP tool generation optimizes for intent coverage: expose tools that match what a user is trying to accomplish, not tools that match what the API offers. This is a higher-level abstraction, and it requires human judgment about what users want — judgment that cannot come from parsing an OpenAPI spec alone.

The overlay mechanism is the right place for that judgment to live. A workflow definition in an overlay is explicit, reviewable, versioned, and kept close to the spec that backs it. When the underlying API changes, the overlay is the first thing to update, and the generator catches any mismatch between the overlay's referenced operations and the spec. The same invariant the CLI framework enforces — spec is the source of truth, customizations live in overlays, mismatches are caught at generation time — holds for MCP workflow tools.
