use std::process::ExitCode;

use atlas_cli_derive::atlas_cli;
use clap::Parser;

#[derive(atlas_cli, clap::Parser)]
#[atlas_cli(
    spec = "spec.yaml",
    hierarchy = "hierarchy.yaml",
    excluded_operation_ids = [
        "createGroup",
    ],
)]
pub struct Cli {
    #[clap(subcommand)]
    sub_command: CliSubCommands,
}

#[tokio::main]
async fn main() -> Result<(), ExitCode> {
    tracing_subscriber::fmt::init();

    let command = Cli::parse();

    command.execute().await
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn clusters_list_generated_operation_is_a_get_on_the_clusters_url() {
        use mongodb_atlas_cli::atlas::Operation;

        let op = ListClusterDetailsV20230101::parse_from(["clusters list", "--pretty", "true"]);

        assert_eq!(op.method(), http::Method::GET);
        assert_eq!(op.url(), "/api/atlas/v2/clusters?pretty=true");
        assert!(op.version().to_string().contains("2023-01-01"));
    }

    #[test]
    fn path_parameters_lose_their_curly_braces() {
        use mongodb_atlas_cli::atlas::Operation;

        let op = GetGroupBackupCompliancePolicyV20230101::parse_from([
            "compliance-policy read",
            "--group-id",
            "65d609455c11505db4a12c76",
        ]);
        assert_eq!(
            op.url(),
            "/api/atlas/v2/groups/65d609455c11505db4a12c76/backupCompliancePolicy"
        );
    }

    #[test]
    fn cli_parses_clusters_create_with_version_and_body_flags() {
        let cli = Cli::parse_from([
            "cli",
            "clusters",
            "create",
            "--version",
            "v20241023",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--name",
            "myCluster",
            "--advanced-configuration-minimum-enabled-tls-protocol",
            "TLS1_2",
        ]);

        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::Create(probe) = clusters.sub_command else {
            panic!("expected create subcommand");
        };
        assert_eq!(
            probe.rest,
            vec![
                "--group-id",
                "32b6e34b3d91647abb20e7b8",
                "--name",
                "myCluster",
                "--advanced-configuration-minimum-enabled-tls-protocol",
                "TLS1_2",
            ]
        );

        // Body flags are per-version: re-parse against the selected version's
        // struct and confirm the request-body flags landed.
        let fwd = CreateGroupClusterV20241023::parse_from(
            std::iter::once("clusters create".to_owned()).chain(probe.rest),
        );
        assert_eq!(fwd.groupId, "32b6e34b3d91647abb20e7b8");
        assert_eq!(fwd.name.as_deref(), Some("myCluster"));
        assert_eq!(
            fwd.advancedConfiguration_minimumEnabledTlsProtocol.as_deref(),
            Some("TLS1_2")
        );
    }

    #[test]
    fn cli_parses_nested_cloud_backups_entity() {
        // Cloud Backups is a multi-entity group: the entity name is a
        // subcommand under the group.
        let cli = Cli::parse_from([
            "cli",
            "cloud-backups",
            "compliance-policy",
            "read",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
        ]);

        let CliSubCommands::CloudBackups(backups) = cli.sub_command else {
            panic!("expected cloud backups subcommand");
        };
        let CloudBackupsSubcommands::CompliancePolicy(entity) = backups.sub_command else {
            panic!("expected compliance policy subcommand");
        };
        let CloudBackupsCompliancePolicySubcommands::Read(probe) = entity.sub_command else {
            panic!("expected read subcommand");
        };
        assert_eq!(probe.rest, vec!["--group-id", "32b6e34b3d91647abb20e7b8"]);
    }

    #[test]
    fn cli_parses_flattened_component_command() {
        // Clusters is a single-entity group; its components are nested
        // subcommands under the flattened group command.
        let cli = Cli::parse_from([
            "cli",
            "clusters",
            "advanced-configuration-options",
            "read",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--cluster-name",
            "myCluster",
        ]);

        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::AdvancedConfigurationOptions(component) = clusters.sub_command
        else {
            panic!("expected advanced configuration options subcommand");
        };
        let ClustersClusterAdvancedConfigurationOptionsSubcommands::Read(probe) =
            component.sub_command
        else {
            panic!("expected read subcommand");
        };
        assert_eq!(
            probe.rest,
            vec!["--group-id", "32b6e34b3d91647abb20e7b8", "--cluster-name", "myCluster"]
        );
    }

    use std::io::Write;

    /// Write a request-body fixture to the temp dir so clio's parser accepts
    /// the path at parse time.
    fn temp_json(label: &str, body: &str) -> std::path::PathBuf {
        let path = std::env::temp_dir()
            .join(format!("atlas-cli-golden-{label}-{}.json", std::process::id()));
        let mut file = std::fs::File::create(&path).unwrap();
        file.write_all(body.as_bytes()).unwrap();
        path
    }

    #[test]
    fn prepare_body_builds_json_from_body_flags() {
        let mut cmd = UpdateGroupBackupCompliancePolicyV20231001::parse_from([
            "compliance-policy update",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--authorized-email",
            "grace@example.com",
            "--authorized-user-first-name",
            "Grace",
            "--authorized-user-last-name",
            "Hopper",
            "--copy-protection-enabled",
            "--restore-window-days",
            "7",
        ]);
        let bytes = cmd.prepare_body().unwrap().unwrap();
        let body: serde_json::Value = serde_json::from_slice(&bytes).unwrap();
        assert_eq!(body["authorizedEmail"], "grace@example.com");
        // Real JSON scalars come off the typed flags, not strings.
        assert_eq!(body["copyProtectionEnabled"], true);
        assert_eq!(body["restoreWindowDays"], 7);
    }

    #[test]
    fn help_file_parses_without_required_args() {
        // `--help-file` short-circuits at the probe, before the version
        // struct re-parse demands `--group-id`.
        let cli = Cli::parse_from(["cli", "clusters", "create", "--help-file"]);
        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::Create(probe) = clusters.sub_command else {
            panic!("expected create subcommand");
        };
        assert!(probe.help_file);
        assert!(probe.rest.is_empty());
    }

    #[test]
    fn help_file_prints_the_request_body_schema() {
        // The schema the probe prints for the selected version is the version
        // struct's BODY_SCHEMA: valid object JSON that --file input must match.
        let schema: serde_json::Value =
            serde_json::from_str(CreateGroupClusterV20241023::BODY_SCHEMA).unwrap();
        assert_eq!(schema["type"], "object");
        assert!(schema["properties"]["name"].is_object());
    }

    #[test]
    fn bodyless_commands_have_no_help_file_flag() {
        // list carries no request body: its probe defines no --help-file, so
        // the probe's allow_hyphen_values rest capture takes the token (and
        // the version re-parse later rejects it as unknown).
        let cli = Cli::parse_from(["cli", "clusters", "list", "--help-file"]);
        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::List(probe) = clusters.sub_command else {
            panic!("expected list subcommand");
        };
        assert_eq!(probe.rest, vec!["--help-file".to_owned()]);
    }

    #[test]
    fn prepare_body_rejects_missing_required_fields() {
        let mut cmd = UpdateGroupBackupCompliancePolicyV20231001::parse_from([
            "compliance-policy update",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--authorized-email",
            "grace@example.com",
            "--copy-protection-enabled",
        ]);
        let error = cmd.prepare_body().unwrap_err();
        assert!(error.contains("required"), "unexpected: {error}");
    }

    #[test]
    fn prepare_body_rejects_file_combined_with_flags() {
        let path = temp_json(
            "xor",
            r#"{"authorizedEmail":"a@b.c","authorizedUserFirstName":"Ada","authorizedUserLastName":"Lovelace"}"#,
        );
        let mut cmd = UpdateGroupBackupCompliancePolicyV20231001::parse_from([
            "compliance-policy update",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--file",
            path.to_str().unwrap(),
            "--authorized-email",
            "grace@example.com",
        ]);
        let error = cmd.prepare_body().unwrap_err();
        std::fs::remove_file(path).ok();
        assert!(error.contains("not both"), "unexpected: {error}");
    }

    #[test]
    fn prepare_body_reads_and_validates_file() {
        let path = temp_json(
            "read",
            r#"{"authorizedEmail":"a@b.c","authorizedUserFirstName":"Ada","authorizedUserLastName":"Lovelace"}"#,
        );
        let mut cmd = UpdateGroupBackupCompliancePolicyV20231001::parse_from([
            "compliance-policy update",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--file",
            path.to_str().unwrap(),
        ]);
        let body = cmd.prepare_body().unwrap().unwrap();
        std::fs::remove_file(path).ok();
        assert!(body.starts_with(br#"{"authorizedEmail""#));
    }

    #[test]
    fn prepare_body_rejects_schema_invalid_file() {
        let path = temp_json("invalid", r#"{"cloudProvider":"KUBERNETES"}"#);
        let mut cmd = CreateGroupBackupExportBucketV20230101::parse_from([
            "export-buckets create",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
            "--file",
            path.to_str().unwrap(),
        ]);
        let error = cmd.prepare_body().unwrap_err();
        std::fs::remove_file(path).ok();
        assert!(error.contains("invalid"), "unexpected: {error}");
    }

    #[test]
    fn prepare_body_returns_empty_byte_for_bodyless_version() {
        use mongodb_atlas_cli::atlas::Operation;
        let cmd = ListClusterDetailsV20230101::parse_from(["clusters list"]);
        // No request body: no --file, nothing to prepare, nothing to send.
        assert_eq!(cmd.request_body(), bytes::Bytes::new());
    }

    /// An unsupported-shape body (oneOf export bucket request) has no flat
    /// flags: --file is the only way in, and omitting it is an error.
    #[test]
    fn complex_body_requires_file() {
        let mut cmd = CreateGroupBackupExportBucketV20240530::parse_from([
            "export-buckets create",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
        ]);
        let error = cmd.prepare_body().unwrap_err();
        assert!(error.contains("request body is required"), "unexpected: {error}");
    }
}
