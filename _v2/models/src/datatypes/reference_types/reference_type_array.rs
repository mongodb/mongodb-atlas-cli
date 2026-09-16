use serde::{Deserialize, Serialize};

use crate::datatypes::DataType;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ReferenceTypeArray {
    entries_type: Box<DataType>,
}

impl ReferenceTypeArray {
    /// Infallible: [`DataType`] can't be an empty string or hashmap itself.
    pub fn new(entries_type: DataType) -> Self {
        Self {
            entries_type: Box::new(entries_type),
        }
    }

    pub fn entries_type(&self) -> &DataType {
        &self.entries_type
    }
}
