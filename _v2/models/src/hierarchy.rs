use std::collections::BTreeMap;

use serde::{Deserialize, Serialize};

use crate::operation_id::OperationId;

/// The CLI hierarchy: group -> entities -> operations.
///
/// `Deserialize` lets a manual `cli.yaml` drive code generation instead of
/// inferring the structure from the spec (see `/atlas-cli-codegen`).
#[derive(Default, Debug, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Hierarchy {
    #[serde(default)]
    groups: BTreeMap<String, Group>,
}

impl Hierarchy {
    pub fn new() -> Self {
        Hierarchy {
            groups: Default::default(),
        }
    }

    pub fn group(&mut self, key: String) -> &mut Group {
        self.groups.entry(key).or_default()
    }

    pub fn groups(&self) -> impl Iterator<Item = (&String, &Group)> {
        self.groups.iter()
    }
}

#[derive(Default, Debug, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Group {
    #[serde(default)]
    pub description: Option<String>,
    #[serde(default)]
    entities: BTreeMap<String, Entity>,
}

impl Group {
    pub fn entity(&mut self, key: String) -> &mut Entity {
        self.entities.entry(key).or_default()
    }

    pub fn entities(&self) -> impl Iterator<Item = (&String, &Entity)> {
        self.entities.iter()
    }
}

#[derive(Default, Debug, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Entity {
    #[serde(default)]
    pub description: Option<String>,
    // crudl
    pub create: Option<OperationId>,
    pub read: Option<OperationId>,
    pub update: Option<OperationId>,
    pub delete: Option<OperationId>,
    pub list: Option<OperationId>,

    #[serde(default)]
    pub actions: BTreeMap<String, OperationId>,

    // components
    #[serde(default)]
    pub components: BTreeMap<String, Component>,
}

impl Entity {
    pub fn component(&mut self, key: String) -> &mut Component {
        self.components.entry(key).or_default()
    }

    pub fn set_component_action(&mut self, component: &str, action: &str, value: OperationId) {
        let component = self.component(component.to_string());

        match action {
            "read" => component.read = Some(value),
            "update" => component.update = Some(value),
            _ => {
                component.actions.insert(action.to_string(), value);
            }
        }
    }
}

#[derive(Default, Debug, Serialize, Deserialize)]
#[serde(deny_unknown_fields)]
pub struct Component {
    #[serde(default)]
    pub description: Option<String>,
    pub read: Option<OperationId>,
    pub update: Option<OperationId>,

    #[serde(default)]
    pub actions: BTreeMap<String, OperationId>,
}
