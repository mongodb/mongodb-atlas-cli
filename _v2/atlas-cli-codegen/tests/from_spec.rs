use std::collections::BTreeSet;

use atlas_cli_codegen::{
    CodegenError, CodegenOptions, GeneratedCli, GeneratedEntity, GeneratedFlag, GeneratedGroup,
    GeneratedOperation, from_config,
};

fn entity_operation_names(entity: &GeneratedEntity) -> Vec<String> {
    entity.operations.iter().map(|o| o.variant_name.clone()).collect()
}

const SPEC: &str = include_str!("fixtures/spec.yaml");
const HIERARCHY: &str = include_str!("fixtures/cli-hierarchy.yaml");

fn cli(options: CodegenOptions) -> GeneratedCli {
    from_config(SPEC, HIERARCHY, &options).unwrap()
}

fn default_options() -> CodegenOptions {
    CodegenOptions {
        excluded_operation_ids: BTreeSet::new(),
    }
}

fn group<'a>(cli: &'a GeneratedCli, ident: &str) -> &'a GeneratedGroup {
    cli.groups.iter().find(|g| g.ident == ident).unwrap()
}

fn entity<'a>(group: &'a GeneratedGroup, ident: &str) -> &'a GeneratedEntity {
    group.entities.iter().find(|e| e.ident == ident).unwrap()
}

fn operation<'a>(entity: &'a GeneratedEntity, variant: &str) -> &'a GeneratedOperation {
    entity
        .operations
        .iter()
        .find(|o| o.variant_name == variant)
        .unwrap()
}

#[test]
fn groups_are_built_from_hierarchy_sorted() {
    let cli = cli(default_options());
    let idents = cli.groups.iter().map(|g| g.ident.as_str()).collect::<Vec<_>>();
    assert_eq!(idents, ["CloudBackups", "Clusters"]);
}

#[test]
fn single_entity_group_is_flattened() {
    let cli = cli(default_options());
    let clusters = group(&cli, "Clusters");

    // "Clusters" has one entity (Cluster), so there is nothing nested at the
    // group level; the entity carries the operations.
    assert_eq!(clusters.entities.len(), 1);
    let cluster = entity(clusters, "Cluster");
    assert_eq!(cluster.operations.len(), 2);
    assert_eq!(
        cluster.operations.iter().map(|o| o.variant_name.as_str()).collect::<Vec<_>>(),
        ["Create", "Update"]
    );

    let create = operation(cluster, "Create");
    assert_eq!(create.probe_ident, "CreateGroupClusterProbe");
    assert_eq!(create.version_enum_ident, "CreateGroupClusterVersion");
}

#[test]
fn multi_entity_group_nests_per_entity_commands() {
    let cli = cli(default_options());
    let backups = group(&cli, "CloudBackups");

    assert_eq!(backups.entities.len(), 2);
    assert_eq!(
        backups.entities.iter().map(|e| e.ident.as_str()).collect::<Vec<_>>(),
        ["Cluster", "CompliancePolicy"]
    );

    let compliance = entity(backups, "CompliancePolicy");
    assert_eq!(entity_operation_names(compliance), ["Read"]);
    assert_eq!(compliance.operations[0].probe_ident, "GetGroupBackupCompliancePolicyProbe");

    let cluster = entity(backups, "Cluster");
    assert!(cluster.operations.is_empty());
    assert_eq!(cluster.components.len(), 1);
    let schedule = &cluster.components[0];
    assert_eq!(schedule.id, "Schedule");
    assert_eq!(
        schedule.operations.iter().map(|o| o.variant_name.as_str()).collect::<Vec<_>>(),
        ["Read"]
    );
    assert_eq!(schedule.operations[0].probe_ident, "StartGroupClusterBackupProbe");
}

#[test]
fn configured_descriptions_are_carried_on_every_level() {
    let cli = cli(default_options());
    let clusters = group(&cli, "Clusters");
    assert_eq!(clusters.description.as_deref(), Some("Cluster operations, grouped."));
    assert_eq!(entity(clusters, "Cluster").description.as_deref(), Some("Manage your clusters."));

    let backups = group(&cli, "CloudBackups");
    let compliance = entity(backups, "CompliancePolicy");
    assert_eq!(compliance.description.as_deref(), Some("Backup compliance policy."));
    assert_eq!(
        entity(backups, "Cluster").components[0].description.as_deref(),
        Some("Backup schedules.")
    );
}

#[test]
fn operation_versions_oldest_first_with_struct_names() {
    let cli = cli(default_options());
    let create = operation(entity(group(&cli, "Clusters"), "Cluster"), "Create");

    let variants = create
        .versions
        .iter()
        .map(|v| v.variant_ident.as_str())
        .collect::<Vec<_>>();
    assert_eq!(variants, ["V20230101", "V20241023"]);

    let structs = create
        .versions
        .iter()
        .map(|v| v.struct_ident.as_str())
        .collect::<Vec<_>>();
    assert_eq!(structs, ["CreateGroupClusterV20230101", "CreateGroupClusterV20241023"]);
}

#[test]
fn flags_from_parameters_with_optionality() {
    let cli = cli(default_options());
    let create = operation(entity(group(&cli, "Clusters"), "Cluster"), "Create");

    assert_eq!(
        create.flags,
        vec![
            GeneratedFlag {
                ident: "groupId".to_owned(),
                description: "Unique 24-hexadecimal digit string that identifies your project.".to_owned(),
                required: true,
                list: false,
            },
            GeneratedFlag {
                ident: "pretty".to_owned(),
                description: "Flag that indicates whether the response body should be pretty printed.".to_owned(),
                required: false,
                list: false,
            },
        ]
    );

    let update = operation(entity(group(&cli, "Clusters"), "Cluster"), "Update");
    assert_eq!(update.flags.len(), 1);
    assert_eq!(update.flags[0].ident, "groupId");
    assert!(update.flags[0].required);
}

#[test]
fn body_flags_derived_from_request_body_flattened() {
    let cli = cli(default_options());
    let create = operation(entity(group(&cli, "Clusters"), "Cluster"), "Create");

    assert_eq!(
        create.versions[0].body_flags,
        vec![
            GeneratedFlag {
                ident: "name".to_owned(),
                description: String::new(),
                required: true,
                list: false,
            },
            GeneratedFlag {
                ident: "person_first_name".to_owned(),
                description: String::new(),
                required: false,
                list: false,
            },
            GeneratedFlag {
                ident: "person_last_name".to_owned(),
                description: String::new(),
                required: false,
                list: false,
            },
            GeneratedFlag {
                ident: "providers".to_owned(),
                description: String::new(),
                required: false,
                list: true,
            },
            GeneratedFlag {
                ident: "region".to_owned(),
                description: String::new(),
                required: false,
                list: false,
            },
        ]
    );
}

#[test]
fn description_is_the_full_operation_description() {
    let cli = cli(default_options());
    let create = operation(entity(group(&cli, "Clusters"), "Cluster"), "Create");
    assert_eq!(
        create.description,
        "Creates one cluster in the specified project.\n\n\
         Clusters contain a group of hosts that maintain the same data set."
    );
}

#[test]
fn excluded_operation_ids_are_skipped() {
    let options = CodegenOptions {
        excluded_operation_ids: BTreeSet::from(["updateGroupCluster".to_owned()]),
    };
    let cli = cli(options);
    let cluster = entity(group(&cli, "Clusters"), "Cluster");
    assert_eq!(cluster.operations.len(), 1);
    assert_eq!(cluster.operations[0].variant_name, "Create");
}

#[test]
fn operation_missing_from_spec_is_an_error() {
    let hierarchy = r#"
groups:
  Clusters:
    entities:
      Cluster:
        create: doesNotExist
"#;
    let err = from_config(SPEC, hierarchy, &default_options()).unwrap_err();
    assert!(matches!(err, CodegenError::OperationNotFoundInSpec { op_id } if op_id == "doesNotExist"));
}

#[test]
fn uncovered_operations_error_by_default() {
    let spec = r#"
openapi: 3.0.3
info:
  title: coverage
  version: "1"
components:
  schemas:
    Body:
      type: object
      properties:
        name: { type: string }
paths:
  /one:
    post:
      description: First op.
      operationId: createOne
      tags: [Thing]
      responses:
        "200":
          description: OK
          content:
            application/vnd.atlas.2024-01-01+json:
              schema: { $ref: '#/components/schemas/Body' }
  /two:
    post:
      description: Second op.
      operationId: createTwo
      tags: [Thing]
      responses:
        "200":
          description: OK
          content:
            application/vnd.atlas.2024-01-01+json:
              schema: { $ref: '#/components/schemas/Body' }
"#;
    // Only `createOne` is in the hierarchy; `createTwo` is neither referenced
    // nor excluded.
    let hierarchy = r#"
groups:
  Thing:
    entities:
      One:
        create: createOne
"#;
    let err = from_config(spec, hierarchy, &default_options()).unwrap_err();
    assert!(matches!(
        err,
        CodegenError::UncoveredOperations { operation_ids } if operation_ids == ["createTwo"]
    ));
}

#[test]
fn excluded_operation_ids_mute_the_error() {
    let spec = r#"
openapi: 3.0.3
info:
  title: coverage
  version: "1"
components:
  schemas:
    Body:
      type: object
      properties:
        name: { type: string }
paths:
  /one:
    post:
      description: First op.
      operationId: createOne
      tags: [Thing]
      responses:
        "200":
          description: OK
          content:
            application/vnd.atlas.2024-01-01+json:
              schema: { $ref: '#/components/schemas/Body' }
  /two:
    post:
      description: Second op.
      operationId: createTwo
      tags: [Thing]
      responses:
        "200":
          description: OK
          content:
            application/vnd.atlas.2024-01-01+json:
              schema: { $ref: '#/components/schemas/Body' }
"#;
    let hierarchy = r#"
groups:
  Thing:
    entities:
      One:
        create: createOne
"#;
    let options = CodegenOptions {
        excluded_operation_ids: BTreeSet::from(["createTwo".to_owned()]),
    };
    let cli = from_config(spec, hierarchy, &options).unwrap();
    assert_eq!(cli.groups.len(), 1);
}

#[test]
fn operations_without_versions_are_skipped() {
    let spec = r#"
openapi: 3.0.3
info:
  title: no-versions
  version: "1"
paths:
  /ready:
    get:
      description: Readiness probe.
      operationId: getReady
      tags:
        - Infra
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema:
                type: object
"#;
    let hierarchy = r#"
groups:
  Infra:
    entities:
      Ready:
        read: getReady
"#;
    let cli = from_config(spec, hierarchy, &default_options()).unwrap();
    assert_eq!(cli.groups.len(), 0);
}
