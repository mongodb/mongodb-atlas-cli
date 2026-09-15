use std::collections::HashMap;

use openapiv3_resolve::{ResolvedParameter, ResolvedParameterData, Shared, openapiv3::PathStyle};
use serde::{Deserialize, Serialize};
use thiserror::Error;

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct ParameterizedUrl {
    pub parts: Vec<Part>,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub enum Part {
    Const(String),
    Parameter(UrlParameter),
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct UrlParameter {
    pub description: String,
    pub name: String,
}

#[derive(Debug, Error)]
pub enum UrlParameterParseError {
    #[error(
        "Only PathStyle::Simple is supported for parameters, got '{path_style:?}' instead for {parameter_name}"
    )]
    InvalidPathStyle {
        path_style: PathStyle,
        parameter_name: String,
    },
    #[error("Only required header parameters are supported, '{parameter_name}' is not required")]
    ParameterNotRequired { parameter_name: String },
    #[error("Description is missing")]
    MissingDescription,
}

impl UrlParameter {
    fn from_parameter_data_path_style(
        parameter_data: &ResolvedParameterData,
        style: &PathStyle,
    ) -> Result<Self, UrlParameterParseError> {
        let PathStyle::Simple = style else {
            return Err(UrlParameterParseError::InvalidPathStyle {
                path_style: style.clone(),
                parameter_name: parameter_data.name.to_owned(),
            });
        };

        if !parameter_data.required {
            return Err(UrlParameterParseError::ParameterNotRequired {
                parameter_name: parameter_data.name.to_owned(),
            });
        }

        let name = parameter_data.name.to_owned();
        let description = parameter_data
            .description
            .clone()
            .and_then(|s| (!s.is_empty()).then_some(s))
            .ok_or(UrlParameterParseError::MissingDescription)?;

        Ok(Self { name, description })
    }
}

#[derive(Debug, Error)]
pub enum ParameterizedUrlParseError {
    #[error(transparent)]
    UrlParameterParseError(#[from] UrlParameterParseError),
    #[error(
        "The parameter '{parameter_name}' was used in the path, but not found in the parameter list"
    )]
    MissingParameter { parameter_name: String },
}

impl ParameterizedUrl {
    pub fn from_path_and_resolved_parameters(
        path: &str,
        parameters: &Vec<Shared<ResolvedParameter>>,
    ) -> Result<Self, ParameterizedUrlParseError> {
        let mut url_parameters = parameters
            .iter()
            .filter_map(|p| match &**p {
                ResolvedParameter::Path {
                    parameter_data,
                    style,
                } => Some(
                    UrlParameter::from_parameter_data_path_style(parameter_data, style)
                        .map(|p| (p.name.clone(), p)),
                ),
                _ => None,
            })
            .collect::<Result<HashMap<String, UrlParameter>, UrlParameterParseError>>()?;

        let mut parts = Vec::new();
        for part in path.trim_start_matches('/').split('/') {
            if let Some(parameter_name) = part.strip_prefix('{').and_then(|s| s.strip_suffix('}')) {
                let Some(parameter) = url_parameters.remove(parameter_name) else {
                    return Err(ParameterizedUrlParseError::MissingParameter {
                        parameter_name: parameter_name.to_owned(),
                    });
                };

                parts.push(Part::Parameter(parameter));
            } else {
                parts.push(Part::Const(part.to_string()))
            }
        }

        Ok(Self { parts })
    }
}

#[cfg(test)]
pub mod tests {
    use openapiv3_resolve::{ResolvedOpenAPI, ResolvedParameter, Shared, openapiv3::OpenAPI};

    use super::*;

    const SPEC: &str = r#"
openapi: 3.0.3
info:
  title: test
  version: "1"
paths:
  "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements:":
    get:
      parameters:
        - { name: groupId, in: path, required: true, description: Group id, schema: { type: string } }
        - { name: clusterName, in: path, required: true, description: Cluster name, schema: { type: string } }
        - { name: clusterView, in: path, required: true, description: Cluster view, schema: { type: string } }
        - { name: databaseName, in: path, required: true, description: Database name, schema: { type: string } }
        - { name: collectionName, in: path, required: true, description: Collection name, schema: { type: string } }
      responses:
        "200":
          description: ok
  "/{groupId}/{clusterName}/{databaseName}":
    get:
      parameters:
        - { name: groupId, in: path, required: true, description: Group id, schema: { type: string } }
        - { name: clusterName, in: path, required: true, description: Cluster name, schema: { type: string } }
        - { name: databaseName, in: path, required: true, description: Database name, schema: { type: string } }
      responses:
        "200":
          description: ok
"#;

    fn path_parameter(name: &str, description: &str) -> UrlParameter {
        UrlParameter {
            description: description.to_owned(),
            name: name.to_owned(),
        }
    }

    // Shared values only exist inside a resolved document; the doc must outlive the borrows.
    fn resolved_parameters<'a>(
        doc: &'a ResolvedOpenAPI,
        path: &str,
    ) -> &'a Vec<Shared<ResolvedParameter>> {
        &doc.paths().paths[path].get.as_ref().unwrap().parameters
    }

    #[test]
    fn parses_and_interleaves_const_and_parameter_parts() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/{clusterView}/{databaseName}/{collectionName}/collStats/measurements:";
        let url = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap();

        assert_eq!(
            url.parts,
            [
                Part::Const(String::from("api")),
                Part::Const(String::from("atlas")),
                Part::Const(String::from("v2")),
                Part::Const(String::from("groups")),
                Part::Parameter(path_parameter("groupId", "Group id")),
                Part::Const(String::from("clusters")),
                Part::Parameter(path_parameter("clusterName", "Cluster name")),
                Part::Parameter(path_parameter("clusterView", "Cluster view")),
                Part::Parameter(path_parameter("databaseName", "Database name")),
                Part::Parameter(path_parameter("collectionName", "Collection name")),
                Part::Const(String::from("collStats")),
                Part::Const(String::from("measurements:")),
            ]
        );
    }

    #[test]
    fn parses_all_dynamic_path() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/{groupId}/{clusterName}/{databaseName}";
        let url = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap();

        assert_eq!(
            url.parts,
            [
                Part::Parameter(path_parameter("groupId", "Group id")),
                Part::Parameter(path_parameter("clusterName", "Cluster name")),
                Part::Parameter(path_parameter("databaseName", "Database name")),
            ]
        );
    }

    #[test]
    fn parses_all_const_path_without_parameters() {
        let url =
            ParameterizedUrl::from_path_and_resolved_parameters("/api/atlas/v2/ready", &vec![])
                .unwrap();

        assert_eq!(
            url.parts,
            [
                Part::Const(String::from("api")),
                Part::Const(String::from("atlas")),
                Part::Const(String::from("v2")),
                Part::Const(String::from("ready")),
            ]
        );
    }

    #[test]
    fn parses_empty_path() {
        let url = ParameterizedUrl::from_path_and_resolved_parameters("", &vec![]).unwrap();

        assert_eq!(url.parts, [Part::Const(String::from(""))]);
    }

    #[test]
    fn errors_when_parameter_is_missing_from_parameter_list() {
        let err =
            ParameterizedUrl::from_path_and_resolved_parameters("/groups/{missingName}", &vec![])
                .unwrap_err();

        assert!(matches!(
            err,
            ParameterizedUrlParseError::MissingParameter { parameter_name } if parameter_name == "missingName"
        ));
    }
}
