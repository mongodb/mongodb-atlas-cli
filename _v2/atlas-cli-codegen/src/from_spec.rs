//! `cli.yaml` (manual hierarchy) + `spec.yaml` -> [`GeneratedCli`].
//!
//! The hierarchy is authored by hand (mirroring `models::hierarchy::Hierarchy`:
//! group -> entity -> create/read/update/delete/list + actions + components),
//! so operations land under the right commands instead of under whatever noun
//! heuristic would guess. Operations never listed do not get generated.

use std::collections::{BTreeMap, BTreeSet};

use models::datatypes::DataType;
use models::datatypes::reference_types::ReferenceType;
use models::datatypes::value_type::ValueType;
use models::hierarchy::{Entity, Hierarchy};
use models::operation::ParamIn;
use models::operation_id::OperationId;
use openapiv3::OpenAPI;
use openapiv3_resolve::ResolvedOpenAPI;

use crate::ir::*;
use crate::{CodegenError, CodegenOptions};

/// Parse the two yaml documents and build the full CLI intermediate.
pub fn from_config(
    spec_yaml: &str,
    hierarchy_yaml: &str,
    options: &CodegenOptions,
) -> Result<GeneratedCli, CodegenError> {
    let openapi: OpenAPI = serde_yaml::from_str(spec_yaml)
        .map_err(|error| CodegenError::InvalidOpenAPISpec(error.to_string()))?;
    let resolved = ResolvedOpenAPI::try_from(&openapi)
        .map_err(|error| CodegenError::InvalidOpenAPISpec(error.to_string()))?;
    let hierarchy: Hierarchy = serde_yaml::from_str(hierarchy_yaml)
        .map_err(CodegenError::InvalidHierarchy)?;
    build(&resolved, &hierarchy, options)
}

/// Build the IR from an already-resolved spec and a manual hierarchy.
fn build(
    spec: &ResolvedOpenAPI,
    hierarchy: &Hierarchy,
    options: &CodegenOptions,
) -> Result<GeneratedCli, CodegenError> {
    let parsed = models::Spec::from_resolved_openapi_spec(spec)?;

    // Every operation must be covered by the hierarchy or explicitly
    // excluded; tracks which ids the hierarchy references.
    let mut referenced = BTreeSet::new();

    let mut groups = Vec::new();
    for (group_name, group) in hierarchy.groups() {
        let mut entities = Vec::new();
        for (entity_name, entity) in group.entities() {
            let operations = entity_operations(
                entity,
                &parsed.operations,
                &options.excluded_operation_ids,
                &mut referenced,
            )?;
            let mut components = Vec::new();
            for (component_name, component) in &entity.components {
                let component_operations = component_operations(
                    component,
                    &parsed.operations,
                    &options.excluded_operation_ids,
                    &mut referenced,
                )?;
                if component_operations.is_empty() {
                    continue;
                }
                components.push(GeneratedComponent {
                    id: component_name.clone(),
                    ident: ident_of(component_name),
                    description: component.description.clone(),
                    operations: component_operations,
                });
            }
            if operations.is_empty() && components.is_empty() {
                continue;
            }
            entities.push(GeneratedEntity {
                id: entity_name.clone(),
                ident: ident_of(entity_name),
                description: entity.description.clone(),
                operations,
                components,
            });
        }
        if entities.is_empty() {
            continue;
        }
        groups.push(GeneratedGroup {
            name: group_name.clone(),
            ident: ident_of(group_name),
            description: group.description.clone(),
            entities,
        });
    }

    let uncovered = parsed
        .operations
        .keys()
        .map(|id| id.to_string())
        .filter(|id| !referenced.contains(id) && !options.excluded_operation_ids.contains(id))
        .collect::<Vec<_>>();
    if !uncovered.is_empty() {
        return Err(CodegenError::UncoveredOperations {
            operation_ids: uncovered,
        });
    }

    Ok(GeneratedCli { groups })
}

/// Direct operations of one entity: create/read/update/delete/list, then
/// actions. Components are built separately.
fn entity_operations(
    entity: &Entity,
    operations: &BTreeMap<OperationId, models::Operation>,
    excluded: &BTreeSet<String>,
    referenced: &mut BTreeSet<String>,
) -> Result<Vec<GeneratedOperation>, CodegenError> {
    let mut out = Vec::new();

    for (slot, op_id) in [
        ("create", entity.create.as_ref()),
        ("read", entity.read.as_ref()),
        ("update", entity.update.as_ref()),
        ("delete", entity.delete.as_ref()),
        ("list", entity.list.as_ref()),
    ] {
        let Some(op_id) = op_id else {
            continue;
        };
        push_operation(
            &mut out,
            slot,
            OperationKind::Crud,
            op_id,
            operations,
            excluded,
            referenced,
        )?;
    }

    for (action, op_id) in &entity.actions {
        push_operation(
            &mut out,
            action,
            OperationKind::Action,
            op_id,
            operations,
            excluded,
            referenced,
        )?;
    }

    Ok(out)
}

/// read/update/actions of one component, variant names from the slot only
/// (they live under the component's own subcommand enum).
fn component_operations(
    component: &models::hierarchy::Component,
    operations: &BTreeMap<OperationId, models::Operation>,
    excluded: &BTreeSet<String>,
    referenced: &mut BTreeSet<String>,
) -> Result<Vec<GeneratedOperation>, CodegenError> {
    let mut out = Vec::new();
    if let Some(op_id) = component.read.as_ref() {
        push_operation(
            &mut out,
            "read",
            OperationKind::Crud,
            op_id,
            operations,
            excluded,
            referenced,
        )?;
    }
    if let Some(op_id) = component.update.as_ref() {
        push_operation(
            &mut out,
            "update",
            OperationKind::Crud,
            op_id,
            operations,
            excluded,
            referenced,
        )?;
    }
    for (action, op_id) in &component.actions {
        push_operation(
            &mut out,
            action,
            OperationKind::Action,
            op_id,
            operations,
            excluded,
            referenced,
        )?;
    }
    Ok(out)
}

fn push_operation(
    out: &mut Vec<GeneratedOperation>,
    slot: &str,
    kind: OperationKind,
    op_id: &OperationId,
    operations: &BTreeMap<OperationId, models::Operation>,
    excluded: &BTreeSet<String>,
    referenced: &mut BTreeSet<String>,
) -> Result<(), CodegenError> {
    let op_id_string = op_id.to_string();
    if excluded.contains(&op_id_string) {
        return Ok(());
    }
    let operation = operations
        .get(op_id)
        .ok_or_else(|| CodegenError::OperationNotFoundInSpec {
            op_id: op_id_string.clone(),
        })?;
    referenced.insert(op_id_string.clone());

    if operation.versions.is_empty() {
        tracing::trace!(op_id = %op_id_string, "skipping operation: no versioned media types");
        return Ok(());
    }

    let operation_flags = operation
        .located_parameters()
        .map(|(location, name, description, required)| GeneratedFlag {
            name: name.to_owned(),
            ident: flag_ident(name),
            description: description.to_owned(),
            required,
            list: false,
            location: match location {
                ParamIn::Path => FlagLocation::Path,
                ParamIn::Query => FlagLocation::Query,
                ParamIn::Header => FlagLocation::Header,
            },
            value_kind: FlagValueKind::String,
        })
        .collect();

    let variant_name = ident_of(slot);

    let probe_base = pascal_case_of(op_id);
    let versions = operation
        .versions
        .iter()
        .map(|(version, operation_version)| {
            let (variant_ident, api_version) = match version {
                models::versioned_mediatype::Version::Stable(date) => (
                    format!("V{}", date.to_string().replace('-', "")),
                    ApiVersion::Stable(
                        u32::from(date.year),
                        u32::from(date.month),
                        u32::from(date.day),
                    ),
                ),
                models::versioned_mediatype::Version::Upcoming(date) => (
                    format!("V{}", date.to_string().replace('-', "")),
                    ApiVersion::Upcoming(
                        u32::from(date.year),
                        u32::from(date.month),
                        u32::from(date.day),
                    ),
                ),
                models::versioned_mediatype::Version::Preview => {
                    ("Preview".to_owned(), ApiVersion::Preview)
                }
            };
            let body_flags = operation_version
                .request_body
                .as_ref()
                .map(request_body_flags)
                .unwrap_or_default();
            let body_schema = operation_version
                .request_body
                .as_ref()
                .map(|body| body.to_json_schema().to_string());
            GeneratedVersion {
                struct_ident: format!("{probe_base}{variant_ident}"),
                variant_ident,
                body_flags,
                body_schema,
                api_version,
            }
        })
        .collect();

    out.push(GeneratedOperation {
        variant_name,
        description: operation.description.clone(),
        kind,
        probe_ident: format!("{probe_base}Probe"),
        version_enum_ident: format!("{probe_base}Version"),
        versions,
        flags: operation_flags,
        method: operation.http_verb.to_string(),
        url_template: operation.url_template(),
    });
    Ok(())
}

/// The JSON scalar kind of a value type: what `serde_json::json!(...)` should
/// produce when the flag is turned into a request body.
fn value_kind(value_type: &ValueType) -> FlagValueKind {
    match value_type {
        ValueType::Boolean(_) => FlagValueKind::Boolean,
        ValueType::Integer(_) => FlagValueKind::Integer,
        ValueType::Double(_) => FlagValueKind::Double,
        ValueType::String(_) | ValueType::Enum(_) => FlagValueKind::String,
    }
}

/// Request body -> flat flags.
///
/// Only supports simple bodies: an object whose leaves are value types or
/// arrays of value types. Objects are flattened (`person.first_name` ->
/// `--person-first-name`). Unsupported leaves (arrays of objects, oneOf,
/// nested arrays, ...) are skipped individually with a warning; the supported
/// leaves still generate flags.
fn body_flag(name: &str, required: bool, list: bool, value_kind: FlagValueKind) -> GeneratedFlag {
    GeneratedFlag {
        name: name.to_owned(),
        ident: flag_ident(name),
        description: String::new(),
        required,
        list,
        location: FlagLocation::Body,
        value_kind,
    }
}

fn request_body_flags(body: &DataType) -> Vec<GeneratedFlag> {
    let body = match body {
        DataType::ReferenceType(ReferenceType::Optional(optional)) => optional.data_type(),
        body => body,
    };
    let DataType::ReferenceType(ReferenceType::Object(object)) = body else {
        tracing::trace!("skipping request body flags: unsupported body shape");
        return Vec::new();
    };
    let mut flags = Vec::new();
    for (name, property) in object.properties() {
        request_body_flags_for_property(&mut flags, name, property, true);
    }
    flags
}

/// Collect the flat flags under `name`. `required` is inherited from the
/// enclosing (already-unwrapped) property: an optional property makes every
/// deeper leaf optional. Unsupported leaves are skipped, keep the rest.
fn request_body_flags_for_property(
    flags: &mut Vec<GeneratedFlag>,
    name: &str,
    datatype: &DataType,
    required: bool,
) {
    match datatype {
        DataType::ValueType(value_type) => {
            flags.push(body_flag(name, required, false, value_kind(value_type)));
        }
        DataType::ReferenceType(ReferenceType::Optional(optional)) => {
            request_body_flags_for_property(flags, name, optional.data_type(), false);
        }
        DataType::ReferenceType(ReferenceType::Array(array)) => match array.entries_type() {
            // ponytail: arrays of objects are unsupported by design.
            DataType::ValueType(value_type) => {
                flags.push(body_flag(name, required, true, value_kind(value_type)));
            }
            _ => tracing::trace!(
                flag = name,
                "skipping request body flag: only arrays of value types are supported"
            ),
        },
        DataType::ReferenceType(ReferenceType::Object(object)) => {
            for (sub_name, sub_type) in object.properties() {
                request_body_flags_for_property(
                    flags,
                    &format!("{name}.{sub_name}"),
                    sub_type,
                    required,
                );
            }
        }
        DataType::ReferenceType(ReferenceType::OneOf(_)) => tracing::trace!(
            flag = name,
            "skipping request body flag: unions are not supported"
        ),
    }
}

/// PascalCase of the whole operation id: verb capitalized + every noun.
fn pascal_case_of(op_id: &models::operation_id::OperationId) -> String {
    let mut out = op_id.verb.as_ref().to_owned();
    out.replace_range(0..1, &out[0..1].to_ascii_uppercase().to_string());
    for noun in &op_id.nouns {
        out.push_str(&noun.to_string());
    }
    out
}

/// Any string to a PascalCase identifier: split on non-alphanumeric,
/// capitalize each word, keep abbreviations intact.
fn ident_of(raw: &str) -> String {
    let mut out = String::new();
    for word in raw.split(|c: char| !c.is_ascii_alphanumeric()) {
        if word.is_empty() {
            continue;
        }
        let mut chars = word.chars();
        if let Some(first) = chars.next() {
            out.push(first.to_ascii_uppercase());
        }
        out.extend(chars);
    }
    if out.is_empty() {
        out = "_".to_owned();
    }
    out
}

/// Keep a flag name a valid and non-reserved field identifier.
fn flag_ident(name: &str) -> String {
    let mut sanitized = name
        .chars()
        .map(|c| if c.is_ascii_alphanumeric() || c == '_' { c } else { '_' })
        .collect::<String>();
    if sanitized.is_empty() {
        sanitized.push('_');
    }
    if syn::parse_str::<syn::Ident>(&sanitized).is_ok() && !KEYWORDS.contains(&sanitized.as_str()) {
        sanitized
    } else {
        format!("_{sanitized}")
    }
}

const KEYWORDS: &[&str] = &[
    "as", "async", "await", "break", "const", "continue", "crate", "dyn", "else", "enum",
    "extern", "false", "fn", "for", "if", "impl", "in", "let", "loop", "match", "mod", "move",
    "mut", "pub", "ref", "return", "self", "Self", "static", "struct", "super", "trait", "true",
    "try", "type", "unsafe", "use", "where", "while",
];

#[cfg(test)]
mod tests {
    use super::*;
    use models::datatypes::value_type::{ValueType, ValueTypeString};

    fn str_ty() -> DataType {
        DataType::ValueType(ValueType::String(ValueTypeString::new(None)))
    }

    fn object(props: BTreeMap<String, DataType>) -> DataType {
        DataType::ReferenceType(ReferenceType::try_new_object(props).unwrap())
    }

    fn optional(datatype: DataType) -> DataType {
        DataType::ReferenceType(ReferenceType::new_optional(datatype))
    }

    fn flag(ident: &str, required: bool, list: bool) -> GeneratedFlag {
        dot_flag(ident, ident, required, list, FlagValueKind::String)
    }

    /// A flag whose OpenAPI name differs from its (mangled) field ident, e.g.
    /// the flattened `person.first_name` property.
    fn dot_flag(
        name: &str,
        ident: &str,
        required: bool,
        list: bool,
        value_kind: FlagValueKind,
    ) -> GeneratedFlag {
        GeneratedFlag {
            name: name.to_owned(),
            ident: ident.to_owned(),
            description: String::new(),
            required,
            list,
            location: FlagLocation::Body,
            value_kind,
        }
    }

    #[test]
    fn keeps_required_and_optional_scalars() {
        let body = object(BTreeMap::from([
            ("name".into(), str_ty()),
            ("region".into(), optional(str_ty())),
        ]));
        assert_eq!(
            request_body_flags(&body),
            vec![flag("name", true, false), flag("region", false, false)]
        );
    }

    #[test]
    fn array_of_value_types_is_a_list_flag() {
        let body = object(BTreeMap::from([(
            "providers".into(),
            DataType::ReferenceType(ReferenceType::new_array(str_ty())),
        )]));
        assert_eq!(request_body_flags(&body), vec![flag("providers", true, true)]);
    }

    #[test]
    fn flattens_object_properties() {
        let body = object(BTreeMap::from([(
            "person".into(),
            object(BTreeMap::from([
                ("first_name".into(), str_ty()),
                ("last_name".into(), optional(str_ty())),
            ])),
        )]));
        assert_eq!(
            request_body_flags(&body),
            vec![
                dot_flag("person.first_name", "person_first_name", true, false, FlagValueKind::String),
                dot_flag("person.last_name", "person_last_name", false, false, FlagValueKind::String),
            ]
        );
    }

    #[test]
    fn optional_property_makes_flattened_leaves_optional() {
        let body = object(BTreeMap::from([(
            "person".into(),
            optional(object(BTreeMap::from([("first_name".into(), str_ty())]))),
        )]));
        assert_eq!(
            request_body_flags(&body),
            vec![dot_flag("person.first_name", "person_first_name", false, false, FlagValueKind::String)]
        );
    }

    #[test]
    fn skips_unsupported_leaves_keeps_supported_ones() {
        let body = object(BTreeMap::from([
            ("name".into(), str_ty()),
            // Array of objects: unsupported, skipped.
            (
                "tags".into(),
                DataType::ReferenceType(ReferenceType::new_array(object(BTreeMap::new()))),
            ),
        ]));
        assert_eq!(request_body_flags(&body), vec![flag("name", true, false)]);
    }

    #[test]
    fn skips_non_object_bodies() {
        assert_eq!(request_body_flags(&str_ty()), vec![]);
        assert_eq!(
            request_body_flags(&DataType::ReferenceType(ReferenceType::new_array(str_ty()))),
            vec![]
        );
    }

    fn typed(value_type: ValueType) -> DataType {
        DataType::ValueType(value_type)
    }

    #[test]
    fn value_kinds_follow_the_leaf_type() {
        use models::datatypes::value_type::{
            ValueTypeBoolean, ValueTypeDouble, ValueTypeInteger,
        };
        let body = object(BTreeMap::from([
            ("enabled".into(), typed(ValueType::Boolean(ValueTypeBoolean {}))),
            ("port".into(), typed(ValueType::Integer(ValueTypeInteger {}))),
            ("ratio".into(), typed(ValueType::Double(ValueTypeDouble {}))),
        ]));
        assert_eq!(
            request_body_flags(&body),
            vec![
                dot_flag("enabled", "enabled", true, false, FlagValueKind::Boolean),
                dot_flag("port", "port", true, false, FlagValueKind::Integer),
                dot_flag("ratio", "ratio", true, false, FlagValueKind::Double),
            ]
        );
    }

    #[test]
    fn value_kind_of_array_entries_is_carried() {
        use models::datatypes::value_type::ValueTypeInteger;
        let body = object(BTreeMap::from([(
            "ports".into(),
            DataType::ReferenceType(ReferenceType::new_array(typed(ValueType::Integer(
                ValueTypeInteger {},
            )))),
        )]));
        assert_eq!(
            request_body_flags(&body),
            vec![dot_flag("ports", "ports", true, true, FlagValueKind::Integer)]
        );
    }
}
