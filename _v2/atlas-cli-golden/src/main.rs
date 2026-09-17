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
    let command = Cli::parse();

    command.execute().await
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn cli_parses_clusters_create_with_version() {
        let cli = Cli::parse_from([
            "cli",
            "clusters",
            "create",
            "--version",
            "v20241023",
            "--group-id",
            "32b6e34b3d91647abb20e7b8",
        ]);

        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::Create(probe) = clusters.sub_command else {
            panic!("expected create subcommand");
        };
        assert_eq!(probe.rest, vec!["--group-id", "32b6e34b3d91647abb20e7b8"]);

        let fwd = CreateGroupClusterV20241023::parse_from(
            std::iter::once("clusters create".to_owned()).chain(probe.rest),
        );
        assert_eq!(fwd.groupId, "32b6e34b3d91647abb20e7b8");
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
}
