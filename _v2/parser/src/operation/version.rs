use std::collections::HashMap;

use models::{
    datatypes::DataType,
    versioned_mediatype::{MediaType, Version},
};
use openapiv3_resolve::ResolvedMediaType;
use serde::{Deserialize, Serialize};
use thiserror::Error;

/// A version of an operation (determined by version header, example: application/vnd.atlas.2024-10-23+json)
/// A version can have multiple response types, for example `getOrgBillingCostExplorerUsage` can both return csv and json, different schemas
///
/// Limitations:
/// - we only support JSON requests, so maximum 1 request body
#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct OperationVersion {
    pub request_body: Option<DataType>,
    pub responses: HashMap<MediaType, DataType>,
}

#[derive(Debug, Error)]
pub enum OperationVersionParseError {}

impl OperationVersion {
    pub fn from_version_requests_responses(
        version: Version,
        request_bodies: Option<HashMap<MediaType, &ResolvedMediaType>>,
        response_bodies: HashMap<MediaType, &ResolvedMediaType>,
    ) -> Result<Self, OperationVersionParseError> {
        todo!()
    }
}
