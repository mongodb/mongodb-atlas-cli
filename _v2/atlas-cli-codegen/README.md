# atlas-cli-codegen

Intermediate layer between the OpenAPI spec, a *manually authored CLI
hierarchy*, and the `atlas-cli-derive` proc macro.

The macro stays dumb. Its whole job is: read `spec.yaml` and `cli.yaml`, hand
both to this crate, and append the returned token stream to the derived item.
All logic worth testing lives here, in a normal library.

## Pipeline

```
spec.yaml  ─┐
             ├→ ResolvedOpenAPI + models::Hierarchy → GeneratedCli (this crate) → TokenStream
cli.yaml   ─┘
```

`cli.yaml` mirrors [`models::hierarchy::Hierarchy`](/models/src/hierarchy.rs):
`group → entity → create/read/update/delete/list + actions + components`. The
hierarchy is ours to author, so operations land under the right commands
instead of under whatever a spec-inference heuristic would guess. Operations
never listed in `cli.yaml` are simply not generated.

Every group, entity, and component takes an optional `description`; without
it, codegen falls back to a `Manage your Atlas CLI <name>` template.

```yaml
# cli.yaml
groups:
  Clusters:                 # single entity → flattened group command
    description: Manage your clusters.
    entities:
      Cluster:
        description: Manage cluster lifecycle.
        create: createGroupCluster
        read: getGroupCluster
        update: updateGroupCluster
        delete: deleteGroupCluster
        list: listClusterDetails
        actions:
          status: getGroupClusterStatus
        components:
          AdvancedConfigurationOptions:   # nested command: clusters advanced-configuration-options …
            description: Read and update advanced configuration.
            read: getGroupClusterProcessArgs
            update: updateGroupClusterProcessArgs
  Cloud Backups:            # multiple entities → nested entity commands
    entities:
      CompliancePolicy:
        description: Manage your backup compliance policy.
        read: getGroupBackupCompliancePolicy
        update: updateGroupBackupCompliancePolicy
      ExportBuckets:
        create: createGroupBackupExportBucket
        list: listGroupBackupExportBuckets
```

## Derive surface

```rust
#[derive(atlas_cli, clap::Parser)]
#[atlas_cli(
    spec = "path/to/spec.yaml",   // relative to the consuming crate's CARGO_MANIFEST_DIR
    hierarchy = "cli.yaml",
    excluded_operation_ids = ["..."]  // deliberate skips
)]
pub struct Cli {
    #[clap(subcommand)]
    sub_command: CliSubCommands,
}
```

The derive emits only additional items; the `struct Cli` itself stays put and
clap keeps deriving `Parser` on it.

### Coverage

Every spec `operationId` must be covered by exactly one of:

- a slot in `cli.yaml` (this is the point: it forces the hierarchy to be
  exhaustive and explicit), or
- `excluded_operation_ids` (the "we don't generate this one" catalog).

Neither → `compile_error!` listing the missing ids. This catalog is expected
to be large for a partial CLI; that is fine. A typo'd slot name is caught by
the hierarchy parser (`#[serde(deny_unknown_fields)]` rejects unknown keys, so
an action written at the wrong nesting level fails loudly instead of being
silently dropped).

### `cli.yaml` shape

Entity and component *actions* nest under an `actions:` key (they are part of
`models::hierarchy`'s schema, not free-form entity keys):

```yaml
      CompliancePolicy:
        read: getGroupBackupCompliancePolicy
        actions:
          disable: disableGroupBackupCompliancePolicy
      Cluster:
        components:
          MongoDBEmployeeAccess:
            actions:
              grant: grantGroupClusterMongoDbEmployeeAccess
              revoke: revokeGroupClusterMongoDbEmployeeAccess
```

`read`/`update` on a component are component-level slots and stay flat.

## Tree shapes

- **Single entity** — the group commands are flattened: the group variant
  holds one command struct whose subcommands are the entity's operations
  directly (`clusters create`, `clusters read`, ...).
- **Multiple entities** — the group variant holds one command struct whose
  subcommands are per-entity commands, and each entity command holds the
  entity's operations (`cloud-backups compliance-policy read`,
  `cloud-backups export-buckets list`, ...). Entity types are scoped by group
  (`{Group}{Entity}Command`) so entities named e.g. `Cluster` do not collide
  across groups.
- **Components** — a third level under an entity: the entity command has a
  component subcommand per component, each with its own operation enum
  (`clusters advanced-configuration-options read`). Component types are
  scoped by group + entity (`{Group}{Entity}{Component}Command`).

## Generated items (per operation)

For every operation referenced by the hierarchy (and not excluded):

- 1 probe struct (`CreateGroupClusterProbe`): `version` + `rest: Vec<String>`
  captured verbatim, `disable_help_flag`.
- 1 version enum (`CreateGroupClusterVersion`, `#[derive(ValueEnum)]`).
- 1 `clap::Parser` struct per API version (`CreateGroupClusterV20241023`),
  fields generated from the operation's path/query/header parameters.
- `impl Probe::execute`: match on the selected version, re-parse `rest`
  against the version struct, delegate.
- `impl VersionStruct::execute`: `todo!()` for now.

Variant names come from the slot keys verbatim: `create` → `Create`,
`read` → `Read`, an action key `restartPrimaries` → `RestartPrimaries`, a
component action `AdvancedConfigurationOptions` + `read` →
`AdvancedConfigurationOptionsRead`.

POC scope/deferred:

- Flags are plain `String` / `Option<String>`; typed flags (bool/int/enum)
  are a later pass.
- `x-xgen-operation-id-override` is ignored; naming uses the spec
  `operationId`.
