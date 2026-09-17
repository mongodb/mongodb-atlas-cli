#[derive(Debug, clap::Parser)]
pub struct Cli {
    #[clap(subcommand)]
    sub_command: CliSubCommands,
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
    /// Create a new cluster (version tree, testing, leave out in real cli)
    Create(CreateGroupClusterProbe),
    /// Create a update cluster (version tree, testing, leave out in real cli)
    Update,
}

#[derive(Debug, clap::Args)]
pub struct CreateGroupClusterProbe {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,
}

#[derive(Clone, Debug, clap::ValueEnum)]
pub enum CreateGroupClusterVersion {
    V20260901,
    V20250101,
}

#[derive(Debug, clap::Args)]
pub struct CreateGroupClusterV20260901 {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,
}

#[derive(Debug, clap::Args)]
pub struct CreateGroupClusterV20250101 {
    #[arg(long, value_enum, default_value_t = CreateGroupClusterVersion::V20260901)]
    version: CreateGroupClusterVersion,
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

#[derive(Debug, clap::Args)]
pub struct CreateBackupV20260802 {
    #[arg(long, value_enum, default_value_t = CreateBackupVersionVersion::V20260902)]
    version: CreateBackupVersionVersion,
}
