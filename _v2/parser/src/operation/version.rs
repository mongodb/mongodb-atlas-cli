use std::collections::HashMap;

use models::versioned_mediatype::Version;
use openapiv3_resolve::ResolvedMediaType;
use serde::{Deserialize, Serialize};
use thiserror::Error;

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct OperationVersion {
    pub request: (),
    pub responses: (),
}

#[derive(Debug, Error)]
pub enum OperationVersionParseError {}

impl OperationVersion {
    pub fn from_version_requests_responses(
        version: Version,
        request_bodies: Option<HashMap<String, &ResolvedMediaType>>,
        response_bodies: HashMap<String, &ResolvedMediaType>,
    ) -> Result<Self, OperationVersionParseError> {
        todo!()
    }
}
