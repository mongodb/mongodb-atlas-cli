use std::collections::BTreeMap;

use serde::{Deserialize, Serialize};
use thiserror::Error;

use crate::datatypes::reference_types::ReferenceType;

#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub struct ReferenceTypeOneOf {
    discriminator_field: String,
    options: BTreeMap<String, ReferenceType>,
}

#[derive(Debug, Error)]
pub enum ReferenceTypeOneOfTryNewError {
    #[error("discriminator field must not be empty")]
    EmptyString,
    #[error("options must not be empty")]
    EmptyMap,
    #[error("options must not contain empty-string keys")]
    EmptyMapKey,
}

impl ReferenceTypeOneOf {
    pub fn discriminator_field(&self) -> &str {
        &self.discriminator_field
    }

    pub fn options(&self) -> &BTreeMap<String, ReferenceType> {
        &self.options
    }

    pub fn try_new(
        discriminator_field: String,
        options: BTreeMap<String, ReferenceType>,
    ) -> Result<Self, ReferenceTypeOneOfTryNewError> {
        if discriminator_field.is_empty() {
            return Err(ReferenceTypeOneOfTryNewError::EmptyString);
        }
        if options.is_empty() {
            return Err(ReferenceTypeOneOfTryNewError::EmptyMap);
        }
        if options.keys().any(String::is_empty) {
            return Err(ReferenceTypeOneOfTryNewError::EmptyMapKey);
        }
        Ok(Self {
            discriminator_field,
            options,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::datatypes::value_type::ValueTypeBoolean;
    use crate::datatypes::{DataType, ValueType};

    fn bool_array_reference() -> ReferenceType {
        ReferenceType::new_array(DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})))
    }

    #[test]
    fn rejects_empty_discriminator() {
        let err = ReferenceTypeOneOf::try_new(String::new(), BTreeMap::new()).unwrap_err();
        assert!(matches!(err, ReferenceTypeOneOfTryNewError::EmptyString));
    }

    #[test]
    fn rejects_empty_options() {
        let err = ReferenceTypeOneOf::try_new("type".to_string(), BTreeMap::new()).unwrap_err();
        assert!(matches!(err, ReferenceTypeOneOfTryNewError::EmptyMap));
    }

    #[test]
    fn rejects_empty_option_key() {
        let mut options = BTreeMap::new();
        options.insert(String::new(), bool_array_reference());
        let err = ReferenceTypeOneOf::try_new("type".to_string(), options).unwrap_err();
        assert!(matches!(err, ReferenceTypeOneOfTryNewError::EmptyMapKey));
    }

    #[test]
    fn accepts_valid_input() {
        let mut options = BTreeMap::new();
        options.insert("default".to_string(), bool_array_reference());
        let one_of = ReferenceTypeOneOf::try_new("type".to_string(), options).unwrap();
        assert_eq!(one_of.discriminator_field, "type");
    }
}
