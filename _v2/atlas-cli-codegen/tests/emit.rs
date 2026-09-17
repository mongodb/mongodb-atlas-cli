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
