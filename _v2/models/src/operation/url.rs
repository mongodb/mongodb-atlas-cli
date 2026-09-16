use std::collections::BTreeMap;

use openapiv3_resolve::{ResolvedParameter, ResolvedParameterData, Shared, openapiv3::{PathStyle, QueryStyle}};
use serde::{Deserialize, Serialize};
use thiserror::Error;

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct ParameterizedUrl {
    pub parts: Vec<Part>,
    pub query_parameters: Vec<UrlParameter>,
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
    pub required: bool,
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
    #[error(
        "Only QueryStyle::Form is supported for parameters, got '{query_style:?}' instead for {parameter_name}"
    )]
    InvalidQueryStyle {
        query_style: QueryStyle,
        parameter_name: String,
    },
    #[error("Only required header parameters are supported, '{parameter_name}' is not required")]
    ParameterNotRequired { parameter_name: String },
    #[error(
        "Only 'none' or 'false' is supported for allow_empty_value, got 'true' instead for '{parameter_name}'"
    )]
    AllowEmptyValue { parameter_name: String },
    #[error("Description is missing")]
    MissingDescription,
}

impl UrlParameter {
    fn from_parameter_data(
        name: String,
        description: Option<String>,
        required: bool,
    ) -> Result<Self, UrlParameterParseError> {
        let description = description
            .and_then(|s| (!s.is_empty()).then_some(s))
            .ok_or(UrlParameterParseError::MissingDescription)?;

        Ok(Self {
            name,
            description,
            required,
        })
    }

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

        Self::from_parameter_data(
            parameter_data.name.to_owned(),
            parameter_data.description.clone(),
            parameter_data.required,
        )
    }

    fn from_parameter_data_query_style(
        parameter_data: &ResolvedParameterData,
        style: &QueryStyle,
        allow_empty_value: Option<bool>,
    ) -> Result<Self, UrlParameterParseError> {
        let QueryStyle::Form = style else {
            return Err(UrlParameterParseError::InvalidQueryStyle {
                query_style: style.clone(),
                parameter_name: parameter_data.name.to_owned(),
            });
        };

        if allow_empty_value == Some(true) {
            return Err(UrlParameterParseError::AllowEmptyValue {
                parameter_name: parameter_data.name.to_owned(),
            });
        }

        Self::from_parameter_data(
            parameter_data.name.to_owned(),
            parameter_data.description.clone(),
            parameter_data.required,
        )
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
            .collect::<Result<BTreeMap<String, UrlParameter>, UrlParameterParseError>>()?;

        let query_parameters = parameters
            .iter()
            .filter_map(|p| match &**p {
                ResolvedParameter::Query {
                    parameter_data,
                    style,
                    allow_empty_value,
                    ..
                } => Some(UrlParameter::from_parameter_data_query_style(
                    parameter_data,
                    style,
                    *allow_empty_value,
                )),
                _ => None,
            })
            .collect::<Result<Vec<UrlParameter>, UrlParameterParseError>>()?;

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

        Ok(Self {
            parts,
            query_parameters,
        })
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
components:
  parameters:
    envelope:
      description: Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
      in: query
      name: envelope
      schema:
        default: false
        type: boolean
    federationSettingsId:
      description: Unique 24-hexadecimal digit string that identifies your federation.
      in: path
      name: federationSettingsId
      required: true
      schema:
        example: 55fa922fb343282757d9554e
        pattern: ^([a-f0-9]{24})$
        type: string
    groupId:
      description: |-
        Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

        **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
      in: path
      name: groupId
      required: true
      schema:
        example: 32b6e34b3d91647abb20e7b8
        pattern: ^([a-f0-9]{24})$
        type: string
    itemsPerPage:
      description: Number of items that the response returns per page.
      in: query
      name: itemsPerPage
      schema:
        default: 100
        maximum: 500
        minimum: 1
        type: integer
    pageNum:
      description: Number of the page that displays the current set of the total objects that the response returns.
      in: query
      name: pageNum
      schema:
        default: 1
        minimum: 1
        type: integer
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
  "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders":
    get:
      parameters:
        - $ref: '#/components/parameters/federationSettingsId'
        - $ref: '#/components/parameters/envelope'
        - $ref: '#/components/parameters/itemsPerPage'
        - $ref: '#/components/parameters/pageNum'
        - description: The protocols of the target identity providers.
          in: query
          name: protocol
          schema:
            items:
              default: SAML
              enum:
                - SAML
                - OIDC
              type: string
            type: array
        - description: The types of the target identity providers.
          in: query
          name: idpType
          schema:
            items:
              default: WORKFORCE
              enum:
                - WORKFORCE
                - WORKLOAD
              type: string
            type: array
      responses:
        "200":
          description: ok
  "/api/atlas/v2/groups/{groupId}/streams/accountDetails":
    get:
      parameters:
        - $ref: '#/components/parameters/groupId'
        - $ref: '#/components/parameters/envelope'
        - description: One of "aws", "azure" or "gcp".
          in: query
          name: cloudProvider
          required: true
          schema:
            type: string
        - description: The cloud provider specific region name, i.e. "US_EAST_1" for cloud provider "aws".
          in: query
          name: regionName
          required: true
          schema:
            type: string
      responses:
        "200":
          description: ok
  # Synthetic paths below: the real Atlas spec has no non-form query styles and no
  # allowEmptyValue=true, so these are used to exercise those validation rules only.
  "/query-style/{id}":
    get:
      parameters:
        - { name: id, in: path, required: true, description: Id, schema: { type: string } }
        - { name: filter, in: query, style: deepObject, description: Filter, schema: { type: object } }
      responses:
        "200":
          description: ok
  "/query-empty/{id}":
    get:
      parameters:
        - { name: id, in: path, required: true, description: Id, schema: { type: string } }
        - { name: emptyVal, in: query, allowEmptyValue: true, description: Empty value, schema: { type: string } }
      responses:
        "200":
          description: ok
"#;

    fn path_parameter(name: &str, description: &str) -> UrlParameter {
        UrlParameter {
            description: description.to_owned(),
            name: name.to_owned(),
            required: true,
        }
    }

    fn query_parameter(name: &str, description: &str, required: bool) -> UrlParameter {
        UrlParameter {
            description: description.to_owned(),
            name: name.to_owned(),
            required,
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

    fn query_parameter_names_and_required(url: &ParameterizedUrl) -> Vec<(String, bool)> {
        url.query_parameters
            .iter()
            .map(|p| (p.name.clone(), p.required))
            .collect()
    }

    #[test]
    fn parses_real_required_query_parameters() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/api/atlas/v2/groups/{groupId}/streams/accountDetails";
        let url = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap();

        assert_eq!(
            query_parameter_names_and_required(&url),
            [
                (String::from("envelope"), false),
                (String::from("cloudProvider"), true),
                (String::from("regionName"), true),
            ]
        );
    }

    #[test]
    fn parses_real_optional_query_parameters() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/api/atlas/v2/federationSettings/{federationSettingsId}/identityProviders";
        let url = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap();

        assert_eq!(
            query_parameter_names_and_required(&url),
            [
                (String::from("envelope"), false),
                (String::from("itemsPerPage"), false),
                (String::from("pageNum"), false),
                (String::from("protocol"), false),
                (String::from("idpType"), false),
            ]
        );
        assert_eq!(
            url.query_parameters[0],
            query_parameter(
                "envelope",
                "Flag that indicates whether Application wraps the response in an `envelope` JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope=true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.",
                false,
            )
        );
    }

    #[test]
    fn errors_on_non_form_query_style() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/query-style/{id}";
        let err = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap_err();

        assert!(matches!(
            err,
            ParameterizedUrlParseError::UrlParameterParseError(
                UrlParameterParseError::InvalidQueryStyle { parameter_name, .. }
            ) if parameter_name == "filter"
        ));
    }

    #[test]
    fn errors_on_allow_empty_value_true() {
        let openapi: OpenAPI = serde_yaml::from_str(SPEC).unwrap();
        let doc = ResolvedOpenAPI::try_from(&openapi).unwrap();

        let path = "/query-empty/{id}";
        let err = ParameterizedUrl::from_path_and_resolved_parameters(
            path,
            resolved_parameters(&doc, path),
        )
        .unwrap_err();

        assert!(matches!(
            err,
            ParameterizedUrlParseError::UrlParameterParseError(
                UrlParameterParseError::AllowEmptyValue { parameter_name }
            ) if parameter_name == "emptyVal"
        ));
    }
}
