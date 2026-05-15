# MCP Tool Design: Research Synthesis

This document synthesizes recent writing (February–May 2026) on building MCP tools, with a focus on what the research says about auto-generation from API specs, tool design at scale, and workflow composition. The goal is to extract what matters for the Atlas CLI → MCP adaptation.

---

## The central finding: 1:1 API-to-tool mapping does not work

This appears in every serious article on the subject. It is not a minor caveat — it is the primary failure mode for MCP servers built by teams coming from API design backgrounds.

The concrete evidence:

* **GitHub Copilot** reduced its tool count from 40 to 13 and saw measurable benchmark improvements.
* **Block rebuilt its Linear MCP server three times**, going from 30+ tools to 2. Each rebuild was a response to observed agent failures, not a premature optimization.
* **Chroma research** tested 18 LLMs and found performance degradation as input length grows, even on simple tasks. The assumption that a model handles the 10,000th token as reliably as the 100th does not hold.
* At 50+ tools, tool schemas alone consume 5–7% of the model's context window before a single user message arrives. That number only gets worse as tool descriptions are made more informative.
* Anthropic's own guidance puts the accuracy degradation threshold at 30–50 tools. Arcade's research (54 Patterns) puts it even lower in practice.

The failure modes past ~20 tools are two distinct problems. The first is context bloat: schemas, names, and descriptions consume tokens that were needed for reasoning. The second is tool hallucination: models start inventing nonexistent tool names, conflating parameters between tools, or calling the right tool with arguments from a different tool's schema. These are different bugs and require different fixes, but both have the same root cause — too many tools in context.

---

## What Cloudflare learned

The [Cloudflare CF CLI Local Explorer post](https://blog.cloudflare.com/cf-cli-local-explorer/) is the closest analog to the Atlas CLI auto-generation problem that exists in public writing. Cloudflare builds a CLI, an SDK, a Terraform provider, and an MCP server from a single source of truth. Their experience is directly applicable.

**The schema-first architecture.** Cloudflare moved beyond OpenAPI as a source of truth and replaced it with a custom TypeScript schema system. Their description: "The schema format is 'just' a set of TypeScript types with conventions, linting, and guardrails to ensure consistency." From this schema, they generate all their surface areas — CLI, SDK, Terraform, MCP — using different templates against the same model. This is structurally identical to the Atlas CLI approach (spec + overlays → generator → template), but Cloudflare went further by owning the schema format itself rather than adapting OpenAPI.

**Consistency enforced at the schema layer.** Cloudflare establishes guardrails like: always use `get` (never `info`), standardize on `--force` (never `--skip-confirmations`), always support `--json`. These are linting rules on the schema, not conventions in documentation. The result is that agents encounter consistent command shapes across all Cloudflare products, which matters because agents build mental models from patterns — inconsistent naming forces them to treat every tool as novel.

**Local-remote API parity.** Local Explorer mirrors the Cloudflare API for local resources (KV, R2, D1, Durable Objects). Developers and agents interact with identical command structures whether targeting local or remote. This eliminates a class of agent errors where the model understands the operation but gets confused about which environment it is targeting.

**Agent-centric context engineering.** The design uses clear signals about whether commands operate on local or remote resources, predictable defaults, and consistent output formatting. These sound like small decisions but they directly affect whether an agent can reliably interpret results across different tools.

The lesson from Cloudflare is that schema quality and consistency matter more than coverage. A smaller, internally consistent schema with good descriptions produces better agent behavior than a complete schema with variable quality.

---

## The OpenAPI auto-generation landscape

Several tools now generate MCP servers from OpenAPI specs: Stainless, FastMCP, AWS's openapi-mcp-server, Speakeasy's Gram, and various open-source generators. Their existence is useful as a signal — the problem is considered solved at the syntax level. The question is whether that matters.

**Neon's conclusion:** treat auto-generation as a starting point, not a final product. Their process:

1. Use the TypeScript SDK as a foundation (not raw OpenAPI).
2. Implement only tools that are genuinely appropriate for LLM use.
3. Build higher-level workflows that combine multiple endpoints — their example is a "streamlined database migration testing" tool that never appears in the raw API.
4. Aggressively prune everything that doesn't pass the "would an agent actually need this" test.

Neon's word is "aggressively." The implication is that the natural output of a generator requires removing the vast majority of what it produces.

**Speakeasy's conclusion:** documentation quality determines tool quality. The generator handles syntax transformation reliably. Semantic quality — helping agents understand *when* and *how* to use a tool — requires human judgment. Specific gaps that cause hallucinations:

* Unclear endpoint purpose (what is this for, not just what it does)
* Vague parameter descriptions
* Missing response examples
* No guidance on parameter lookup patterns (e.g., "get task IDs via `listTasks` first")

Speakeasy ships an `x-speakeasy-mcp` extension for overriding tool names in agent contexts — a direct parallel to Atlas CLI's `x-xgen-atlascli` extension for CLI-specific customizations. This is not a coincidence; it is the right architectural answer. You cannot fix semantic quality by forking the spec.

**The shared takeaway** across Neon, Speakeasy, and Cloudflare: generation handles the mechanical work. Every generator in this space is essentially doing what Atlas CLI's `api-generator` does — parsing a schema and emitting code via a template. The differentiation is in the curation layer on top: overlays, extensions, and higher-level workflow definitions that inject human judgment into the process without breaking the generation pipeline.

---

## Tool design patterns from the research

Arcade's 54-pattern taxonomy organizes the space across three axes — maturity (atomic to orchestrated), integration type, and access pattern — and ten concern categories. The most practically useful patterns for the Atlas use case:

**Outcomes over operations.** Phil Schmid's formulation is the clearest: instead of `getUser`, `listOrders`, and `getStatus`, expose `trackLatestOrder(email)`. The single tool handles all three API calls internally. This reduces context window bloat and API round-trips simultaneously, and the description can speak to user intent rather than API mechanics.

**Flatten arguments.** Nested objects in tool input schemas cause hallucinations because the model has to reason about structure at the same time as it reasons about values. Top-level primitives with constrained types (enums as `Literal`, explicit string formats) reduce the cognitive load on the model. This is a direct implication for any tool generated from an OpenAPI spec, since REST request bodies are often deeply nested.

**Service-prefixed names.** The `{service}_{action}_{resource}` convention (e.g., `atlas_create_cluster`, `atlas_list_projects`) prevents naming collisions when multiple MCP servers are active simultaneously. It also helps the model identify which server owns a tool without needing to read the description.

**Descriptions written for agents, not developers.** Speakeasy's observation that "humans can infer context from brief descriptions; AI agents cannot" is backed up by real failure data. A description that says "Returns cluster details" fails agents. A description that says "Use this to check if a cluster is ready after creation, or to get connection strings for an existing cluster. Requires a project ID and cluster name." succeeds. The difference is intent guidance, not information density.

**Observable signals drive design evolution.** Arcade's framework provides a practical rule: if you see high retry rates, the tool's description needs work. If you see agents calling the same sequence of tools repeatedly across sessions, that sequence should become a single workflow tool. Let observed failure modes drive the curation decisions, rather than trying to anticipate them entirely upfront.

**Structured error responses that enable self-correction.** Errors should include actionable recovery guidance: "rate limited, retry after 30 seconds" rather than just the HTTP status code. Agents that receive structured errors can self-correct; agents that receive opaque errors stall or hallucinate a recovery path.

**The Tools/Resources/Prompts split.** The DDD article makes a useful point about using all three MCP capability types. Tools for active operations with side effects. Resources for read-only contextual data (accessible via URIs, not tool calls). Prompts for templated workflows that guide the agent through a multi-step process. Keeping GET operations in Resources rather than Tools removes them from the tool list without removing their accessibility.

---

## Progressive tool discovery

Amazon Prime Video faced the problem at scale: a centralized MCP server with hundreds of tools, but individual tasks only needing 3–4 of them. Their solution exploits two MCP protocol features: tool list change notifications and session tracking.

The flow:

1. The agent starts with a single "find tools" meta-tool visible in context.
2. When the agent calls it with a problem category (operations, results, training, etc.), the server maps that session to the relevant tools.
3. A notification fires, the agent re-fetches the tool list, and now sees only the 3–4 tools appropriate for its current task.
4. The agent can switch categories mid-session, dropping old tools and loading new ones.

The result: hundreds of tools become manageable without removing coverage, at the cost of one additional round-trip at the start of each task.

The design trade-offs are real. Category design is hard — too many categories reproduces the original problem one level up, and the categories have to be stable enough that the agent can reliably navigate them. The approach also introduces protocol dependency on MCP session tracking, which the 2026 MCP roadmap is still evolving (stateful sessions are a known scaling problem).

Progressive discovery is best understood as a scaling mechanism, not a replacement for curation. Amazon's approach works because they also maintain "clean tool definitions" — the discovery mechanism routes the agent to good tools, not just fewer tools.

---

## The 2026 MCP roadmap and what it means for generation

The official roadmap identifies four priorities: transport scalability (Streamable HTTP at scale), agent communication (Tasks primitive refinement), governance maturation, and enterprise readiness (audit trails, SSO, gateway behavior, configuration portability).

Two items are relevant to tool generation:

**Stateful sessions are a first-class problem.** The protocol's current session model makes horizontal scaling hard. Any tool generator that produces tools relying on server-side session state is building on unstable ground. The architectural preference should be stateless tool execution — parameters carry everything needed, server-side state is minimal. This aligns with the REST-like design of the Atlas CLI's executor (the `CommandRequest` struct is self-contained).

**Enterprise features as extensions.** The roadmap explicitly endorses handling enterprise concerns through extensions rather than core spec changes. This is the same philosophy as the overlay system: extend the spec without forking it. For tool generation, it reinforces the pattern of using extension namespaces (`x-mcp`, `x-xgen-atlascli`) for tooling-specific concerns.

---

## Synthesis for the Atlas CLI → MCP adaptation

Drawing the thread across all of this:

**The auto-generation pipeline is the right foundation.** Every source confirms that generators handle the mechanical transformation correctly. Stainless, FastMCP, Speakeasy's Gram — they all do what `api-generator` does. The Atlas CLI is not behind the curve; the pipeline is sound.

**The gap is the curation layer.** Neon, Cloudflare, and Speakeasy all arrived at the same place: the generator is the starting point. What makes tools agent-usable is the work that happens in overlays — description overrides, tool groupings, workflow definitions, skip directives. This is already the mechanism; it just needs to be expanded with agent-centric concerns.

**Workflow tools are the primary product.** Every source that discusses production agent failures points to multi-step operations as the place where 1:1 tools fail most visibly. The `x-mcp-workflows` overlay extension described in the alternatives document is not a nice-to-have — it is the mechanism that separates tools that work from tools that work for agents.

**Keep the raw operation layer, but filter it aggressively.** The Atlas API has ~700 operations. The right exposure is probably 20–40 workflow tools plus 40–60 curated raw tools, not 700. GitHub Copilot's trajectory (40→13) and Block's (30+→2) are reference points, but those numbers are appropriate for narrow-scope servers. A general-purpose Atlas MCP server has broader scope and can sustain more tools than a single-purpose Linear integration.

**Progressive discovery is the scaling answer.** Once the server has curated workflow tools and filtered raw tools, progressive discovery (Amazon's pattern) addresses the remaining context problem. The implementation is small — one meta-tool plus category mapping — but it unlocks the ability to register significantly more tools without context bloat. Design the categories around Atlas domains: clusters, projects, networking, users, backups, search.

**The 5–15 tool guidance is for narrow servers.** Phil Schmid's recommendation is widely cited but applies to servers with a single well-defined purpose. A general-purpose Atlas MCP server is more like Amazon Prime Video's case — many tools organized into domains — and the progressive discovery pattern is the appropriate solution for that scale.

---

## Sources

* [Cloudflare CF CLI Local Explorer](https://blog.cloudflare.com/cf-cli-local-explorer/)
* [54 Patterns for Building Better MCP Tools — Arcade](https://www.arcade.dev/blog/mcp-tool-patterns)
* [MCP is Not the Problem, It's Your Server — Phil Schmid](https://www.philschmid.de/mcp-best-practices)
* [Auto-generating MCP Servers from OpenAPI Schemas: Yay or Nay? — Neon](https://neon.com/blog/autogenerating-mcp-servers-openai-schemas)
* [Generating MCP Tools from OpenAPI: Benefits, Limits, Best Practices — Speakeasy](https://www.speakeasy.com/mcp/tool-design/generate-mcp-tools-from-openapi)
* [Progressive Tool Discovery for MCP Servers to Manage Context at Scale — Amazon / ZenML](https://www.zenml.io/llmops-database/progressive-tool-discovery-for-mcp-servers-to-manage-context-at-scale)
* [Building Scalable MCP Servers with Domain-Driven Design — Chris Hughes](https://medium.com/@chris.p.hughes10/building-scalable-mcp-servers-with-domain-driven-design-fb9454d4c726)
* [A Practical Guide to Architectures of Agentic Applications — Speakeasy](https://www.speakeasy.com/mcp/using-mcp/ai-agents/architecture-patterns)
* [The 2026 MCP Roadmap — Model Context Protocol Blog](https://blog.modelcontextprotocol.io/posts/2026-mcp-roadmap/)
* [MCP Tool Design: Why Your AI Agent Is Failing — AWS Heroes / DEV Community](https://dev.to/aws-heroes/mcp-tool-design-why-your-ai-agent-is-failing-and-how-to-fix-it-40fc)
* [When Too Many Tools Become Too Much Context: RAG-MCP — Pankaj Pandey](https://medium.com/@pankaj_pandey/when-too-many-tools-become-too-much-context-a-deep-dive-into-rag-mcp-9b628c8476d3)
