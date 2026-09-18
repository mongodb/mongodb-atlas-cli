//! Intermediate layer between the OpenAPI spec, a manual CLI hierarchy, and
//! the `atlas-cli-derive` proc macro, so the macro stays dumb and the
//! interesting logic is testable.
//!
//! Pipeline: `spec.yaml + cli.yaml -> ResolvedOpenAPI + models::Hierarchy ->
//! GeneratedCli (this crate) -> TokenStream`.
#![cfg_attr(not(test), deny(clippy::unwrap_used, clippy::expect_used))]

use std::collections::BTreeSet;

use thiserror::Error;

pub mod emit;
pub mod from_spec;
pub mod ir;

pub use emit::to_tokens;
pub use from_spec::from_config;
pub use ir::{
    ApiVersion, FlagLocation, GeneratedCli, GeneratedComponent, GeneratedEntity, GeneratedFlag,
    GeneratedGroup, GeneratedOperation, GeneratedVersion,
};

/// Knobs for the `#[atlas_cli(...)]` attribute on the derived `Cli` struct.
#[derive(Debug, Clone, Default)]
pub struct CodegenOptions {
    /// Spec `operationId`s to skip even when the hierarchy references them.
    pub excluded_operation_ids: BTreeSet<String>,
}

#[derive(Debug, Error)]
pub enum CodegenError {
    #[error("invalid or unresolvable OpenAPI spec: {0}")]
    InvalidOpenAPISpec(String),
    #[error("invalid hierarchy yaml: {0}")]
    InvalidHierarchy(#[from] serde_yaml::Error),
    #[error("invalid spec: {0}")]
    InvalidSpec(#[from] models::SpecParseError),
    #[error("operation `{op_id}` is referenced in the hierarchy but not in the spec")]
    OperationNotFoundInSpec { op_id: String },
    #[error(
        "{count} operation(s) are covered neither by the hierarchy nor by `excluded_operation_ids`: {ids}",
        count = operation_ids.len(),
        ids = operation_ids.join(", ")
    )]
    UncoveredOperations { operation_ids: Vec<String> },
}
