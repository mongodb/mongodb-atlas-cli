//! Integration test for [`models::datatypes::DataType::to_json_schema`]
//! against the real Atlas Admin API spec.
//!
//! The `parser` dev-dependency turns the repo's OpenAPI spec into
//! [`DataType`]s (single source of truth for the conversion), then this crate
//! snapshots the JSON Schemas it produces. Lives here so the snapshot sits in
//! the crate under test, next to the code.

use models::datatypes::json_schema::is_valid;
use models::operation_id::OperationId;
use openapiv3_resolve::openapiv3::OpenAPI;

/// `createGroupCluster`: one JSON Schema per request body and per response
/// body (a snapshot file each, named `<operation>_<version>_request` /
/// `<operation>_<version>_response_<media_type>`).
#[test]
fn create_group_cluster_json_schemas() {
    let openapi_spec_str = include_str!("../../../tools/internal/specs/spec-with-overlays.yaml");
    let openapi_spec: OpenAPI =
        serde_yaml::from_str(openapi_spec_str).expect("valid openapi spec");
    let resolved_openapi_spec =
        openapiv3_resolve::ResolvedOpenAPI::try_from(&openapi_spec)
            .expect("all references are valid");
    let spec =
        parser::Spec::from_resolved_openapi_spec(&resolved_openapi_spec).expect("parsing succeeds");

    let operation_id: OperationId = "createGroupCluster".parse().unwrap();
    let operation = spec.operations.get(&operation_id).expect("operation exists");

    for (version, operation_version) in &operation.versions {
        if let Some(request_body) = &operation_version.request_body {
            let schema = request_body.to_json_schema();
            assert!(is_valid(&schema));
            let name = format!("create_group_cluster_{version}_request");
            insta::assert_json_snapshot!(name, schema);
        }
        for (media_type, response_body) in &operation_version.response_bodies {
            let schema = response_body.to_json_schema();
            assert!(is_valid(&schema));
            let name = format!("create_group_cluster_{version}_response_{media_type}");
            insta::assert_json_snapshot!(name, schema);
        }
    }
}
