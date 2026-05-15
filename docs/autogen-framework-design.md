# Auto-generation Framework: Technical Design

This document describes the auto-generation framework used in Atlas CLI to turn an OpenAPI spec into executable commands, and what it takes to apply the same approach to a different consumer — specifically an MCP server.

---

## The problem it solves

The Atlas Admin API has hundreds of endpoints. Maintaining hand-written command definitions, flag declarations, and help text for each one does not scale. More importantly, it creates drift: the API changes, commands lag behind, documentation goes stale.

The framework solves this by making the OpenAPI spec the only source of truth. Every command, every parameter, every description, and every supported version is derived from the spec. Adding or changing an API endpoint requires no human intervention on the CLI side — running `make gen-api-commands` picks it up.

---

## High-level pipeline

```
spec.yaml
    │
    ▼
apply-overlay         (OpenAPI Overlay spec; pure YAML patching)
    │
    ▼
spec-with-overlays.yaml
    │
    ▼
api-generator         (Go binary; reads spec, emits Go source via text/template)
    │
    ▼
internal/api/commands.go   (static Go variable; the full command tree)
    │
    ▼
CLI builder           (iterates commands.go at startup, registers Cobra commands)
```

Each stage has a single, narrow responsibility. The generator does not know anything about Cobra. The CLI builder does not know anything about OpenAPI. The shared data model — the `Command` struct — is the contract between them.

---

## Stage 1: The upstream spec

`tools/internal/specs/spec.yaml` is imported from the Atlas API team and committed as-is. Nobody edits it by hand. When the API team ships a new spec, you replace this file and re-run generation.

This is the key invariant: the spec file is never modified in this repository. All Atlas CLI-specific concerns live in overlays.

---

## Stage 2: Overlays

The overlay tool (`tools/cmd/apply-overlay`) reads the spec and applies a set of YAML overlay files in sorted alphabetical order, producing `spec-with-overlays.yaml`.

Overlays follow the [OpenAPI Overlay 1.0 spec](https://spec.openapis.org/overlay/v1.0.0.html). Each overlay file has a list of `actions`, where every action targets a specific path in the spec using a JSONPath selector and either `update`s or `remove`s content at that path.

```yaml
overlay: 1.0.0
info:
  title: Configure groupId flag
  version: 1.0.0
actions:
  - target: $.components.parameters.groupId
    update:
      x-xgen-atlascli:
        aliases:
          - projectId
```

All CLI-specific additions live under the `x-xgen-atlascli` extension key. This namespace is the contract between overlay authors and the generator. The generator reads nothing outside this namespace (except standard OpenAPI fields like `operationId`, `parameters`, `responses`, and `description`).

Overlays handle:

* **Flag aliases** — `groupId` is also accepted as `projectId` because that flag name predates the auto-gen framework.
* **Description overrides** — when the upstream description is wrong or too short for CLI use, the overlay replaces it.
* **Skip directives** — `x-xgen-atlascli: skip: true` excludes an operation from generation entirely.
* **Operation ID overrides** — when the upstream `operationId` is too long or inconsistent with CLI conventions, the overlay renames it without touching the spec.
* **Command aliases** — entire command names can have aliases at the operation level.
* **Watchers** — declarative polling behavior for long-running operations (described in detail below).
* **Usage examples** — structured examples that appear in `--help` output.

Each overlay file addresses one concern. `groupID_flag.yaml` is about the `groupId` alias. `watchers.yaml` is about all watcher definitions. Keeping them separate means the diff for any change is small and reviewable in isolation.

---

## Stage 3: Code generation

`tools/cmd/api-generator` is a Go binary that:

1. Loads and validates the overlaid spec using `kin-openapi/openapi3`.
2. Iterates every path/verb combination and calls `operationToCommand()` for each one.
3. Groups commands by their OpenAPI tag.
4. Sorts groups alphabetically and commands within each group by `operationId`.
5. Validates all watcher definitions (see below).
6. Executes a `text/template` against the sorted command tree.
7. Runs the output through `go/format` before writing it to disk.

The generator can produce two output types via `--output-type`:

* `commands` — the full command tree as a Go variable (`internal/api/commands.go`)
* `metadata` — documentation metadata for the docs generator (`tools/cmd/docs/metadata.go`)

Both use the same parse/validate pipeline; only the template differs.

### The Command data model

The generator emits values of these types, defined in `tools/shared/api`:

```go
type Command struct {
    OperationID       string
    ShortOperationID  string
    Aliases           []string
    Description       string
    RequestParameters RequestParameters
    Versions          []CommandVersion
    Watcher           *WatcherProperties
}

type RequestParameters struct {
    URL             string       // e.g. /api/atlas/v2/groups/{groupId}/clusters
    QueryParameters []Parameter
    URLParameters   []Parameter
    Verb            string       // e.g. http.MethodGet
}

type CommandVersion struct {
    Version              Version    // StableVersion | UpcomingVersion | PreviewVersion
    Sunset               *time.Time
    PublicPreview        bool
    Deprecated           bool
    RequestContentType   string     // e.g. "json"
    ResponseContentTypes []string   // e.g. ["json", "csv"]
}

type Parameter struct {
    Name        string
    Aliases     []string
    Short       string
    Description string
    Required    bool
    Type        ParameterType  // {IsArray bool, Type "string"|"int"|"bool"}
}
```

These types live in `tools/shared/api`, which has no CLI dependencies. That is deliberate — anything that needs to understand commands imports this package, not Cobra or any CLI-specific package.

### Version model

Atlas API versions take one of three forms:

* `2023-02-01` — a stable, date-stamped version
* `2023-02-01.upcoming` — not yet generally available, same date format
* `preview` — undated; subject to breaking changes at any time

The `Version` interface has `StabilityLevel()`, `Less()`, `Equal()`, and `String()`. The generator sorts versions oldest-first and drops any version whose `Sunset` date is already in the past at generation time.

Atlas uses versioned content types to express both version and format in a single header value:

```
application/vnd.atlas.2023-02-01+json
```

The generator parses these from the `responses` and `requestBody` content keys to build `CommandVersion.ResponseContentTypes` and `CommandVersion.RequestContentType`.

### What gets skipped

An operation is skipped (produces no command) if:

* It has `x-xgen-atlascli: skip: true` in the overlay.
* It has zero non-sunset versions after the sunset filter runs (the endpoint is fully retired).
* It has more or fewer than exactly one tag (the tag becomes the command group name; ambiguity is a generator error).

---

## Stage 4: Runtime wiring

`internal/cli/api/api.go` is the only file that knows about Cobra. `Builder()` does one thing: iterate `api.Commands` and register every command.

For each `Group`, it creates a `cobra.Command` with `Use` set to the camelCase tag name.

For each `Command` within a group, `convertAPIToCobraCommand()`:

1. Sets `Use` to the camelCase `operationId` (or `shortOperationId` if set, with the original as an alias).
2. Registers a `--version` flag listing all supported versions; the default is the newest stable version.
3. Registers `--output` (content type / format) and `--output-file`.
4. Registers `--file` for mutations that have a request body.
5. Registers `--watch` and `--watch-timeout` if the command has a watcher.
6. Registers every URL parameter and query parameter as a typed `--flag`.
7. Installs a `NormalizeFunc` that rewrites flag aliases to canonical names before Cobra processes them.
8. In `PreRunE`, fills untouched flags from the user's profile (so `--projectId` does not need to be typed every time).
9. In `RunE`, builds a `CommandRequest`, executes it via the `Executor`, formats the output, and optionally runs the watcher.

The executor (`internal/api/executor.go`) is a thin wrapper around an HTTP client. It calls `ConvertToHTTPRequest()`, which builds the full URL, query string, and `Accept`/`Content-Type` headers, then executes the request.

---

## Watchers

Some API operations are asynchronous. You call `POST /clusters` and the cluster spends several minutes provisioning. Watchers let the user pass `--watch` and block until the cluster reaches a stable state.

A watcher is defined in the overlay as a property on the triggering operation:

```yaml
x-xgen-atlascli:
  watcher:
    get:
      operation-id: getGroupCluster     # which operation to poll
      version: "2024-08-05"
      params:
        clusterName: body:$.name        # extract from mutation response
        groupId: body:$.groupId
    expect:
      match:
        path: $.stateName               # JSONPath into the poll response
        values:
          - IDLE
```

The `params` mapping uses a three-function DSL:

* `input:<flag>` — copy the value from the original mutation's CLI flag
* `body:<jsonpath>` — extract from the mutation's JSON response body
* `const:<value>` — use a literal string

The generator validates that:

* The `operation-id` exists in the command tree.
* The specified `version` exists on that operation.
* Every parameter named in `params` exists on the watcher operation.
* No required parameters on the watcher operation are missing from `params`.

Misconfigured watchers are caught at generation time, not at runtime.

At runtime, `Watcher.Wait()` polls once per second using `WatchOne()`, which executes the GET command and checks the HTTP code and/or a JSONPath condition against the response body.

---

## Applying this to an MCP server

The data model and the HTTP execution layer are fully reusable. The only thing that changes is the consumer layer — instead of Cobra commands, you register MCP tools.

### What stays the same

* The upstream spec and overlay files. You can add a new extension namespace (`x-mcp`) alongside `x-xgen-atlascli` in the same overlay files if you need MCP-specific customizations, or keep them in separate overlay files and apply both sets.
* The apply-overlay binary. It's a generic OpenAPI Overlay implementation with no Atlas CLI knowledge.
* The `tools/shared/api` data model. `Command`, `Parameter`, `CommandVersion`, `Version`, and `WatcherProperties` are the portable intermediate representation.
* The version parsing and sorting logic.
* The watcher validation logic.

### What changes

**The generator template.** Replace `commands.go.tmpl` with a template that emits your target language's tool registration syntax. For a TypeScript MCP server, this might emit an array of `{ name, description, inputSchema, handler }` objects. For Python, it might emit a list of `@mcp.tool` decorated functions. The template receives the same sorted `GroupedAndSortedCommands` data — only the output format changes.

**The executor consumer.** Instead of `internal/cli/api`, you write an MCP tool handler. That handler receives the tool call's arguments (which map to `Parameters`), constructs a `CommandRequest`, calls `executor.ExecuteCommand()`, and returns the result as an MCP tool response.

**Flag-to-schema translation.** Cobra flags become JSON Schema properties in the MCP `inputSchema`. The `Parameter` struct has everything you need: `Name`, `Description`, `Required`, `Type.Type` (`string`/`int`/`bool`), and `Type.IsArray`. URL parameters and query parameters are both flat properties on the schema — there is no distinction from the tool caller's perspective.

**Version selection.** In the CLI, the user passes `--version`. In MCP, you either expose this as an optional input property or pin the newest stable version at generation time. Pinning at generation time is simpler and better for most MCP use cases, since tool callers are not expected to know Atlas API versioning.

**Watchers.** In the CLI, watchers block synchronously with a polling loop. In MCP, you can model this as a streaming response, as a separate `waitFor<Operation>` tool, or by polling in the tool handler before returning. The `WatcherProperties` struct is the spec for all three approaches.

**Skipping operations.** Not every Atlas API endpoint makes sense as an MCP tool. Use the existing `x-xgen-atlascli: skip: true` mechanism, or add `x-mcp: skip: true` in a new overlay. The generator only needs a small change to check the new flag.

### Recommended architecture

```
spec.yaml
    │
    ▼
apply-overlay (existing tool, unchanged)
    │
    ▼
spec-with-overlays.yaml
    │
    ├──► api-generator --output-type commands  →  internal/api/commands.go  (CLI)
    │
    └──► mcp-generator --output-type tools     →  src/generated/tools.ts    (MCP)
         (new binary; same parsing logic, different template)
```

Both generators can share the spec parsing code. If the codebase is Go, extract the `specToCommands()` conversion into a shared package and import it in both binaries. If the MCP server is in a different language, the generator is still written in Go and emits code in the target language — the same way `api-generator` emits Go from a Go program.

The clean separation between parse/convert (generator) and consume (CLI or MCP) is what makes this adaptation tractable. You are not refactoring the existing CLI; you are adding a parallel consumer that reads the same intermediate data.

---

## Invariants worth preserving

A few properties of the existing framework are worth carrying over deliberately, because they are easy to break when adapting to a new context.

**Validation at generation time.** Watcher correctness, parameter deduplication, and required-field coverage are all caught during `make gen-api-commands`. This keeps the runtime clean. If you add MCP-specific validation (e.g., tool name uniqueness, schema type compatibility), put it in the generator, not in the server startup code.

**The spec is never forked.** Every customization goes through an overlay. The moment you start editing `spec.yaml` directly, you lose the ability to upgrade from the upstream API team without a painful merge.

**The generated file is committed.** `commands.go` is checked into version control even though it is generated. This means CI sees exactly what the server will run, diffs are reviewable in PRs, and there is no build-time dependency on the spec at server startup. Apply the same practice to whatever file the MCP generator produces.

**Sort order is deterministic.** Groups and commands are sorted alphabetically. Templates iterate with `range .` on sorted slices, not on maps. This ensures the generated file is stable across runs even when the spec's internal map ordering changes.
