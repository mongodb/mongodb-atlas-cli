use std::collections::HashMap;

use openapiv3_resolve::{ResolvedParameter, ResolvedParameterData, Shared};
use serde::{Deserialize, Serialize};
use thiserror::Error;

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Headers {
    pub parameters: HashMap<String, HeaderParameter>,
}

#[derive(Debug, Error)]
pub enum HeadersParseError {
    #[error(transparent)]
    HeaderParameterParseError(#[from] HeaderParameterParseError),
}

impl Headers {
    pub fn from_resolved_parameters(
        resolved_parameters: &Vec<Shared<ResolvedParameter>>,
    ) -> Result<Headers, HeadersParseError> {
        let parameters = resolved_parameters
            .iter()
            .filter_map(|p| match &**p {
                ResolvedParameter::Header { parameter_data, .. } => Some(
                    HeaderParameter::from_resolved_parameter_data(parameter_data)
                        .map(|parameter| (parameter.name.clone(), parameter)),
                ),
                _ => None,
            })
            .collect::<Result<HashMap<String, HeaderParameter>, _>>()?;

        Ok(Self { parameters })
    }
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct HeaderParameter {
    pub description: String,
    pub name: String,
    pub required: bool,
}

#[derive(Debug, Error)]
pub enum HeaderParameterParseError {
    #[error("Description is missing")]
    MissingDescription,
}

impl HeaderParameter {
    pub fn from_resolved_parameter_data(
        parameter_data: &ResolvedParameterData,
    ) -> Result<Self, HeaderParameterParseError> {
        let description = parameter_data
            .description
            .clone()
            .and_then(|s| (!s.is_empty()).then_some(s))
            .ok_or(HeaderParameterParseError::MissingDescription)?;

        let name = parameter_data.name.clone();
        let required = parameter_data.required;

        Ok(Self {
            description,
            name,
            required,
        })
    }
}
