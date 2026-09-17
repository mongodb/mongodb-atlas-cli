//! `cli.yaml` (manual hierarchy) + `spec.yaml` -> [`GeneratedCli`].
//!
//! The hierarchy is authored by hand (mirroring `models::hierarchy::Hierarchy`:
//! group -> entity -> create/read/update/delete/list + actions + components),
//! so operations land under the right commands instead of under whatever noun
//! heuristic would guess. Operations never listed do not get generated.

use std::collections::{BTreeMap, BTreeSet};

use models::hierarchy::{Entity, Hierarchy};
use models::operation_id::OperationId;
use openapiv3::OpenAPI;
use openapiv3_resolve::{ResolvedOpenAPI, ResolvedParameter};

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

    // Per-operation raw parameters (path/query/header), used for flag
    // generation. Kept keyed by the spec operationId string.
    let mut flags: BTreeMap<String, Vec<GeneratedFlag>> = BTreeMap::new();
    for (_path, item) in spec.paths().paths.iter() {
        for (_verb, operation) in item.iter() {
            let Some(op_id) = operation.operation_id.as_deref() else {
                continue;
            };
            let operation_flags = operation
                .parameters
                .iter()
                .filter_map(|p| match &**p {
                    ResolvedParameter::Cookie { .. } => None,
                    _ => {
                        let data = p.parameter_data();
                        Some(GeneratedFlag {
                            ident: flag_ident(&data.name),
                            description: data.description.clone().unwrap_or_default(),
                            required: data.required,
                        })
                    }
                })
                .collect();
            flags.insert(op_id.to_owned(), operation_flags);
        }
    }

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
                &flags,
                &options.excluded_operation_ids,
                &mut referenced,
            )?;
            let mut components = Vec::new();
            for (component_name, component) in &entity.components {
                let component_operations = component_operations(
                    component,
                    &parsed.operations,
                    &flags,
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
    flags: &BTreeMap<String, Vec<GeneratedFlag>>,
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
        push_operation(&mut out, slot, op_id, operations, flags, excluded, referenced)?;
    }

    for (action, op_id) in &entity.actions {
        push_operation(&mut out, action, op_id, operations, flags, excluded, referenced)?;
    }

    Ok(out)
}

/// read/update/actions of one component, variant names from the slot only
/// (they live under the component's own subcommand enum).
fn component_operations(
    component: &models::hierarchy::Component,
    operations: &BTreeMap<OperationId, models::Operation>,
    flags: &BTreeMap<String, Vec<GeneratedFlag>>,
    excluded: &BTreeSet<String>,
    referenced: &mut BTreeSet<String>,
) -> Result<Vec<GeneratedOperation>, CodegenError> {
    let mut out = Vec::new();
    if let Some(op_id) = component.read.as_ref() {
        push_operation(&mut out, "read", op_id, operations, flags, excluded, referenced)?;
    }
    if let Some(op_id) = component.update.as_ref() {
        push_operation(&mut out, "update", op_id, operations, flags, excluded, referenced)?;
    }
    for (action, op_id) in &component.actions {
        push_operation(&mut out, action, op_id, operations, flags, excluded, referenced)?;
    }
    Ok(out)
}

fn push_operation(
    out: &mut Vec<GeneratedOperation>,
    slot: &str,
    op_id: &OperationId,
    operations: &BTreeMap<OperationId, models::Operation>,
    flags: &BTreeMap<String, Vec<GeneratedFlag>>,
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
    let operation_flags = flags.get(&op_id_string).cloned().unwrap_or_default();

    if operation.versions.is_empty() {
        eprintln!("Skipping operation `{op_id_string}`: no versioned media types");
        return Ok(());
    }

    let variant_name = ident_of(slot);

    let probe_base = pascal_case_of(op_id);
    let versions = operation
        .versions
        .keys()
        .map(|version| {
            let variant_ident = match version {
                models::versioned_mediatype::Version::Stable(date)
                | models::versioned_mediatype::Version::Upcoming(date) => {
                    format!("V{}", date.to_string().replace('-', ""))
                }
                models::versioned_mediatype::Version::Preview => "Preview".to_owned(),
            };
            GeneratedVersion {
                struct_ident: format!("{probe_base}{variant_ident}"),
                variant_ident,
            }
        })
        .collect();

    out.push(GeneratedOperation {
        variant_name,
        description: operation.description.clone(),
        probe_ident: format!("{probe_base}Probe"),
        version_enum_ident: format!("{probe_base}Version"),
        versions,
        flags: operation_flags,
    });
    Ok(())
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
