mod reference_type_array;
mod reference_type_object;
mod reference_type_one_of;

pub use reference_type_array::ReferenceTypeArray;
pub use reference_type_object::{ReferenceTypeObject, ReferenceTypeObjectTryNewError};
pub use reference_type_one_of::{ReferenceTypeOneOf, ReferenceTypeOneOfTryNewError};

use std::collections::HashMap;

use serde::{Deserialize, Serialize};
use thiserror::Error;

use crate::datatypes::DataType;

/// A reference to another schema. The variant decides how the referenced
/// schema is combined with its siblings.
#[derive(Clone, Debug, PartialEq, Eq, Serialize, Deserialize)]
pub enum ReferenceType {
    Array(ReferenceTypeArray),
    Object(ReferenceTypeObject),
    OneOf(ReferenceTypeOneOf),
    // AnyOf = Object with Optional fields
    // - important with AnyOf -> our OpenApi spec does have schemas with enum types with 1 entry that should be merged through anyOf and should become the discriminator
    // AllOf = Merge all schemas into 1 object
}

#[derive(Debug, Error)]
pub enum ReferenceTypeTryNewError {
    #[error(transparent)]
    ReferenceTypeObjectTryNewError(#[from] ReferenceTypeObjectTryNewError),
    #[error(transparent)]
    ReferenceTypeOneOfTryNewError(#[from] ReferenceTypeOneOfTryNewError),
}

impl ReferenceType {
    /// Infallible: [`ReferenceTypeArray`] holds a [`DataType`], which can't be an empty string or map.
    pub fn new_array(entries_type: DataType) -> Self {
        Self::Array(ReferenceTypeArray::new(entries_type))
    }

    pub fn try_new_object(
        properties: HashMap<String, DataType>,
    ) -> Result<Self, ReferenceTypeTryNewError> {
        Ok(Self::Object(ReferenceTypeObject::try_new(properties)?))
    }

    pub fn try_new_one_of(
        discriminator_field: String,
        options: HashMap<String, ReferenceType>,
    ) -> Result<Self, ReferenceTypeTryNewError> {
        Ok(Self::OneOf(ReferenceTypeOneOf::try_new(
            discriminator_field,
            options,
        )?))
    }
}
