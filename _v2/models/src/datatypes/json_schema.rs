//! Convert a [`DataType`] into a JSON Schema (`serde_json::json`).
//!
//! Every schema [`DataType::to_json_schema`] produces is meta-valid by
//! construction; [`is_valid`] and [`validate`] are provided as a safety net
//! for any schema, auto-detecting the draft like `jsonschema::meta` does.

use serde_json::{Map, Value, json};

use jsonschema::ValidationError;

use crate::datatypes::reference_types::{ReferenceTypeObject, ReferenceTypeOneOf};
use crate::datatypes::value_type::{
    ValueType, ValueTypeInteger, ValueTypeString, ValueTypeStringValidation,
};
use crate::datatypes::{DataType, ReferenceType};

impl DataType {
    /// Convert to a JSON Schema.
    pub fn to_json_schema(&self) -> Value {
        match self {
            DataType::ReferenceType(reference_type) => reference_type_schema(reference_type),
            DataType::ValueType(value_type) => value_type_schema(value_type),
        }
    }
}

/// Check that `schema` is valid against the JSON Schema meta-schema.
pub fn is_valid(schema: &Value) -> bool {
    jsonschema::meta::is_valid(schema)
}

/// Validate `schema` against the JSON Schema meta-schema.
pub fn validate(schema: &Value) -> Result<(), ValidationError<'_>> {
    jsonschema::meta::validate(schema)
}

fn value_type_schema(value_type: &ValueType) -> Value {
    match value_type {
        ValueType::Boolean(_) => json!({"type": "boolean"}),
        ValueType::Double(_) => json!({"type": "number"}),
        ValueType::Integer(integer) => integer_schema(integer),
        ValueType::String(string) => string_schema(string),
        ValueType::Enum(enumeration) => json!({"enum": enumeration.clone().into_inner()}),
    }
}

/// An integer schema carries its numeric bounds when the source spec declared
/// them (`minimum`/`maximum`), so a wrong value is rejected before the request
/// reaches the API.
fn integer_schema(integer: &ValueTypeInteger) -> Value {
    let mut keywords = Map::new();
    keywords.insert("type".to_string(), json!("integer"));
    if let Some(minimum) = integer.minimum {
        keywords.insert("minimum".to_string(), json!(minimum));
    }
    if let Some(maximum) = integer.maximum {
        keywords.insert("maximum".to_string(), json!(maximum));
    }
    Value::Object(keywords)
}

fn string_schema(string: &ValueTypeString) -> Value {
    let Some(validation) = string.validation() else {
        return json!({"type": "string"});
    };
    match validation {
        ValueTypeStringValidation::Regex { regex } => {
            json!({"type": "string", "pattern": regex})
        }
        ValueTypeStringValidation::Length { min, max } => {
            let mut keywords = Map::new();
            keywords.insert("type".to_string(), json!("string"));
            if let Some(min) = min {
                keywords.insert("minLength".to_string(), json!(min));
            }
            if let Some(max) = max {
                keywords.insert("maxLength".to_string(), json!(max));
            }
            Value::Object(keywords)
        }
    }
}

fn reference_type_schema(reference_type: &ReferenceType) -> Value {
    match reference_type {
        ReferenceType::Optional(optional) => nullable(optional.data_type().to_json_schema()),
        ReferenceType::Array(array) => json!({
            "type": "array",
            "items": array.entries_type().to_json_schema()
        }),
        ReferenceType::Object(object) => object_schema(object),
        ReferenceType::OneOf(one_of) => one_of_schema(one_of),
    }
}

/// Widen a schema's `type` to include `"null"` — the JSON Schema way to mark
/// a field nullable (`{"type": ["integer", "null"]}`, valid since draft-07).
/// Schemas with no `type` keyword (a `oneOf`) fall back to an `anyOf` with
/// `{"type": "null"}`.
fn nullable(schema: Value) -> Value {
    let Some(object) = schema.as_object() else {
        return json!({"anyOf": [schema, {"type": "null"}]});
    };
    let Some(Value::String(kind)) = object.get("type") else {
        return json!({"anyOf": [schema, {"type": "null"}]});
    };
    let mut object = object.clone();
    object.insert("type".to_string(), json!([kind, "null"]));
    Value::Object(object)
}

fn object_schema(object: &ReferenceTypeObject) -> Value {
    let properties = object.properties();
    if properties.is_empty() {
        // An empty property map is a free-form object.
        return json!({"type": "object"});
    }
    let mut schema = Map::new();
    schema.insert("type".to_string(), json!("object"));
    let properties: Map<String, Value> = properties
        .iter()
        .map(|(name, data_type)| (name.clone(), data_type.to_json_schema()))
        .collect();
    schema.insert("properties".to_string(), Value::Object(properties));

    // A property the parser wrapped in `Optional` may be absent; every other
    // property is required. The parser marks a required-but-nullable field
    // `Optional` too, which drops it from `required` (the model cannot tell
    // "absent" and "null" apart, so `required` errs toward permissive).
    let required = object
        .properties()
        .iter()
        .filter(|(_, data_type)| {
            !matches!(
                data_type,
                DataType::ReferenceType(ReferenceType::Optional(_))
            )
        })
        .map(|(name, _)| name.clone())
        .collect::<Vec<_>>();
    if !required.is_empty() {
        schema.insert("required".to_string(), json!(required));
    }

    Value::Object(schema)
}

fn one_of_schema(one_of: &ReferenceTypeOneOf) -> Value {
    let mut variants = Vec::with_capacity(one_of.options().len());
    let mut mapping = Map::new();
    for (discriminator_value, reference_type) in one_of.options() {
        let schema = DataType::ReferenceType(reference_type.clone()).to_json_schema();
        let schema = pin_discriminator(schema, one_of.discriminator_field(), discriminator_value);
        variants.push(schema.clone());
        mapping.insert(discriminator_value.clone(), schema);
    }
    json!({
        "oneOf": variants,
        "discriminator": {
            "propertyName": one_of.discriminator_field(),
            "mapping": mapping
        }
    })
}

/// Pin a variant's discriminator field to its tag value with a `const`, so a
/// JSON-Schema `oneOf` selects exactly one variant. Plain JSON-Schema has no
/// `discriminator` keyword — the validator ignores it — so without the `const`
/// every structurally-equal variant would match and `oneOf` would reject the
/// value as ambiguous.
fn pin_discriminator(schema: Value, field: &str, value: &str) -> Value {
    let object = match schema.as_object() {
        Some(object) if object.contains_key("properties") => object,
        _ => return schema,
    };
    let mut properties = object
        .get("properties")
        .and_then(Value::as_object)
        .cloned()
        .unwrap_or_default();
    properties.insert(field.to_string(), json!({"const": value}));
    let mut object = object.clone();
    object.insert("properties".to_string(), Value::Object(properties));
    Value::Object(object)
}

#[cfg(test)]
mod tests {
    use std::collections::BTreeMap;

    use serde_json::json;

    use super::*;
    use crate::datatypes::value_type::{
        ValueTypeBoolean, ValueTypeDouble, ValueTypeEnum, ValueTypeInteger,
    };

    /// Convert, then assert the result is meta-valid.
    fn schema_of(data_type: DataType) -> Value {
        let schema = data_type.to_json_schema();
        assert!(is_valid(&schema), "meta-invalid schema: {schema}");
        schema
    }

    #[test]
    fn scalar_value_types() {
        assert_eq!(
            schema_of(DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {}))),
            json!({"type": "boolean"})
        );
        assert_eq!(
            schema_of(DataType::ValueType(ValueType::Integer(
                ValueTypeInteger::new()
            ))),
            json!({"type": "integer"})
        );
        assert_eq!(
            schema_of(DataType::ValueType(ValueType::Double(ValueTypeDouble {}))),
            json!({"type": "number"})
        );
    }

    #[test]
    fn strings() {
        let plain = DataType::ValueType(ValueType::String(ValueTypeString::new(None)));
        assert_eq!(schema_of(plain), json!({"type": "string"}));

        let length = DataType::ValueType(ValueType::String(ValueTypeString::new(Some(
            ValueTypeStringValidation::Length {
                min: Some(1),
                max: Some(128),
            },
        ))));
        assert_eq!(
            schema_of(length),
            json!({"type": "string", "minLength": 1, "maxLength": 128})
        );

        let pattern = DataType::ValueType(ValueType::String(ValueTypeString::new(Some(
            ValueTypeStringValidation::Regex {
                regex: "^([a-f0-9]{24})$".to_string(),
            },
        ))));
        assert_eq!(
            schema_of(pattern),
            json!({"type": "string", "pattern": "^([a-f0-9]{24})$"})
        );
    }

    #[test]
    fn enums() {
        let enumeration = ValueTypeEnum::try_new(vec!["a".to_string(), "b".to_string()]).unwrap();
        assert_eq!(
            schema_of(DataType::ValueType(ValueType::Enum(enumeration))),
            json!({"enum": ["a", "b"]})
        );
    }

    #[test]
    fn arrays() {
        let array = DataType::ReferenceType(ReferenceType::new_array(DataType::ValueType(
            ValueType::Boolean(ValueTypeBoolean {}),
        )));
        assert_eq!(
            schema_of(array),
            json!({"type": "array", "items": {"type": "boolean"}})
        );
    }

    #[test]
    fn optional_is_nullable() {
        let optional = DataType::ReferenceType(ReferenceType::new_optional(DataType::ValueType(
            ValueType::Integer(ValueTypeInteger::new()),
        )));
        assert_eq!(schema_of(optional), json!({"type": ["integer", "null"]}));
    }

    #[test]
    fn optional_object_keeps_properties() {
        let mut properties = BTreeMap::new();
        properties.insert(
            "id".to_string(),
            DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})),
        );
        let object = DataType::ReferenceType(ReferenceType::try_new_object(properties).unwrap());
        let optional = DataType::ReferenceType(ReferenceType::new_optional(object));
        assert_eq!(
            schema_of(optional),
            json!({
                "type": ["object", "null"],
                "properties": {"id": {"type": "boolean"}},
                "required": ["id"]
            })
        );
    }

    #[test]
    fn optional_one_of_uses_any_of_fallback() {
        let mut properties = BTreeMap::new();
        properties.insert(
            "id".to_string(),
            DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})),
        );
        let mut options = BTreeMap::new();
        options.insert(
            "default".to_string(),
            ReferenceType::try_new_object(properties).unwrap(),
        );
        let one_of = DataType::ReferenceType(
            ReferenceType::try_new_one_of("type".to_string(), options).unwrap(),
        );
        let optional = DataType::ReferenceType(ReferenceType::new_optional(one_of));
        let schema = schema_of(optional);
        let variant = json!({
            "oneOf": [{
                "type": "object",
                "properties": {
                    "id": {"type": "boolean"},
                    "type": {"const": "default"}
                },
                "required": ["id"]
            }],
            "discriminator": {
                "propertyName": "type",
                "mapping": {
                    "default": {
                        "type": "object",
                        "properties": {
                            "id": {"type": "boolean"},
                            "type": {"const": "default"}
                        },
                        "required": ["id"]
                    }
                }
            }
        });
        assert_eq!(schema, json!({"anyOf": [variant, {"type": "null"}]}));
    }

    #[test]
    fn empty_object_is_free_form() {
        let object =
            DataType::ReferenceType(ReferenceType::try_new_object(BTreeMap::new()).unwrap());
        assert_eq!(schema_of(object), json!({"type": "object"}));
    }

    #[test]
    fn object_with_properties() {
        let mut properties = BTreeMap::new();
        properties.insert(
            "name".to_string(),
            DataType::ValueType(ValueType::String(ValueTypeString::new(None))),
        );
        properties.insert(
            "age".to_string(),
            DataType::ValueType(ValueType::Integer(ValueTypeInteger::new())),
        );
        let object = DataType::ReferenceType(ReferenceType::try_new_object(properties).unwrap());
        assert_eq!(
            schema_of(object),
            json!({
                "type": "object",
                "properties": {
                    "age": {"type": "integer"},
                    "name": {"type": "string"}
                },
                "required": ["age", "name"]
            })
        );
    }

    #[test]
    fn one_of_keeps_discriminator() {
        let mut properties = BTreeMap::new();
        properties.insert(
            "id".to_string(),
            DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {})),
        );
        let mut options = BTreeMap::new();
        options.insert(
            "default".to_string(),
            ReferenceType::try_new_object(properties).unwrap(),
        );
        let one_of = DataType::ReferenceType(
            ReferenceType::try_new_one_of("type".to_string(), options).unwrap(),
        );
        assert_eq!(
            schema_of(one_of),
            json!({
                "oneOf": [{
                    "type": "object",
                    "properties": {
                        "id": {"type": "boolean"},
                        "type": {"const": "default"}
                    },
                    "required": ["id"]
                }],
                "discriminator": {
                    "propertyName": "type",
                    "mapping": {
                        "default": {
                            "type": "object",
                            "properties": {
                                "id": {"type": "boolean"},
                                "type": {"const": "default"}
                            },
                            "required": ["id"]
                        }
                    }
                }
            })
        );
    }

    #[test]
    fn meta_validation_accepts_valid_and_rejects_invalid() {
        let valid = json!({
            "type": "object",
            "properties": {
                "name": {"type": "string"},
                "age": {"type": "integer", "minimum": 0}
            }
        });
        assert!(is_valid(&valid));
        assert!(validate(&valid).is_ok());

        let invalid = json!({"type": "invalid_type", "minimum": "not_a_number"});
        assert!(!is_valid(&invalid));
        assert!(validate(&invalid).is_err());
    }
}
