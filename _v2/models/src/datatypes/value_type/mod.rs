use serde::{Deserialize, Serialize};

mod string_value_type;
mod value_type_enum;

pub use string_value_type::{ValueTypeString, ValueTypeStringValidation};
pub use value_type_enum::ValueTypeEnum;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub enum ValueType {
    Boolean(ValueTypeBoolean),
    Enum(ValueTypeEnum),
    Double(ValueTypeDouble),
    Integer(ValueTypeInteger),
    String(ValueTypeString),
}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ValueTypeBoolean {}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ValueTypeDouble {}

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ValueTypeInteger {
    pub minimum: Option<i64>,
    pub maximum: Option<i64>,
}

impl ValueTypeInteger {
    pub const fn new() -> Self {
        Self {
            minimum: None,
            maximum: None,
        }
    }

    pub const fn with_bounds(minimum: Option<i64>, maximum: Option<i64>) -> Self {
        Self { minimum, maximum }
    }
}

impl Default for ValueTypeInteger {
    fn default() -> Self {
        Self::new()
    }
}
