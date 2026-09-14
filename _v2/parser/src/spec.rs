use std::collections::BTreeMap;

use models::operation_id::OperationId;
use openapiv3_resolve::ResolvedOpenAPI;
use thiserror::Error;

use crate::{Operation, OperationParseError};

pub struct Spec {
    pub operations: BTreeMap<OperationId, Operation>,
}

#[derive(Debug, Error)]
pub enum SpecParseError {
    #[error("Failed to parse operation (PATH={path}, VERB={verb}): {error}")]
    FailedToParseOperation {
        path: String,
        verb: String,
        error: OperationParseError,
    },
}

impl Spec {
    pub fn from_resolved_openapi_spec(spec: &ResolvedOpenAPI) -> Result<Self, SpecParseError> {
        let mut operations = BTreeMap::new();

        for (path, item) in spec.paths().paths.iter() {
            for (verb, operation) in item.iter() {
                let operation =
                    Operation::from_path_verb_and_resolved_operation(path, verb, operation)
                        .map_err(|error| SpecParseError::FailedToParseOperation {
                            path: path.to_owned(),
                            verb: verb.to_owned(),
                            error,
                        })?;

                let operation_id = operation.operation_id.clone();
                operations.insert(operation_id, operation);
            }
        }

        Ok(Self { operations })
    }
}
