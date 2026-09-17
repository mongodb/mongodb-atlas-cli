Desired macro:
```
#[derive(atlas_cli, clap::Parser)]
#[atlas_cli(spec = "../spec.yaml", excluded_operation_ids = ["excluded_operation_id_1", "excluded_operation_id_2"])]
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

This will generate code similar to the code in the `src/` directory.
The spec yaml will be transformed using the `_v2/models/src/spec.rs`, then that spec will be used to generate the code.

The code below will expand to 1 probe struct + 1 struct for all the versions.
Also the impl block which generates the match logic and versioning logic.
The information is comming by the looked up operation id.
```
#[atlas_cli(operation_id = CreateGroupCluster)]
```

Fields should be generated from parameters on the operation.
Execute implementation can stay `todo!` for now.
