use anyhow::{Context, Result, bail};
use openapiv3::OpenAPI;
use std::collections::BTreeSet;
use tracing::{error, info};

use models::{hierarchy::Hierarchy, operation_id::OperationId};

pub fn cli_config(openapi: &OpenAPI) -> Result<Hierarchy> {
    let mut hierarchy = Hierarchy::new();
    clusters_config(&mut hierarchy, openapi).context("configuring the clusters config")?;

    Ok(hierarchy)
}

fn clusters_config(hierarchy: &mut Hierarchy, openapi: &OpenAPI) -> Result<()> {
    let clusters_str = "Clusters".to_string();
    let clusters_group = hierarchy.group(clusters_str.clone());
    let cluster = clusters_group.entity("Cluster".to_string());

    // Make sure we included all for the POC
    // - Collect all cluster operation ids
    // - When we use them, remove the operation id
    // - Log the ones we forgot
    let mut operation_ids = openapi
        .paths
        .iter()
        .flat_map(|(_path, item)| {
            item.as_item()
                .unwrap()
                .iter()
                .filter_map(|(_verb, operation)| {
                    if operation.deprecated {
                        return None;
                    }

                    operation.tags.contains(&clusters_str).then(|| {
                        Ok(operation
                            .operation_id
                            .as_ref()
                            .context("operation id missing")?
                            .parse::<OperationId>()?)
                    })
                })
        })
        .collect::<Result<BTreeSet<OperationId>>>()?;

    let mut take_operation_id = |value: &str| -> Result<OperationId> {
        let operation_id = value.parse::<OperationId>()?;
        if operation_ids.remove(&operation_id) {
            Ok(operation_id)
        } else {
            bail!("operation id is not removed yet: {operation_id}")
        }
    };

    let mut take_operation_id_option = |value| Some(take_operation_id(value)).transpose();

    cluster.create = take_operation_id_option("createGroupCluster")?;
    cluster.read = take_operation_id_option("getGroupCluster")?;
    cluster.update = take_operation_id_option("updateGroupCluster")?;
    cluster.delete = take_operation_id_option("deleteGroupCluster")?;
    cluster.list = take_operation_id_option("listClusterDetails")?;

    cluster.actions.insert(
        "status".to_string(),
        take_operation_id("getGroupClusterStatus")?,
    );

    cluster.actions.insert(
        "restartPrimaries".to_string(),
        take_operation_id("restartGroupClusterPrimaries")?,
    );

    cluster.actions.insert(
        "validateConfiguration".to_string(),
        take_operation_id("validateGroupClusterConfigurations")?,
    );

    cluster.actions.insert(
        "upgradeToFlexCluster".to_string(),
        take_operation_id("upgradeGroupClusterTenantUpgrade")?,
    );

    cluster.actions.insert(
        "listProviderRegions".to_string(),
        take_operation_id("listGroupClusterProviderRegions")?,
    );

    cluster.set_component_action(
        "MongoDBEmployeeAccess",
        "grant",
        take_operation_id("grantGroupClusterMongoDbEmployeeAccess")?,
    );

    cluster.set_component_action(
        "MongoDBEmployeeAccess",
        "revoke",
        take_operation_id("revokeGroupClusterMongoDbEmployeeAccess")?,
    );

    cluster.set_component_action(
        "FeatureCompatibilityVersion",
        "pin",
        take_operation_id("pinGroupClusterFeatureCompatibilityVersion")?,
    );

    cluster.set_component_action(
        "FeatureCompatibilityVersion",
        "unpin",
        take_operation_id("unpinGroupClusterFeatureCompatibilityVersion")?,
    );

    cluster.set_component_action(
        "AdvancedConfigurationOptions",
        "get",
        take_operation_id("getGroupClusterProcessArgs")?,
    );

    cluster.set_component_action(
        "AdvancedConfigurationOptions",
        "update",
        take_operation_id("updateGroupClusterProcessArgs")?,
    );

    cluster.set_component_action(
        "SqlInterface",
        "read",
        take_operation_id("getGroupClusterSqlInterface")?,
    );

    cluster.set_component_action(
        "SqlInterface",
        "update",
        take_operation_id("updateGroupClusterSqlInterface")?,
    );

    cluster.set_component_action(
        "SampleDataset",
        "progress",
        take_operation_id("getGroupSampleDatasetLoad")?,
    );

    cluster.set_component_action(
        "SampleDataset",
        "request",
        take_operation_id("requestGroupSampleDatasetLoad")?,
    );

    error!(?operation_ids, "not all operation ids are assigned");

    Ok(())
}
