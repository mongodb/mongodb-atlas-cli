use std::{collections::BTreeMap, str::FromStr};

use models::{
    http_verb::{Verb, VerbError},
    operation_id::{OperationId, OperationIdParseError},
    versioned_mediatype::{
        MediaType, Version, VersionedAcceptHeader, VersionedAcceptHeaderParseError,
    },
};
use openapiv3_resolve::{
    ResolvedMediaType, ResolvedOperation, ResolvedParameter, Shared, indexmap::IndexMap,
    openapiv3::StatusCode,
};
use serde::{Deserialize, Serialize};
use thiserror::Error;

mod headers;
mod url;
mod version;

use headers::Headers;
use url::ParameterizedUrl;
use version::OperationVersion;

use crate::operation::{
    headers::HeadersParseError, url::ParameterizedUrlParseError,
    version::OperationVersionParseError,
};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Operation {
    pub description: String,
    pub operation_id: OperationId,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub operation_id_override: Option<OperationId>,
    pub http_verb: Verb,
    pub url: ParameterizedUrl,
    pub headers: Headers,

    pub versions: BTreeMap<Version, OperationVersion>,
}

#[derive(Debug, Error)]
pub enum OperationParseError {
    #[error("Description is missing")]
    MissingDescription,
    #[error("OperationID is missing")]
    MissingOperationId,
    #[error(transparent)]
    InvalidOperationId(#[from] OperationIdParseError),
    #[error("Invalid OperationId override, wrong type")]
    InvalidOperationIdOverrideType,
    #[error(transparent)]
    InvalidOperationIdOverride(OperationIdParseError),
    #[error(transparent)]
    InvalidHTTPVerb(#[from] VerbError),
    #[error(transparent)]
    ParameterizedUrlParseError(#[from] ParameterizedUrlParseError),
    #[error(transparent)]
    HeadersParseError(#[from] HeadersParseError),
    #[error("Cookie parameters are not supported")]
    UnsupportedCookieParameter,
    #[error(transparent)]
    VersionedAcceptHeaderParseError(#[from] VersionedAcceptHeaderParseError),
    #[error(transparent)]
    OperationVersionParseError(#[from] OperationVersionParseError),
}

impl Operation {
    pub fn from_path_verb_and_resolved_operation(
        path: &str,
        verb: &str,
        operation: &ResolvedOperation,
    ) -> Result<Self, OperationParseError> {
        let description = operation
            .description
            .clone()
            .and_then(|s| (!s.is_empty()).then_some(s))
            .ok_or(OperationParseError::MissingDescription)?;

        let operation_id = operation
            .operation_id
            .as_deref()
            .ok_or(OperationParseError::MissingOperationId)?
            .parse()?;

        let operation_id_override = operation
            .extensions
            .get("x-xgen-operation-id-override")
            .map(|v| match v.as_str() {
                Some(s) => s
                    .parse::<OperationId>()
                    .map_err(OperationParseError::InvalidOperationIdOverride),
                None => Err(OperationParseError::InvalidOperationIdOverrideType),
            })
            .transpose()?;

        let http_verb = Verb::try_new(verb)?;

        Self::reject_cookie_parameters(&operation.parameters)?;

        let url = ParameterizedUrl::from_path_and_resolved_parameters(path, &operation.parameters)?;
        let headers = Headers::from_resolved_parameters(&operation.parameters)?;

        let mut version_requests = if let Some(request_body) = operation.request_body.as_deref() {
            Self::resolved_media_types_to_versions(&request_body.content)?
        } else {
            BTreeMap::new()
        };

        let mut version_responses =
            BTreeMap::<Version, BTreeMap<MediaType, &ResolvedMediaType>>::new();

        for (status_code, response) in &operation.responses.responses {
            let StatusCode::Code(code) = *status_code else {
                continue;
            };

            if !(200..=299).contains(&code) {
                continue;
            }

            let local_version_responses =
                Self::resolved_media_types_to_versions(&response.content)?;

            for (version, versioned_response) in local_version_responses {
                for (content_type, resolved_media_type) in versioned_response {
                    let version_response = version_responses.entry(version).or_default();
                    version_response.insert(content_type, resolved_media_type);
                }
            }
        }

        let mut operation_versions = BTreeMap::new();

        for (version, response_bodies) in version_responses {
            let request_bodies = version_requests.remove(&version);
            let operation_version = OperationVersion::from_version_requests_responses(
                version,
                request_bodies,
                response_bodies,
            )?;

            operation_versions.insert(version, operation_version);
        }

        Ok(Operation {
            description,
            operation_id,
            operation_id_override,
            http_verb,
            url,
            headers,
            versions: operation_versions,
        })
    }

    fn reject_cookie_parameters(
        parameters: &Vec<Shared<ResolvedParameter>>,
    ) -> Result<(), OperationParseError> {
        if parameters
            .iter()
            .any(|p| matches!(&**p, ResolvedParameter::Cookie { .. }))
        {
            return Err(OperationParseError::UnsupportedCookieParameter);
        }
        Ok(())
    }

    fn resolved_media_types_to_versions<'a>(
        content: &'a IndexMap<String, ResolvedMediaType>,
    ) -> Result<
        BTreeMap<Version, BTreeMap<MediaType, &'a ResolvedMediaType>>,
        VersionedAcceptHeaderParseError,
    > {
        let mut version_requests =
            BTreeMap::<Version, BTreeMap<MediaType, &ResolvedMediaType>>::new();
        for (content_type, resolved_media_type) in content.iter() {
            let Ok(versioned_accept_header) = VersionedAcceptHeader::from_str(content_type) else {
                // Legacy content types (e.g. `application/json`) have no
                // version to attribute an OperationVersion to.
                eprintln!("Skipping content type without a version: {content_type}");
                continue;
            };
            let version = version_requests
                .entry(versioned_accept_header.version)
                .or_default();

            let media_type = versioned_accept_header.mediatype;

            version.insert(media_type, resolved_media_type);
        }

        Ok(version_requests)
    }
}
