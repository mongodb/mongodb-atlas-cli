use models::Spec;
use openapiv3_resolve::openapiv3::OpenAPI;
use openapiv3_resolve::ResolvedOpenAPI;

#[test]
fn parse_openapi_spec_in_repo() {
    let openapi_spec_str =
        include_str!("../../../tools/internal/specs/spec-with-overlays.yaml");
    let openapi_spec: OpenAPI =
        serde_yaml::from_str(openapi_spec_str).expect("valid openapi spec");
    let resolved_openapi_spec =
        ResolvedOpenAPI::try_from(&openapi_spec).expect("all references are valid");

    let spec =
        Spec::from_resolved_openapi_spec(&resolved_openapi_spec).expect("parsing succeeds");
    insta::assert_json_snapshot!(spec);
}
