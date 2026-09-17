Desired macro:
```
#[derive(atlas_cli, clap::Parser)]
#[atlas_cli(spec = "../spec.yaml")]
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

#[derive(atlas_cli, Debug, clap::Subcommand)]
pub enum ClustersSubcommands {
    #[atlas_cli(operation_id = CreateGroupCluster)]
    Create,
    #[atlas_cli(operation_id = UpdateGroupCluster)]
    Update,
}
```
