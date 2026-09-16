use std::collections::BTreeMap;

use crate::operation_id::OperationId;

#[derive(Debug)]
pub struct Hierarchy {
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

#[derive(Default, Debug)]
pub struct Group {
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

#[derive(Default, Debug)]
pub struct Entity {
    // crudl
    pub create: Option<OperationId>,
    pub read: Option<OperationId>,
    pub update: Option<OperationId>,
    pub delete: Option<OperationId>,
    pub list: Option<OperationId>,

    pub actions: BTreeMap<String, OperationId>,

    // components
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

#[derive(Default, Debug)]
pub struct Component {
    pub read: Option<OperationId>,
    pub update: Option<OperationId>,

    pub actions: BTreeMap<String, OperationId>,
}
