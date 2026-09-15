use std::collections::HashMap;

use models::{
    http_verb::{Verb, VerbError},
    operation_id::{OperationId, OperationIdParseError},
    versioned_mediatype::Version,
};
use openapiv3_resolve::{ResolvedOperation, ResolvedParameter, Shared};
use serde::{Deserialize, Serialize};
use thiserror::Error;

mod headers;
mod url;
mod version;

use headers::Headers;
use url::ParameterizedUrl;
use version::OperationVersion;

use crate::operation::{
    headers::{HeaderParameterParseError, HeadersParseError},
    url::ParameterizedUrlParseError,
};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Operation {
    pub description: String,
    pub operation_id: OperationId,
    pub operation_id_override: Option<OperationId>,
    pub http_verb: Verb,
    pub url: ParameterizedUrl,
    pub headers: Headers,

    pub versions: HashMap<Version, OperationVersion>,
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

        Ok(Operation {
            description,
            operation_id,
            operation_id_override,
            http_verb,
            url,
            headers,
            versions: Default::default(),
        })
    }

    fn reject_cookie_parameters(
        parameters: &Vec<Shared<ResolvedParameter>>,
    ) -> Result<(), OperationParseError> {
        if parameters.iter().any(|p| {
            matches!(
                &**p,
                ResolvedParameter::Cookie {
                    ..
                }
            )
        }) {
            return Err(OperationParseError::UnsupportedCookieParameter);
        }
        Ok(())
    }
}
