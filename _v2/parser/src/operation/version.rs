use std::collections::{BTreeMap, HashSet};

use models::datatypes::DataType;
use models::datatypes::reference_types::ReferenceType;
use models::datatypes::value_type::{
    ValueType, ValueTypeBoolean, ValueTypeDouble, ValueTypeEnum, ValueTypeInteger, ValueTypeString,
    ValueTypeStringValidation,
};
use models::versioned_mediatype::{MediaType, Version};
use openapiv3_resolve::{
    NestedSchema, ResolvedDiscriminator, ResolvedMediaType, ResolvedSchema, ResolvedSchemaKind,
    ResolvedType, indexmap::IndexMap,
};
use serde::{Deserialize, Serialize};
use thiserror::Error;

/// A version of an operation (determined by version header, example: application/vnd.atlas.2024-10-23+json)
/// A version can have multiple response types, for example `getOrgBillingCostExplorerUsage` can both return csv and json, different schemas
///
/// Limitations:
/// - we only support JSON requests, so maximum 1 request body
#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct OperationVersion {
    pub version: Version,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub request_body: Option<DataType>,
    pub response_bodies: BTreeMap<MediaType, DataType>,
}

#[derive(Debug, Error)]
pub enum OperationVersionParseError {
    #[error("Media type has no schema")]
    MissingSchema,
    #[error("Array schema is missing its items")]
    MissingArrayItems,
    #[error("Union option {0} is not a reference type")]
    UnionOptionNotReference(String),
    #[error("Invalid enum: {0}")]
    InvalidEnum(String),
    #[error("String length constraint does not fit in u32")]
    InvalidLength,
    #[error("Unsupported schema kind")]
    UnsupportedSchemaKind,
    #[error(transparent)]
    ReferenceTypeTryNewError(#[from] models::datatypes::reference_types::ReferenceTypeTryNewError),
}

impl OperationVersion {
    pub fn from_version_requests_responses(
        version: Version,
        request_bodies: Option<BTreeMap<MediaType, &ResolvedMediaType>>,
        response_bodies: BTreeMap<MediaType, &ResolvedMediaType>,
    ) -> Result<Self, OperationVersionParseError> {
        // Parse the JSON request body
        let request_body_datatype = if let Some(request_bodies) = request_bodies {
            match request_bodies.get(&MediaType::Json) {
                Some(resolved_media_type) if resolved_media_type.schema.is_some() => {
                    Some(Self::resolved_media_type_to_datatype(resolved_media_type)?)
                }
                Some(_) => {
                    eprintln!("Skipping JSON request body without a schema");
                    None
                }
                None => {
                    // Only non-JSON media types (e.g. gzip log bodies); the
                    // model supports JSON requests only.
                    eprintln!("Skipping request body: only JSON media types are supported");
                    None
                }
            }
        } else {
            None
        };

        // Parse all response bodies
        let mut response_bodies_datatypes = BTreeMap::new();
        for (media_type, resolved_media_type) in response_bodies {
            if resolved_media_type.schema.is_none() {
                eprintln!("Skipping response media type without a schema: {media_type:?}");
                continue;
            }
            let data_type = Self::resolved_media_type_to_datatype(resolved_media_type)?;
            response_bodies_datatypes.insert(media_type, data_type);
        }

        Ok(Self {
            version,
            request_body: request_body_datatype,
            response_bodies: response_bodies_datatypes,
        })
    }

    fn resolved_media_type_to_datatype(
        resolved_media_type: &ResolvedMediaType,
    ) -> Result<DataType, OperationVersionParseError> {
        let schema = resolved_media_type
            .schema
            .as_ref()
            .ok_or(OperationVersionParseError::MissingSchema)?;
        let mut visited = HashSet::new();
        Self::resolved_schema_to_datatype(schema, &mut visited)
    }

    /// Converts a resolved schema into a [`DataType`].
    ///
    /// Conversion rules:
    /// - a discriminator mapping defines a [`ReferenceType::OneOf`] even
    ///   without a `oneOf`/`anyOf` list (quirk 1.1)
    /// - `oneOf`: mapping, or a field each alternative pins with a distinct
    ///   single-value enum, or a merged enum, or a merged object
    /// - `anyOf`: like `oneOf`, and without a discriminator the alternatives'
    ///   objects merge with optional fields (`AnyOf = Object with Optional fields`)
    /// - `allOf` merges the alternatives into one object
    /// - `type: integer` stays an integer, `type: number` becomes a double
    /// - a free-form `type: object` becomes an object with no properties
    /// - a schema that is already being converted on this path (a recursive
    ///   reference) becomes an object with no properties
    /// - a `nullable` schema is wrapped in [`ReferenceType::Optional`]
    fn resolved_schema_to_datatype(
        schema: &ResolvedSchema,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let pointer = (schema as *const ResolvedSchema) as usize;
        if !visited.insert(pointer) {
            // A reference back into the schema that contains this one; there
            // is no finite type (quirk 1.5 / org invoices).
            return Self::untyped_object();
        }

        let datatype = if let Some(discriminator) = &schema.schema_data.discriminator
            && !discriminator.mapping.is_empty()
        {
            Self::union_from_discriminator(discriminator, visited)?
        } else {
            match &schema.schema_kind {
                ResolvedSchemaKind::Any(any) => Self::resolved_any_schema_to_datatype(any, visited),
                ResolvedSchemaKind::Type(typ) => Self::resolved_type_to_datatype(typ, visited),
                ResolvedSchemaKind::OneOf { one_of } => Self::one_of_to_datatype(one_of, visited),
                ResolvedSchemaKind::AnyOf { any_of } => Self::any_of_to_datatype(any_of, visited),
                ResolvedSchemaKind::AllOf { all_of } => {
                    Self::merge_objects_to_datatype(all_of, visited)
                }
                ResolvedSchemaKind::Not { .. } => {
                    Err(OperationVersionParseError::UnsupportedSchemaKind)
                }
            }?
        };

        visited.remove(&pointer);

        if schema.schema_data.nullable {
            Ok(Self::wrap_optional(datatype))
        } else {
            Ok(datatype)
        }
    }

    fn resolved_any_schema_to_datatype(
        any: &openapiv3_resolve::ResolvedAnySchema,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        // `SchemaKind::Any` absorbs `oneOf`, `anyOf`, `allOf`, `properties`,
        // `enumeration` and `items` into one bag, so decide in priority order.
        if !any.one_of.is_empty() {
            return Self::one_of_to_datatype(&any.one_of, visited);
        }
        if !any.any_of.is_empty() {
            return Self::any_of_to_datatype(&any.any_of, visited);
        }
        if !any.all_of.is_empty() {
            return Self::merge_objects_to_datatype(&any.all_of, visited);
        }
        if !any.properties.is_empty() {
            return Self::object_to_datatype(&any.properties, &any.required, visited);
        }
        if !any.enumeration.is_empty() {
            return Self::enum_values_to_datatype(
                any.enumeration
                    .iter()
                    .filter_map(|value| {
                        value
                            .as_str()
                            .map(str::to_owned)
                            .or_else(|| value.as_number().map(|number| number.to_string()))
                    })
                    .collect(),
            );
        }
        match any.typ.as_deref() {
            Some("string") => {
                Self::string_to_datatype(any.pattern.as_deref(), any.min_length, any.max_length)
            }
            Some("integer") => Ok(DataType::ValueType(ValueType::Integer(ValueTypeInteger {}))),
            Some("number") => Ok(DataType::ValueType(ValueType::Double(ValueTypeDouble {}))),
            Some("boolean") => Ok(DataType::ValueType(ValueType::Boolean(ValueTypeBoolean {}))),
            Some("array") => match &any.items {
                Some(items) => Self::items_to_datatype(items, visited),
                None => Err(OperationVersionParseError::MissingArrayItems),
            },
            // Free-form object (`type: object` alone) or a map
            // (`additionalProperties`) has no declared fields (quirk 5.4).
            Some("object") => Self::untyped_object(),
            _ => Err(OperationVersionParseError::UnsupportedSchemaKind),
        }
    }

    fn resolved_type_to_datatype(
        typ: &ResolvedType,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        match typ {
            ResolvedType::String(string) => {
                let enumeration = string
                    .enumeration
                    .iter()
                    .flatten()
                    .cloned()
                    .collect::<Vec<_>>();
                if enumeration.is_empty() {
                    Self::string_to_datatype(
                        string.pattern.as_deref(),
                        string.min_length,
                        string.max_length,
                    )
                } else {
                    Self::enum_values_to_datatype(enumeration)
                }
            }
            ResolvedType::Integer(integer) => Self::scalar_with_enum(
                integer
                    .enumeration
                    .iter()
                    .flatten()
                    .map(|value| value.to_string()),
                ValueType::Integer(ValueTypeInteger {}),
            ),
            ResolvedType::Number(number) => Self::scalar_with_enum(
                number
                    .enumeration
                    .iter()
                    .flatten()
                    .map(|value| value.to_string()),
                ValueType::Double(ValueTypeDouble {}),
            ),
            ResolvedType::Boolean(boolean) => Self::scalar_with_enum(
                boolean
                    .enumeration
                    .iter()
                    .flatten()
                    .map(|value| value.to_string()),
                ValueType::Boolean(ValueTypeBoolean {}),
            ),
            ResolvedType::Object(object) => {
                Self::object_to_datatype(&object.properties, &object.required, visited)
            }
            ResolvedType::Array(array) => match &array.items {
                Some(items) => Self::items_to_datatype(items, visited),
                None => Err(OperationVersionParseError::MissingArrayItems),
            },
        }
    }

    fn scalar_with_enum(
        enumeration: impl Iterator<Item = String>,
        otherwise: ValueType,
    ) -> Result<DataType, OperationVersionParseError> {
        let values = enumeration.collect::<Vec<_>>();
        if values.is_empty() {
            Ok(DataType::ValueType(otherwise))
        } else {
            Self::enum_values_to_datatype(values)
        }
    }

    /// Converts a `oneOf` into a union of its alternatives.
    ///
    /// A discriminator mapping is handled before this is reached. Otherwise:
    /// plain-enum alternatives merge into one enum; a field that each
    /// alternative pins with a distinct single-value enum becomes the
    /// discriminator; failing that, the object alternatives merge (there is no
    /// way to tell them apart, quirk 1.2).
    fn one_of_to_datatype(
        one_of: &[NestedSchema],
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        if let Some(enumeration) = Self::merge_enumerations(one_of) {
            return Ok(enumeration);
        }
        if let Some((field, options)) = Self::derive_discriminator(one_of) {
            return Self::union_from_discriminated(field, options, visited);
        }
        Self::merge_objects_to_datatype(one_of, visited)
    }

    /// [`ResolvedSchemaKind::AnyOf`]: alternatives carrying a single-value enum
    /// are discriminated by it; otherwise every alternative's object is merged
    /// with optional fields.
    fn any_of_to_datatype(
        any_of: &[NestedSchema],
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        if let Some(enumeration) = Self::merge_enumerations(any_of) {
            return Ok(enumeration);
        }
        if let Some((field, options)) = Self::derive_discriminator(any_of) {
            return Self::union_from_discriminated(field, options, visited);
        }
        Self::merge_objects_to_datatype(any_of, visited)
    }

    fn union_from_discriminator(
        discriminator: &ResolvedDiscriminator,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let field = discriminator.property_name.clone();
        let options = discriminator
            .mapping
            .iter()
            .map(|(value, nested)| (value.clone(), nested));
        Self::union_from_discriminated(field, options, visited)
    }

    /// Builds a [`ReferenceType::OneOf`] from a discriminator field and its
    /// mapping of discriminator value to alternative schema.
    ///
    /// Every child re-declares the discriminator field as a tag (quirk 1.4);
    /// the discriminator already carries it, so it is dropped from the option.
    fn union_from_discriminated<'a>(
        discriminator_field: String,
        options: impl IntoIterator<Item = (String, &'a NestedSchema)>,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let mut option_types = BTreeMap::new();
        for (value, nested) in options {
            let reference = match Self::resolved_schema_to_datatype(&nested.get(), visited)? {
                DataType::ReferenceType(reference) => reference,
                DataType::ValueType(_) => {
                    return Err(OperationVersionParseError::UnionOptionNotReference(value));
                }
            };
            let option = Self::strip_discriminator_tag(reference, &discriminator_field);
            option_types.insert(value, option);
        }
        let one_of = ReferenceType::try_new_one_of(discriminator_field, option_types)?;
        Ok(DataType::ReferenceType(one_of))
    }

    /// Merges alternatives that are each a plain enum into a single enum,
    /// such as `regionName` which splits a string enum across a `oneOf`.
    fn merge_enumerations(alternatives: &[NestedSchema]) -> Option<DataType> {
        let mut values = Vec::new();
        for alternative in alternatives {
            values.extend(enumeration_of(&alternative.get())?);
        }
        values.sort_unstable();
        values.dedup();
        Self::enum_values_to_datatype(values).ok()
    }

    /// Finds a field that discriminates every alternative: it appears in each
    /// alternative as a single-value enum, with a distinct value.
    fn derive_discriminator(
        alternatives: &[NestedSchema],
    ) -> Option<(String, IndexMap<String, &NestedSchema>)> {
        let mut field: Option<String> = None;
        let mut options = IndexMap::new();
        for alternative in alternatives {
            let (candidate_field, value) = single_enum_of(&alternative.get())?;
            if let Some(known) = &field
                && known != &candidate_field
            {
                return None;
            }
            field = Some(candidate_field);
            options.insert(value, alternative);
        }
        Some((field?, options))
    }

    fn items_to_datatype(
        items: &NestedSchema,
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let entries_type = Self::resolved_schema_to_datatype(&items.get(), visited)?;
        Ok(DataType::ReferenceType(ReferenceType::new_array(
            entries_type,
        )))
    }

    fn object_to_datatype(
        properties: &IndexMap<String, NestedSchema>,
        required: &[String],
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let property_types = Self::property_types(properties, required, visited)?;
        Ok(DataType::ReferenceType(ReferenceType::try_new_object(
            property_types,
        )?))
    }

    fn property_types(
        properties: &IndexMap<String, NestedSchema>,
        required: &[String],
        visited: &mut HashSet<usize>,
    ) -> Result<BTreeMap<String, DataType>, OperationVersionParseError> {
        let required = required.iter().cloned().collect::<HashSet<_>>();
        let mut property_types = BTreeMap::new();
        for (name, nested) in properties {
            let schema = nested.get();
            let mut datatype = Self::resolved_schema_to_datatype(&schema, visited)?;
            let optional = !required.contains(name);
            if optional {
                datatype = Self::wrap_optional(datatype);
            }
            property_types.insert(name.clone(), datatype);
        }
        Ok(property_types)
    }

    /// Merges every alternative's object into one object (`allOf`,
    /// or an `anyOf`/`oneOf` without a discriminable field).
    fn merge_objects_to_datatype(
        alternatives: &[NestedSchema],
        visited: &mut HashSet<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let mut property_types = BTreeMap::new();
        for alternative in alternatives {
            Self::collect_object_properties(&alternative.get(), &mut property_types, visited)?;
        }
        Ok(DataType::ReferenceType(ReferenceType::try_new_object(
            property_types,
        )?))
    }

    /// Collects the object's own properties into `property_types`, marking
    /// optional the fields the object does not require. An `allOf` child that
    /// is itself a union contributes the union's shared properties only.
    fn collect_object_properties(
        schema: &ResolvedSchema,
        property_types: &mut BTreeMap<String, DataType>,
        visited: &mut HashSet<usize>,
    ) -> Result<(), OperationVersionParseError> {
        match &schema.schema_kind {
            ResolvedSchemaKind::Type(ResolvedType::Object(object)) => Self::collect_properties(
                &object.properties,
                &object.required,
                property_types,
                visited,
            ),
            ResolvedSchemaKind::Any(any) if !any.all_of.is_empty() => {
                for nested in &any.all_of {
                    if !nested.is_recursive() {
                        Self::collect_object_properties(&nested.get(), property_types, visited)?;
                    }
                }
                Ok(())
            }
            ResolvedSchemaKind::Any(any)
                if !any.properties.is_empty() && any.one_of.is_empty() && any.any_of.is_empty() =>
            {
                Self::collect_properties(&any.properties, &any.required, property_types, visited)
            }
            ResolvedSchemaKind::Any(any) if !any.one_of.is_empty() || !any.any_of.is_empty() => {
                // A union's shared properties describe the base type; its
                // alternatives are siblings and stay out of the merged object.
                if !any.properties.is_empty() {
                    Self::collect_properties(
                        &any.properties,
                        &any.required,
                        property_types,
                        visited,
                    )?;
                }
                Ok(())
            }
            _ => Ok(()),
        }
    }

    fn collect_properties(
        properties: &IndexMap<String, NestedSchema>,
        required: &[String],
        property_types: &mut BTreeMap<String, DataType>,
        visited: &mut HashSet<usize>,
    ) -> Result<(), OperationVersionParseError> {
        let required = required.iter().cloned().collect::<HashSet<_>>();
        for (name, nested) in properties {
            let schema = nested.get();
            let mut datatype = Self::resolved_schema_to_datatype(&schema, visited)?;
            if schema.schema_data.nullable || !required.contains(name) {
                datatype = Self::wrap_optional(datatype);
            }
            property_types.insert(name.clone(), datatype);
        }
        Ok(())
    }

    /// An object with no properties: a free-form `type: object`, a map
    /// (`additionalProperties`), or a break in a recursive cycle.
    fn untyped_object() -> Result<DataType, OperationVersionParseError> {
        Ok(DataType::ReferenceType(ReferenceType::try_new_object(
            BTreeMap::new(),
        )?))
    }

    /// Wraps a value in [`ReferenceType::Optional`], flattening an existing
    /// `Optional`: absent and `null` collapse into the one "may be absent"
    /// state rather than `Option<Option<T>>`.
    fn wrap_optional(datatype: DataType) -> DataType {
        match datatype {
            DataType::ReferenceType(ReferenceType::Optional(_)) => datatype,
            other => DataType::ReferenceType(ReferenceType::new_optional(other)),
        }
    }

    /// Drops the discriminator field from a union option when the child
    /// re-declares it as a scalar tag (a string or an enum, quirk 1.4); the
    /// discriminator carries the same information and the wire value is kept.
    fn strip_discriminator_tag(
        reference: ReferenceType,
        discriminator_field: &str,
    ) -> ReferenceType {
        let ReferenceType::Object(object) = &reference else {
            return reference;
        };
        let tag_is_scalar = match object.properties().get(discriminator_field) {
            Some(DataType::ValueType(ValueType::Enum(_)))
            | Some(DataType::ValueType(ValueType::String(_))) => true,
            Some(DataType::ReferenceType(ReferenceType::Optional(optional))) => matches!(
                optional.data_type(),
                DataType::ValueType(ValueType::Enum(_)) | DataType::ValueType(ValueType::String(_))
            ),
            _ => false,
        };
        if !tag_is_scalar {
            return reference;
        }
        let mut properties = object.properties().clone();
        if properties.len() <= 1 {
            return reference;
        }
        properties.remove(discriminator_field);
        ReferenceType::try_new_object(properties).unwrap_or(reference)
    }

    fn string_to_datatype(
        pattern: Option<&str>,
        min_length: Option<usize>,
        max_length: Option<usize>,
    ) -> Result<DataType, OperationVersionParseError> {
        let validation = if let Some(pattern) = pattern {
            Some(ValueTypeStringValidation::Regex {
                regex: pattern.to_string(),
            })
        } else if min_length.is_some() || max_length.is_some() {
            Some(ValueTypeStringValidation::Length {
                min: Self::length_to_u32(min_length)?,
                max: Self::length_to_u32(max_length)?,
            })
        } else {
            None
        };
        Ok(DataType::ValueType(ValueType::String(
            ValueTypeString::new(validation),
        )))
    }

    fn length_to_u32(length: Option<usize>) -> Result<Option<u32>, OperationVersionParseError> {
        length
            .map(u32::try_from)
            .transpose()
            .map_err(|_| OperationVersionParseError::InvalidLength)
    }

    fn enum_values_to_datatype(
        values: Vec<String>,
    ) -> Result<DataType, OperationVersionParseError> {
        let enumeration = ValueTypeEnum::try_new(values)
            .map_err(|error| OperationVersionParseError::InvalidEnum(error.to_string()))?;
        Ok(DataType::ValueType(ValueType::Enum(enumeration)))
    }
}

/// Every enum value a plain-enum schema contributes, or `None` if the schema
/// is something other than a plain enum (it has properties or combinators).
/// Only a non-empty enumeration counts; a plain string with no enum values is
/// not mergeable.
fn enumeration_of(schema: &ResolvedSchema) -> Option<Vec<String>> {
    match &schema.schema_kind {
        ResolvedSchemaKind::Type(ResolvedType::String(string)) => {
            let values = string
                .enumeration
                .iter()
                .flatten()
                .cloned()
                .collect::<Vec<_>>();
            (!values.is_empty()).then_some(values)
        }
        ResolvedSchemaKind::Type(ResolvedType::Integer(integer)) => {
            let values = integer
                .enumeration
                .iter()
                .flatten()
                .map(|value| value.to_string())
                .collect::<Vec<_>>();
            (!values.is_empty()).then_some(values)
        }
        ResolvedSchemaKind::Type(ResolvedType::Number(number)) => {
            let values = number
                .enumeration
                .iter()
                .flatten()
                .map(|value| value.to_string())
                .collect::<Vec<_>>();
            (!values.is_empty()).then_some(values)
        }
        ResolvedSchemaKind::Type(ResolvedType::Boolean(boolean)) => {
            let values = boolean
                .enumeration
                .iter()
                .flatten()
                .map(|value| value.to_string())
                .collect::<Vec<_>>();
            (!values.is_empty()).then_some(values)
        }
        ResolvedSchemaKind::Any(any)
            if !any.enumeration.is_empty()
                && any.properties.is_empty()
                && any.one_of.is_empty()
                && any.all_of.is_empty()
                && any.any_of.is_empty()
                && any.items.is_none() =>
        {
            Some(
                any.enumeration
                    .iter()
                    .filter_map(|value| {
                        value
                            .as_str()
                            .map(str::to_owned)
                            .or_else(|| value.as_number().map(|number| number.to_string()))
                    })
                    .collect(),
            )
        }
        _ => None,
    }
}

/// The discriminator values of `schema` as `(field, value)` pairs, where the
/// field is an object property pinned to a single-value enum. Walks an
/// `allOf` base but not a union's alternatives.
fn single_enums(schema: &ResolvedSchema) -> Vec<(String, Vec<String>)> {
    match &schema.schema_kind {
        ResolvedSchemaKind::Type(ResolvedType::Object(object)) => {
            enum_properties_of(&object.properties)
        }
        ResolvedSchemaKind::Any(any) if !any.all_of.is_empty() => any
            .all_of
            .iter()
            .filter(|nested| !nested.is_recursive())
            .flat_map(|nested| single_enums(&nested.get()))
            .collect(),
        ResolvedSchemaKind::Any(any)
            if any.properties.is_empty() && any.one_of.is_empty() && any.any_of.is_empty() =>
        {
            Vec::new()
        }
        ResolvedSchemaKind::Any(any) => enum_properties_of(&any.properties),
        ResolvedSchemaKind::AllOf { all_of } => all_of
            .iter()
            .filter(|nested| !nested.is_recursive())
            .flat_map(|nested| single_enums(&nested.get()))
            .collect(),
        _ => Vec::new(),
    }
}

fn enum_properties_of(properties: &IndexMap<String, NestedSchema>) -> Vec<(String, Vec<String>)> {
    properties
        .iter()
        .filter_map(|(name, nested)| {
            let values = enum_strings_of_resolved(&nested.get());
            (!values.is_empty()).then(|| (name.clone(), values))
        })
        .collect()
}

/// Every enum value a property schema contributes.
fn enum_strings_of_resolved(schema: &ResolvedSchema) -> Vec<String> {
    match &schema.schema_kind {
        ResolvedSchemaKind::Type(ResolvedType::String(string)) => {
            string.enumeration.iter().flatten().cloned().collect()
        }
        ResolvedSchemaKind::Type(ResolvedType::Integer(integer)) => integer
            .enumeration
            .iter()
            .flatten()
            .map(|value| value.to_string())
            .collect(),
        ResolvedSchemaKind::Type(ResolvedType::Number(number)) => number
            .enumeration
            .iter()
            .flatten()
            .map(|value| value.to_string())
            .collect(),
        _ => Vec::new(),
    }
}

/// The first single-value enum field on `schema`, as `(field, value)`.
fn single_enum_of(schema: &ResolvedSchema) -> Option<(String, String)> {
    single_enums(schema)
        .into_iter()
        .find_map(|(field, values)| {
            (values.len() == 1)
                .then(|| values.into_iter().next().unwrap())
                .map(|value| (field, value))
        })
}
