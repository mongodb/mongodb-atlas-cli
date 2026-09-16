//! Snapshot tests for the quirks documented in
//! `auto-generation-v2/quirks` (the marimo notebook `quirks.py`).
//!
//! Each quirk is reproduced as a minimal OpenAPI doc under `parser/fixtures/`
//! and converted through the real pipeline (resolve -> Operation -> datatype).
//! The snapshot shows the `DataType` we produce, proving the quirk is handled
//! rather than crashing or being silently dropped.
#![cfg(test)]

use std::str::FromStr;

use models::datatypes::DataType;
use models::versioned_mediatype::{MediaType, Version, VersionDate};
use openapiv3_resolve::{ResolvedOpenAPI, openapiv3::OpenAPI};

use crate::Operation;

/// Converts the 200 JSON body of `GET /test` in `spec` into a [`DataType`].
fn datatype_of(spec: &str) -> DataType {
    let openapi: OpenAPI = serde_yaml::from_str(spec).expect("valid openapi yaml");
    let resolved = ResolvedOpenAPI::try_from(&openapi).expect("resolves");
    let item = resolved.paths().paths.get("/test").expect("path /test");
    let (_verb, operation) = item
        .iter()
        .find(|(verb, _)| *verb == "get")
        .expect("GET /test");
    let operation = Operation::from_path_verb_and_resolved_operation("/test", "get", operation)
        .expect("operation parses");
    let version = operation
        .versions
        .get(&Version::Stable(
            VersionDate::from_str("2023-01-01").unwrap(),
        ))
        .expect("2023-01-01 version");
    version
        .response_bodies
        .get(&MediaType::Json)
        .expect("json response body")
        .clone()
}

#[test]
fn quirk_1_1_discriminator_without_one_of() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_1_discriminator_without_one_of.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_2_one_of_without_discriminator() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_2_one_of_without_discriminator.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_3_type_object_next_to_one_of() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_3_type_object_next_to_one_of.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_4_unsatisfiable_union_type_object_over_string_branches() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_4_unsatisfiable_union_type_object_over_string_branches.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_5_children_redeclare_discriminator_property() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_5_children_redeclare_discriminator_property.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_6_parent_child_cycle() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_6_parent_child_cycle.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_7_mapping_targets_that_are_not_children() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_7_mapping_targets_that_are_not_children.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_8_mapping_keys_disagree_with_discriminator_enum() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_8_mapping_keys_disagree_with_discriminator_enum.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_9_any_of_used_as_a_tagged_union() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_9_any_of_used_as_a_tagged_union.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_1_10_wrapper_with_one_branch() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_1_10_wrapper_with_one_branch.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_2_2_single_value_enum() {
    let datatype = datatype_of(include_str!("../fixtures/quirk_2_2_single_value_enum.yaml"));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_2_3_unions_of_inline_enums() {
    let datatype = datatype_of(include_str!(
        "../fixtures/quirk_2_3_unions_of_inline_enums.yaml"
    ));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_5_3_nullable() {
    let datatype = datatype_of(include_str!("../fixtures/quirk_5_3_nullable.yaml"));
    insta::assert_json_snapshot!(datatype);
}

#[test]
fn quirk_5_4_free_form_object() {
    let datatype = datatype_of(include_str!("../fixtures/quirk_5_4_free_form_object.yaml"));
    insta::assert_json_snapshot!(datatype);
}
