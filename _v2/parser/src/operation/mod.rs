use std::collections::HashMap;

use models::{http_verb::Verb, operation_id::OperationId, versioned_mediatype::Version};
use openapiv3_resolve::ResolvedOperation;
use thiserror::Error;

use crate::{ParameterizedUrl, parameter::Parameter};
use version::OperationVersion;

mod version;

pub struct Operation {
    pub description: String,
    pub operation_id: OperationId,
    pub overriden_operation_id: Option<OperationId>,
    pub http_verb: Verb,
    pub url: ParameterizedUrl,
    pub header_parameters: HashMap<String, Parameter>,
    pub versions: HashMap<Version, OperationVersion>,
}

#[derive(Debug, Error)]
pub enum OperationParseError {}

impl Operation {
    pub fn from_path_verb_and_resolved_operation(
        path: &str,
        verb: &str,
        operation: &ResolvedOperation,
    ) -> Result<Self, OperationParseError> {
        todo!()
    }
}
