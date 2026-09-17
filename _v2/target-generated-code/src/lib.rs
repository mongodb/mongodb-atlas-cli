use std::process::ExitCode;

use clap::Parser;

#[derive(Debug, clap::Parser)]
pub struct Cli {
    #[clap(subcommand)]
    sub_command: CliSubCommands,
}

impl Cli {
    pub async fn execute(self) -> Result<(), ExitCode> {
        match self.sub_command {
            CliSubCommands::Clusters(c) => match c.sub_command {
                ClustersSubcommands::Create(probe) => probe.execute().await,
                ClustersSubcommands::Update => todo!(),
            },
            CliSubCommands::Backup(_b) => todo!(),
        }
    }
}

#[derive(Debug, clap::Subcommand)]
pub enum CliSubCommands {
    /// Manage your Atlas CLI clusters
    Clusters(ClustersCommand),
    /// Manage your Atlas CLI backups
    Backup(BackupCommand),
}

#[derive(Debug, clap::Args)]
pub struct ClustersCommand {
    #[clap(subcommand)]
    sub_command: ClustersSubcommands,
}

#[derive(Debug, clap::Subcommand)]
pub enum ClustersSubcommands {
    /// Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. This feature is not available for serverless clusters.
    Create(CreateGroupClusterProbe),
    /// Create a update cluster (version tree, testing, leave out in real cli)
    Update,
}

#[derive(Debug, clap::Args)]
#[command(disable_help_flag = true)]
pub struct CreateGroupClusterProbe {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,

    /// Raw args for the selected version's command, captured verbatim.
    #[arg(allow_hyphen_values = true)]
    rest: Vec<String>,
}

impl CreateGroupClusterProbe {
    /// Re-parse the captured raw args against the version-specific command.
    pub async fn execute(self) -> Result<(), ExitCode> {
        match self.version {
            CreateGroupClusterVersion::V20260901 => {
                let args = std::iter::once("clusters create".to_owned()).chain(self.rest);
                let cmd = CreateGroupClusterV20260901::parse_from(args);
                cmd.execute().await
            }
            CreateGroupClusterVersion::V20250101 => {
                let args = std::iter::once("clusters create".to_owned()).chain(self.rest);
                let cmd = CreateGroupClusterV20250101::parse_from(args);
                cmd.execute().await
            }
        }
    }
}

#[derive(Clone, Debug, clap::ValueEnum)]
pub enum CreateGroupClusterVersion {
    V20260901,
    V20250101,
}

/// Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. This feature is not available for serverless clusters.
///
// Please note that using an `instanceSize` of M2 or M5 will create a Flex cluster instead. Support for the `instanceSize` of M2 or M5 will be discontinued in January 2026. We recommend using the Create Flex Cluster API for such configurations moving forward.
///
/// Flags differ per API version (--version). This help describes version V20260901.
#[derive(Debug, clap::Parser)]
pub struct CreateGroupClusterV20260901 {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,
    #[arg(long)]
    blah: Option<String>,
    #[arg(long)]
    debug: bool,
}

impl CreateGroupClusterV20260901 {
    pub async fn execute(&self) -> Result<(), ExitCode> {
        todo!()
    }
}

/// Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. This feature is not available for serverless clusters.
///
// Please note that using an `instanceSize` of M2 or M5 will create a Flex cluster instead. Support for the `instanceSize` of M2 or M5 will be discontinued in January 2026. We recommend using the Create Flex Cluster API for such configurations moving forward.
///
/// Flags differ per API version (--version). This help describes version V20250101.
#[derive(Debug, clap::Parser)]
pub struct CreateGroupClusterV20250101 {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,
}

impl CreateGroupClusterV20250101 {
    pub async fn execute(&self) -> Result<(), ExitCode> {
        todo!()
    }
}

#[derive(Debug, clap::Args)]
pub struct BackupCommand {
    #[clap(subcommand)]
    sub_command: BackupSubcommands,
}

#[derive(Debug, clap::Subcommand)]
pub enum BackupSubcommands {
    Create(CreateBackupVersion),
    Update,
}

#[derive(Debug, clap::Args)]
pub struct CreateBackupVersion {
    #[arg(long, value_enum, default_value_t = CreateBackupVersionVersion::V20260902)]
    version: CreateBackupVersionVersion,
}

#[derive(Copy, Clone, Debug, PartialEq, Eq, PartialOrd, Ord, ::clap::ValueEnum)]
pub enum CreateBackupVersionVersion {
    V20260902,
    V20260802,
}

#[derive(Debug, clap::Args)]
pub struct CreateBackupV20260902 {
    #[arg(long, value_enum, default_value_t = CreateBackupVersionVersion::V20260902)]
    version: CreateBackupVersionVersion,
}

impl CreateBackupV20260902 {
    pub async fn execute(&self) -> Result<(), ExitCode> {
        todo!()
    }
}

#[derive(Debug, clap::Args)]
pub struct CreateBackupV20260802 {
    #[arg(long, value_enum, default_value_t = CreateBackupVersionVersion::V20260902)]
    version: CreateBackupVersionVersion,
}

impl CreateBackupV20260802 {
    pub async fn execute(&self) -> Result<(), ExitCode> {
        todo!()
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use clap::Parser;

    #[test]
    fn probe_forwards_unknown_args_to_versioned_command() {
        let cli = Cli::parse_from([
            "cli",
            "clusters",
            "create",
            "--version",
            "v20260901",
            "--blah",
            "1",
            "--debug",
        ]);
        let CliSubCommands::Clusters(clusters) = cli.sub_command else {
            panic!("expected clusters subcommand");
        };
        let ClustersSubcommands::Create(probe) = clusters.sub_command else {
            panic!("expected create subcommand");
        };
        assert_eq!(probe.rest, vec!["--blah", "1", "--debug"]);

        let fwd = CreateGroupClusterV20260901::parse_from(
            std::iter::once("clusters create".to_owned()).chain(probe.rest),
        );
        assert_eq!(fwd.blah.as_deref(), Some("1"));
        assert!(fwd.debug);
    }
}
