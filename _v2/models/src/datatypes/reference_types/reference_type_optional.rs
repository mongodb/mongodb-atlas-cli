use serde::{Deserialize, Serialize};

use crate::datatypes::DataType;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ReferenceTypeOptional {
    data_type: Box<DataType>,
}

impl ReferenceTypeOptional {
    /// Infallible: [`DataType`] can't be an empty string or hashmap itself.
    pub fn new(data_type: DataType) -> Self {
        Self {
            data_type: Box::new(data_type),
        }
    }

    pub fn data_type(&self) -> &DataType {
        &self.data_type
    }
}
