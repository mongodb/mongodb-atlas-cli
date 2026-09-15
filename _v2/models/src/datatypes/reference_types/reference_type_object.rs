use std::collections::HashMap;

use serde::{Deserialize, Serialize};
use thiserror::Error;

use crate::datatypes::DataType;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ReferenceTypeObject {
    properties: HashMap<String, DataType>,
}

#[derive(Debug, Error)]
pub enum ReferenceTypeObjectTryNewError {
    #[error("properties must not be empty")]
    EmptyMap,
    #[error("properties must not contain empty-string keys")]
    EmptyMapKey,
}

impl ReferenceTypeObject {
    pub fn try_new(
        properties: HashMap<String, DataType>,
    ) -> Result<Self, ReferenceTypeObjectTryNewError> {
        if properties.is_empty() {
            return Err(ReferenceTypeObjectTryNewError::EmptyMap);
        }
        if properties.keys().any(String::is_empty) {
            return Err(ReferenceTypeObjectTryNewError::EmptyMapKey);
        }
        Ok(Self { properties })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::datatypes::value_type::ValueTypeBoolean;
    use crate::datatypes::ValueType;

    fn bool_prop(key: &str) -> (String, DataType) {
        (
            key.to_string(),
            DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})),
        )
    }

    #[test]
    fn rejects_empty_map() {
        let err = ReferenceTypeObject::try_new(HashMap::new()).unwrap_err();
        assert!(matches!(err, ReferenceTypeObjectTryNewError::EmptyMap));
    }

    #[test]
    fn rejects_empty_key() {
        let mut props = HashMap::new();
        let (valid_key, valid_value) = bool_prop("valid");
        props.insert(valid_key, valid_value);
        let (empty_key, empty_key_value) = bool_prop("");
        props.insert(empty_key, empty_key_value);
        let err = ReferenceTypeObject::try_new(props).unwrap_err();
        assert!(matches!(err, ReferenceTypeObjectTryNewError::EmptyMapKey));
    }

    #[test]
    fn accepts_nonempty_map() {
        let mut props = HashMap::new();
        let (key, value) = bool_prop("valid");
        props.insert(key, value);
        let obj = ReferenceTypeObject::try_new(props).unwrap();
        assert!(obj.properties.contains_key("valid"));
    }
}
