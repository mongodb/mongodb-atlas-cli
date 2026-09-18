use std::collections::BTreeSet;

use atlas_cli_codegen::{CodegenOptions, from_config, to_tokens};

const SPEC: &str = include_str!("fixtures/spec.yaml");
const HIERARCHY: &str = include_str!("fixtures/cli-hierarchy.yaml");

fn generated_source() -> String {
    let cli = from_config(
        SPEC,
        HIERARCHY,
        &CodegenOptions {
            excluded_operation_ids: BTreeSet::new(),
        },
    )
    .unwrap();
    let tokens = to_tokens(&cli);
    eprintln!("RAW: {}", tokens);
    let file: syn::File = syn::parse2(tokens).expect("generated code parses as Rust");
    prettyplease::unparse(&file)
}

#[test]
fn emitted_source_matches_golden() {
    insta::assert_snapshot!(generated_source());
}

#[test]
fn emitted_source_starts_with_cli_subcommands() {
    let source = generated_source();
    assert!(source.contains("pub enum CliSubCommands"));
    assert!(source.contains("pub enum ClustersSubcommands"));
    assert!(source.contains("pub enum CloudBackupsSubcommands"));
    assert!(source.contains("pub enum CloudBackupsCompliancePolicySubcommands"));
    assert!(source.contains("pub enum CloudBackupsClusterScheduleSubcommands"));
}

#[test]
fn emitted_source_generates_probe_and_version_structs() {
    let source = generated_source();
    assert!(source.contains("pub struct CreateGroupClusterProbe"));
    assert!(source.contains("pub struct CreateGroupClusterV20241023"));
    assert!(source.contains("pub enum CreateGroupClusterVersion"));
}

#[test]
fn emitted_source_wires_up_request_bodies() {
    let source = generated_source();
    // Versions with a request body take it from --file/stdin XOR the flat
    // flags, and report it through Operation::request_body.
    assert!(source.contains("file: Option<::clio::Input>"));
    assert!(source.contains("pub(crate) mod __atlas_cli_body"));
    assert!(source.contains("fn request_body(&self) -> ::bytes::Bytes"));
    // Body-flag leaves are typed by their schema scalar kind: required
    // scalars stay optional at the CLI (schema validation enforces presence),
    // arrays of scalars become optional Vec fields.
    assert!(source.contains("name: Option<String>"));
    assert!(source.contains("providers: Option<Vec<String>>"));
    assert!(source.contains("person_first_name: Option<String>"));
}

fn struct_source<'a>(source: &'a str, name: &str) -> &'a str {
    let start = source.find(name).unwrap_or_else(|| panic!("{name} in output")) + name.len();
    let body = &source[start..];
    let end = body.find("\n}").expect("struct body ends");
    &body[..end]
}

#[test]
fn emitted_source_gates_help_file_on_request_bodies() {
    let source = generated_source();
    // Operations with a request body short-circuit `--help-file` on the probe
    // (before required flags are validated) and print the request-body schema.
    assert!(struct_source(&source, "pub struct CreateGroupClusterProbe").contains("help_file: bool"));
    // Bodyless operations (GET) get no `--help-file` anywhere.
    assert!(!struct_source(&source, "pub struct GetGroupBackupCompliancePolicyProbe").contains("help_file"));
}

#[test]
fn emitted_source_keeps_bodyless_versions_cloneable() {
    let source = generated_source();
    // getGroupBackupCompliancePolicy is a GET with no request body: no
    // --file, no prepared-body state, still Clone.
    let compliance =
        "#[derive(Debug, Clone, ::clap::Parser)]\npub struct GetGroupBackupCompliancePolicyV20230101";
    assert!(source.contains(compliance));
    // startGroupClusterBackup is a POST with no request body; it must not
    // gain a --file flag either.
    let no_body =
        "#[derive(Debug, Clone, ::clap::Parser)]\npub struct StartGroupClusterBackupV20230101";
    assert!(source.contains(no_body));
}
