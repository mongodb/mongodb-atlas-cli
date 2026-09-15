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
    pub version: Version,
    pub request_body: Option<DataType>,
    pub response_bodies: HashMap<MediaType, DataType>,
}

#[derive(Debug, Error)]
pub enum OperationVersionParseError {
    #[error("Only JSON request bodies are supported")]
    OnlyJsonRequestBodies,
}

impl OperationVersion {
    pub fn from_version_requests_responses(
        version: Version,
        request_bodies: Option<HashMap<MediaType, &ResolvedMediaType>>,
        response_bodies: HashMap<MediaType, &ResolvedMediaType>,
    ) -> Result<Self, OperationVersionParseError> {
        // Parse the JSON request body
        let request_body_datatype = if let Some(request_bodies) = request_bodies {
            if let Some(resolved_media_type) = request_bodies.get(&MediaType::Json) {
                Some(Self::resolved_media_type_to_datatype(resolved_media_type)?)
            } else {
                return Err(OperationVersionParseError::OnlyJsonRequestBodies);
            }
        } else {
            None
        };

        // Parse all response bodies
        let mut response_bodies_datatypes = HashMap::new();
        for (media_type, resolved_media_type) in response_bodies {
            let data_type = Self::resolved_media_type_to_datatype(resolved_media_type)?;
            response_bodies_datatypes.insert(media_type, data_type);
        }

        Ok(Self {
            version,
            request_body: request_body_datatype,
            response_bodies: response_bodies_datatypes,
        })
    }

    fn resolved_media_type_to_datatype(
        _resolved_media_type: &ResolvedMediaType,
    ) -> Result<DataType, OperationVersionParseError> {
        // Conversion rules:
        // - AnyOf = Object with Optional fields, so AnyOf should be converted into Object with Optional
        // - AllOf = merge into Object
        // - OneOf exists
        //
        // Caveats:
        // - important with AnyOf -> our OpenApi spec does have schemas with enum types with 1 entry that should be merged through anyOf and should become the discriminator
        // - integer can be either double or integer based on the field `double: true`
        todo!()
    }
}
