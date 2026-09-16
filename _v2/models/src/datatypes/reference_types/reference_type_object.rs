use std::collections::BTreeMap;

use serde::{Deserialize, Serialize};
use thiserror::Error;

use crate::datatypes::DataType;

/// A JSON object with known properties. An empty property map represents a
/// free-form or unmodelled object (an OpenAPI `type: object` with nothing else).
#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ReferenceTypeObject {
    properties: BTreeMap<String, DataType>,
}

#[derive(Debug, Error)]
pub enum ReferenceTypeObjectTryNewError {
    #[error("properties must not contain empty-string keys")]
    EmptyMapKey,
}

impl ReferenceTypeObject {
    pub fn properties(&self) -> &BTreeMap<String, DataType> {
        &self.properties
    }

    pub fn try_new(
        properties: BTreeMap<String, DataType>,
    ) -> Result<Self, ReferenceTypeObjectTryNewError> {
        if properties.keys().any(String::is_empty) {
            return Err(ReferenceTypeObjectTryNewError::EmptyMapKey);
        }
        Ok(Self { properties })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::datatypes::ValueType;
    use crate::datatypes::value_type::ValueTypeBoolean;

    fn bool_prop(key: &str) -> (String, DataType) {
        (
            key.to_string(),
            DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})),
        )
    }

    #[test]
    fn accepts_empty_map_as_free_form_object() {
        let obj = ReferenceTypeObject::try_new(BTreeMap::new()).unwrap();
        assert!(obj.properties.is_empty());
    }

    #[test]
    fn rejects_empty_key() {
        let mut props = BTreeMap::new();
        let (valid_key, valid_value) = bool_prop("valid");
        props.insert(valid_key, valid_value);
        let (empty_key, empty_key_value) = bool_prop("");
        props.insert(empty_key, empty_key_value);
        let err = ReferenceTypeObject::try_new(props).unwrap_err();
        assert!(matches!(err, ReferenceTypeObjectTryNewError::EmptyMapKey));
    }

    #[test]
    fn accepts_nonempty_map() {
        let mut props = BTreeMap::new();
        let (key, value) = bool_prop("valid");
        props.insert(key, value);
        let obj = ReferenceTypeObject::try_new(props).unwrap();
        assert!(obj.properties.contains_key("valid"));
    }
}
