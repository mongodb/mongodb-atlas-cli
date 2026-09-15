use serde::{Deserialize, Serialize};

use crate::datatypes::{reference_types::ReferenceType, value_type::ValueType};

pub mod reference_types;
pub mod value_type;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub enum DataType {
    ReferenceType(ReferenceType),
    ValueType(ValueType),
}
