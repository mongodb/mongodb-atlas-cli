//! Intermediate representation of the generated CLI.
//!
//! Produced by [`from_spec::from_config`], consumed by
//! [`emit::to_tokens`]. Kept free of proc-macro machinery so it is easy to
//! test (`Spec` -> IR).

/// The whole generated CLI tree.
#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedCli {
    /// One entry per configured group, in hierarchy order.
    pub groups: Vec<GeneratedGroup>,
}

/// A top-level group, from one entry in the manual `cli.yaml` hierarchy.
#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedGroup {
    /// The group name exactly as configured, e.g. `"Cloud Backups"`.
    pub name: String,
    /// PascalCase identifier derived from the name, e.g. `"CloudBackups"`.
    pub ident: String,
    /// Help text; falls back to a template in `emit`.
    pub description: Option<String>,
    /// Entities under this group, in hierarchy order.
    pub entities: Vec<GeneratedEntity>,
}

/// One command entity inside a group.
///
/// A group with a single entity is emitted flattened: its operations become
/// the group's subcommands directly. A group with several entities emits one
/// per-entity command underneath the group command.
#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedEntity {
    /// Entity name, e.g. `"Cluster"`.
    pub id: String,
    pub ident: String,
    /// Help text; falls back to a template in `emit`.
    pub description: Option<String>,
    /// Direct operations (create/read/update/delete/list + actions), in hierarchy order.
    pub operations: Vec<GeneratedOperation>,
    /// Nested component commands, in hierarchy order.
    pub components: Vec<GeneratedComponent>,
}

/// A component command inside an entity (`advanced-configuration-options …`).
#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedComponent {
    /// Component name, e.g. `"AdvancedConfigurationOptions"`.
    pub id: String,
    pub ident: String,
    /// Help text; falls back to a template in `emit`.
    pub description: Option<String>,
    pub operations: Vec<GeneratedOperation>,
}

/// How an operation is categorized in help output: CRUD lifecycle verbs
/// (`create`/`read`/`update`/`delete`/`list`) vs custom actions.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum OperationKind {
    Crud,
    Action,
}

/// One operation, expanded into a probe + version enum + a struct per version.
#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedOperation {
    /// Subcommand variant name, e.g. `"Create"`, `"GetCompliancePolicy"`.
    pub variant_name: String,
    /// First line of the operation description, used as the doc comment.
    pub description: String,
    /// Which help section this operation lands in.
    pub kind: OperationKind,
    /// e.g. `"CreateGroupClusterProbe"`.
    pub probe_ident: String,
    /// e.g. `"CreateGroupClusterVersion"`.
    pub version_enum_ident: String,
    /// Oldest first; the last entry is the default `--version`.
    pub versions: Vec<GeneratedVersion>,
    /// Path/query/header parameters. Shared by every version struct.
    pub flags: Vec<GeneratedFlag>,
    /// HTTP method, e.g. `"GET"`.
    pub method: String,
    /// URL path with `{parameterName}` placeholders, e.g.
    /// `/api/atlas/v2/groups/{groupId}/clusters`.
    pub url_template: String,
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedVersion {
    /// ValueEnum variant name, e.g. `"V20241023"`.
    pub variant_ident: String,
    /// clap::Parser struct name, e.g. `"CreateGroupClusterV20241023"`.
    pub struct_ident: String,
    /// Flat flags derived from this version's request body (per-version, the
    /// shape can change between versions). Empty when the body is absent or
    /// not a supported simple shape.
    pub body_flags: Vec<GeneratedFlag>,
    /// The JSON Schema (`serde_json::Value`, stringified) of this version's
    /// request body, used to validate `--file`/stdin input and flag-built
    /// bodies. `None` when this version has no request body.
    pub body_schema: Option<String>,
    /// The atlas API version this variant maps to.
    pub api_version: ApiVersion,
}

/// What the HTTP request looks like for one API version of an operation.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ApiVersion {
    Stable(u32, u32, u32),
    Upcoming(u32, u32, u32),
    Preview,
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct GeneratedFlag {
    /// Original OpenAPI parameter name, e.g. `"groupId"`.
    pub name: String,
    /// Struct field identifier derived from the name.
    pub ident: String,
    pub description: String,
    pub required: bool,
    /// Repeated flag: `--tag a --tag b` -> `Vec<String>`.
    pub list: bool,
    /// Where the parameter lands in a request.
    pub location: FlagLocation,
    /// JSON scalar kind of a flat body flag, so flag-built bodies carry the
    /// right JSON types (`copyProtectionEnabled` -> `true`, not `"true"`).
    pub value_kind: FlagValueKind,
}

/// The JSON scalar kind of a flat body flag's leaf value type.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum FlagValueKind {
    String,
    Integer,
    Double,
    Boolean,
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum FlagLocation {
    Path,
    Query,
    Header,
    Body,
}
