use serde::{Deserialize, Serialize};

use crate::datatypes::{reference_types::ReferenceType, value_type::ValueType};

pub mod json_schema;
pub mod reference_types;
pub mod value_type;

/// `ReferenceType` and `ValueType` carry disjoint `serde` tags, so an untagged
/// `DataType` nests one level shallower on the wire.
#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
#[serde(untagged)]
pub enum DataType {
    ReferenceType(ReferenceType),
    ValueType(ValueType),
}
