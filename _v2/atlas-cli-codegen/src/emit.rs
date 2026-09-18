//! [`GeneratedCli`] -> `TokenStream`: dumb, positional translation.
//!
//! All logic worth testing lives in `from_spec.rs`; this module only maps the
//! intermediate representation onto clap types. The CLI tree is
//! `Cli -> group -> entity -> {direct ops, component -> ops}`; a single-entity
//! group is flattened onto the group command.

use proc_macro2::{Span, TokenStream};
use quote::quote;
use syn::{Ident, LitStr};

use crate::ir::*;

/// Fully-qualified return type shared by every `execute`.
fn execute_ret() -> TokenStream {
    quote! { ::std::result::Result<(), ::std::process::ExitCode> }
}

/// One line in a help section: the subcommand name and its short about text,
/// mirroring what clap would display for the same docs.
#[derive(Clone, Debug, PartialEq, Eq)]
struct HelpEntry {
    name: String,
    about: String,
}

/// The `help` subcommand clap auto-adds, rendered under its own heading.
fn help_section() -> (String, Vec<HelpEntry>) {
    (
        "Help".to_owned(),
        vec![HelpEntry {
            name: "help".to_owned(),
            about: "Print this message or the help of the given subcommand(s)".to_owned(),
        }],
    )
}

/// What clap displays for a doc comment in a section: first paragraph,
/// trimmed, one trailing `.` removed (clap's `remove_period`).
fn help_about(text: &str) -> String {
    let first_paragraph = text.trim().split("\n\n").next().unwrap_or(text).trim();
    if first_paragraph.ends_with('.') && !first_paragraph.ends_with("..") {
        first_paragraph[..first_paragraph.len() - 1].to_owned()
    } else {
        first_paragraph.to_owned()
    }
}

/// `(name, about)` for one operation: subcommand name + first-sentence about.
fn operation_entry(operation: &GeneratedOperation) -> HelpEntry {
    let (about, _) = split_first_sentence(&operation.description);
    HelpEntry {
        name: kebab_case(&operation.variant_name),
        about: help_about(&about),
    }
}

/// `(name, about)` for a component or entity subcommand.
fn named_entry(name: &str, doc: &str) -> HelpEntry {
    HelpEntry {
        name: kebab_case(name),
        about: help_about(doc),
    }
}

/// `Operations`/`Actions` sections for a command's direct operations. Empty
/// categories are skipped; entries are sorted like clap sorts subcommands.
fn operation_sections(operations: &[GeneratedOperation]) -> Vec<(String, Vec<HelpEntry>)> {
    let sorted = |kind: OperationKind| {
        let mut entries: Vec<HelpEntry> = operations
            .iter()
            .filter(|op| op.kind == kind)
            .map(operation_entry)
            .collect();
        entries.sort_by(|a, b| a.name.cmp(&b.name));
        entries
    };
    let mut sections = Vec::new();
    let crud = sorted(OperationKind::Crud);
    if !crud.is_empty() {
        sections.push(("Operations".to_owned(), crud));
    }
    let actions = sorted(OperationKind::Action);
    if !actions.is_empty() {
        sections.push(("Actions".to_owned(), actions));
    }
    sections
}

/// Full section list for one entity's command: operations, components, help.
fn entity_sections(entity: &GeneratedEntity) -> Vec<(String, Vec<HelpEntry>)> {
    let mut sections = operation_sections(&entity.operations);
    let mut components: Vec<HelpEntry> = entity
        .components
        .iter()
        .map(|component| {
            named_entry(
                &component.ident,
                &help_text(&component.description, || {
                    format!("Manage your Atlas CLI {}", kebab_case(&component.id))
                }),
            )
        })
        .collect();
    components.sort_by(|a, b| a.name.cmp(&b.name));
    if !components.is_empty() {
        sections.push(("Components".to_owned(), components));
    }
    sections.push(help_section());
    sections
}

/// Sections for one component's command: operations (`read`/`update`/actions),
/// then help.
fn component_sections(component: &GeneratedComponent) -> Vec<(String, Vec<HelpEntry>)> {
    let mut sections = operation_sections(&component.operations);
    sections.push(help_section());
    sections
}

/// `Entities` + `help` sections for a multi-entity group command.
fn group_sections(group: &GeneratedGroup) -> Vec<(String, Vec<HelpEntry>)> {
    let mut entities: Vec<HelpEntry> = group
        .entities
        .iter()
        .map(|entity| {
            named_entry(
                &entity.ident,
                &help_text(&entity.description, || {
                    format!("Manage your Atlas CLI {}", kebab_case(&entity.id))
                }),
            )
        })
        .collect();
    entities.sort_by(|a, b| a.name.cmp(&b.name));
    vec![("Entities".to_owned(), entities), help_section()]
}

/// A help-override renderer for one generated command.
struct HelpNode {
    /// `render_<path snake>_help`, the generated fn name inside the help module.
    fn_ident: Ident,
    /// The command's own about line (already display-processed).
    about: String,
    /// Kebab-cased path from the CLI root to this command, for the usage line.
    path: Vec<String>,
    /// Help sections in render order.
    sections: Vec<(String, Vec<HelpEntry>)>,
}

/// `cloud-backups`, `compliance-policy` -> `render_cloud_backups_compliance_policy_help`.
fn help_fn_ident(path: &[String]) -> Ident {
    let mut name = path.join("-").replace('-', "_");
    name.push_str("_help");
    ident(&name)
}

/// `#[command(override_help = __atlas_cli_help::render_<path>_help())]` for a
/// generated command struct, so its help renders sectioned instead of clap's
/// single `Commands:` block.
fn help_override_attr(path: &[String]) -> TokenStream {
    let renderer = help_fn_ident(path);
    quote! { #[command(override_help = __atlas_cli_help::#renderer())] }
}

fn node_name(nodes: &mut Vec<HelpNode>, path: &[String], about: String, sections: Vec<(String, Vec<HelpEntry>)>) {
    nodes.push(HelpNode {
        fn_ident: help_fn_ident(path),
        about,
        path: path.to_vec(),
        sections,
    });
}

fn help_nodes(cli: &GeneratedCli) -> Vec<HelpNode> {
    let mut nodes = Vec::new();
    for group in &cli.groups {
        let group_path = vec![kebab_case(&group.name)];
        let group_about = help_about(&help_text(&group.description, || {
            format!("Manage your Atlas CLI {}", group.name.to_lowercase())
        }));
        if group.entities.len() == 1 {
            // Flattened: the group command IS the entity command.
            node_name(
                &mut nodes,
                &group_path,
                group_about,
                entity_sections(&group.entities[0]),
            );
        } else {
            node_name(
                &mut nodes,
                &group_path,
                group_about,
                group_sections(group),
            );
        }
        for entity in &group.entities {
            // A flattened group has no entity level: components hang straight
            // off the group command, so their path skips the entity segment.
            let component_base: Vec<String> = if group.entities.len() == 1 {
                group_path.clone()
            } else {
                vec![group_path[0].clone(), kebab_case(&entity.id)]
            };
            if group.entities.len() > 1 {
                node_name(
                    &mut nodes,
                    &component_base,
                    help_about(&help_text(&entity.description, || {
                        format!("Manage your Atlas CLI {}", kebab_case(&entity.id))
                    })),
                    entity_sections(entity),
                );
            }
            for component in &entity.components {
                let mut component_path = component_base.clone();
                component_path.push(kebab_case(&component.id));
                node_name(
                    &mut nodes,
                    &component_path,
                    help_about(&help_text(&component.description, || {
                        format!("Manage your Atlas CLI {}", kebab_case(&component.id))
                    })),
                    component_sections(component),
                );
            }
        }
    }
    nodes
}

fn section_tokens(sections: &[(String, Vec<HelpEntry>)]) -> Vec<TokenStream> {
    sections
        .iter()
        .map(|(heading, entries)| {
            let entries_tok = entries.iter().map(|entry| {
                let name = &entry.name;
                let about = &entry.about;
                quote! { (#name, #about) }
            });
            quote! { (#heading, &[#(#entries_tok),*]) }
        })
        .collect()
}

/// The `anstyle` renderer + one wrapper fn per overridden command.
fn help_module(nodes: &[HelpNode]) -> TokenStream {
    let wrappers = nodes.iter().map(|node| {
        let fn_ident = &node.fn_ident;
        let about = &node.about;
        let path = &node.path;
        let sections = section_tokens(&node.sections);
        quote! {
            pub(crate) fn #fn_ident() -> StyledStr {
                render(#about, &[#(#path),*], &[#(#sections),*])
            }
        }
    });
    let module = quote! {
        #[allow(dead_code, reason = "renderers are wired into clap derive")]
        pub(crate) mod __atlas_cli_help {
            use ::clap::builder::StyledStr;

            #(#wrappers)*

            fn heading_style() -> ::anstyle::Style {
                ::anstyle::Style::new().bold().underline()
            }

            fn literal_style() -> ::anstyle::Style {
                ::anstyle::Style::new().bold()
            }

            fn bin_name() -> String {
                std::env::args_os()
                    .next()
                    .and_then(|arg| {
                        std::path::Path::new(&arg)
                            .file_name()
                            .map(|name| name.to_string_lossy().into_owned())
                    })
                    .unwrap_or_else(|| "atlas-cli".to_owned())
            }

            /// `Operations`/`Actions`/`Components`/`Entities`/`Help` sections
            /// styled like clap, plus usage and options.
            fn render(about: &str, path: &[&str], sections: &[(&str, &[(&str, &str)])]) -> StyledStr {
                let heading = heading_style().render().to_string();
                let literal = literal_style().render().to_string();
                let reset = literal_style().render_reset().to_string();

                let mut out = StyledStr::new();
                if !about.is_empty() {
                    out.push_str(about);
                    out.push_str("\n\n");
                }

                out.push_str(&format!("{heading}Usage:{reset} "));
                let mut usage = bin_name();
                for segment in path {
                    usage.push(' ');
                    usage.push_str(segment);
                }
                out.push_str(&format!("{literal}{usage}{reset} <COMMAND>\n\n"));

                let max_name = sections
                    .iter()
                    .flat_map(|(_, entries)| entries.iter())
                    .map(|(name, _)| name.chars().count())
                    .max()
                    .unwrap_or(0);

                for (offset, (heading_text, entries)) in sections.iter().enumerate() {
                    if offset > 0 {
                        out.push_str("\n");
                    }
                    out.push_str(&format!("{heading}{heading_text}:{reset}\n"));
                    for (name, about) in *entries {
                        out.push_str(&format!("  {literal}{name}{reset}"));
                        out.push_str(&" ".repeat(max_name.saturating_sub(name.chars().count()) + 2));
                        out.push_str(about);
                        out.push_str("\n");
                    }
                }

                out.push_str("\n");
                out.push_str(&format!("{heading}Options:{reset}\n"));
                out.push_str(&format!("  {literal}-h{reset}, {literal}--help{reset}  Print help\n"));

                out
            }
        }
    };
    module
}

pub fn to_tokens(cli: &GeneratedCli) -> TokenStream {
    let cli_subcommands = cli_subcommands(cli);
    let cli_execute = cli_execute(cli);
    let groups = cli.groups.iter().map(group);
    let help = help_module(&help_nodes(cli));

    let operations = cli
        .groups
        .iter()
        .flat_map(|g| g.entities.iter())
        .flat_map(|e| {
            e.operations
                .iter()
                .chain(e.components.iter().flat_map(|c| c.operations.iter()))
        })
        .collect::<Vec<_>>();
    let any_body = operations
        .iter()
        .any(|op| op.versions.iter().any(|version| version.body_schema.is_some()));
    let body = any_body
        .then(body_module)
        .unwrap_or_default();
    let operations = operations.into_iter().map(operation);

    quote! {
        #cli_subcommands
        #cli_execute
        #help
        #body
        #(#groups)*
        #(#operations)*
    }
}

fn ident(name: &str) -> Ident {
    Ident::new(name, Span::call_site())
}

fn subcommand_enum_ident(group_ident: &str) -> Ident {
    ident(&format!("{group_ident}Subcommands"))
}

fn command_ident(group_ident: &str) -> Ident {
    ident(&format!("{group_ident}Command"))
}

fn entity_command_ident(group_ident: &str, entity_ident: &str) -> Ident {
    ident(&format!("{group_ident}{entity_ident}Command"))
}

fn entity_subcommand_enum_ident(group_ident: &str, entity_ident: &str) -> Ident {
    ident(&format!("{group_ident}{entity_ident}Subcommands"))
}

fn component_command_ident(group_ident: &str, entity_ident: &str, component_ident: &str) -> Ident {
    ident(&format!("{group_ident}{entity_ident}{component_ident}Command"))
}

fn component_subcommand_enum_ident(
    group_ident: &str,
    entity_ident: &str,
    component_ident: &str,
) -> Ident {
    ident(&format!("{group_ident}{entity_ident}{component_ident}Subcommands"))
}

fn cli_subcommands(cli: &GeneratedCli) -> TokenStream {
    let variants = cli.groups.iter().map(|group| {
        let variant = ident(&group.ident);
        let command = command_ident(&group.ident);
        let doc = help_text(&group.description, || {
            format!("Manage your Atlas CLI {}", group.name.to_lowercase())
        });
        quote! {
            #[doc = #doc]
            #variant(#command),
        }
    });

    quote! {
        #[derive(Debug, ::clap::Subcommand)]
        pub enum CliSubCommands {
            #(#variants)*
        }
    }
}

fn cli_execute(cli: &GeneratedCli) -> TokenStream {
    let ret = execute_ret();
    let arms = cli.groups.iter().map(|group| {
        let variant = ident(&group.ident);
        let inner = group_execute(group);
        quote! {
            CliSubCommands::#variant(c) => #inner,
        }
    });

    quote! {
        impl Cli {
            pub async fn execute(self) -> #ret {
                match self.sub_command {
                    #(#arms)*
                }
            }
        }
    }
}

fn group_execute(group: &GeneratedGroup) -> TokenStream {
    let group_enum = subcommand_enum_ident(&group.ident);

    if group.entities.len() == 1 {
        entity_execute(group, &group.entities[0], &group_enum, ident("c"))
    } else {
        let arms = group.entities.iter().map(|entity| {
            let entity_variant = ident(&entity.ident);
            let entity_enum = entity_subcommand_enum_ident(&group.ident, &entity.ident);
            let inner = entity_execute(group, entity, &entity_enum, ident("b"));
            quote! {
                #group_enum::#entity_variant(b) => #inner,
            }
        });
        quote! { match c.sub_command { #(#arms)* } }
    }
}

/// Dispatch over one entity command: direct operation arms + component arms.
fn entity_execute(
    group: &GeneratedGroup,
    entity: &GeneratedEntity,
    entity_enum: &Ident,
    binding: Ident,
) -> TokenStream {
    let mut arms = op_arms(&entity.operations, entity_enum);
    for component in &entity.components {
        let component_variant = ident(&component.ident);
        let component_enum =
            component_subcommand_enum_ident(&group.ident, &entity.ident, &component.ident);
        let component_arms = op_arms(&component.operations, &component_enum);
        arms.push(quote! {
            #entity_enum::#component_variant(d) => match d.sub_command {
                #(#component_arms)*
            },
        });
    }
    quote! { match #binding.sub_command { #(#arms)* } }
}

fn op_arms(operations: &[GeneratedOperation], subcommand_enum: &Ident) -> Vec<TokenStream> {
    operations
        .iter()
        .map(|operation| {
            let variant = ident(&operation.variant_name);
            quote! { #subcommand_enum::#variant(probe) => probe.execute().await, }
        })
        .collect()
}

/// The command structs + subcommand enums of one group and everything under it.
fn group(group: &GeneratedGroup) -> TokenStream {
    let group_command = command_ident(&group.ident);
    let group_enum = subcommand_enum_ident(&group.ident);
    let group_help_attr = help_override_attr(&[kebab_case(&group.name)]);

    let group_struct = quote! {
        #[derive(Debug, ::clap::Args)]
        #group_help_attr
        pub struct #group_command {
            #[clap(subcommand)]
            sub_command: #group_enum,
        }
    };

    if group.entities.len() == 1 {
        // Flattened: the group command IS the entity command.
        let entity = &group.entities[0];
        let entity_variants = entity_variants(group, entity);
        let group_enum_tok = quote! {
            #[derive(Debug, ::clap::Subcommand)]
            pub enum #group_enum {
                #(#entity_variants)*
            }
        };
        let components = component_commands(group, entity);
        quote! {
            #group_struct
            #group_enum_tok
            #(#components)*
        }
    } else {
        let group_entity_variants = group.entities.iter().map(|entity| {
            let entity_variant = ident(&entity.ident);
            let entity_command = entity_command_ident(&group.ident, &entity.ident);
            let doc = help_text(&entity.description, || {
                format!("Manage your Atlas CLI {}", kebab_case(&entity.id))
            });
            quote! {
                #[doc = #doc]
                #entity_variant(#entity_command),
            }
        });
        let group_enum_tok = quote! {
            #[derive(Debug, ::clap::Subcommand)]
            pub enum #group_enum {
                #(#group_entity_variants)*
            }
        };

        let entities = group.entities.iter().map(|entity| {
            let entity_command = entity_command_ident(&group.ident, &entity.ident);
            let entity_enum = entity_subcommand_enum_ident(&group.ident, &entity.ident);
            let entity_help_attr =
                help_override_attr(&[kebab_case(&group.name), kebab_case(&entity.id)]);
            let entity_struct = quote! {
                #[derive(Debug, ::clap::Args)]
                #entity_help_attr
                pub struct #entity_command {
                    #[clap(subcommand)]
                    sub_command: #entity_enum,
                }
            };
            let entity_variants = entity_variants(group, entity);
            let entity_enum_tok = quote! {
                #[derive(Debug, ::clap::Subcommand)]
                pub enum #entity_enum {
                    #(#entity_variants)*
                }
            };
            let components = component_commands(group, entity);
            quote! {
                #entity_struct
                #entity_enum_tok
                #(#components)*
            }
        });

        quote! {
            #group_struct
            #group_enum_tok
            #(#entities)*
        }
    }
}

/// Variants of one entity's subcommand enum: direct operations + components.
fn entity_variants(group: &GeneratedGroup, entity: &GeneratedEntity) -> Vec<TokenStream> {
    let mut variants = operation_variants(&entity.operations);
    for component in &entity.components {
        let component_variant = ident(&component.ident);
        let component_command =
            component_command_ident(&group.ident, &entity.ident, &component.ident);
        let doc = help_text(&component.description, || {
            format!("Manage your Atlas CLI {}", kebab_case(&component.id))
        });
        variants.push(quote! {
            #[doc = #doc]
            #component_variant(#component_command),
        });
    }
    variants
}

/// The per-component command struct + subcommand enum.
fn component_commands(group: &GeneratedGroup, entity: &GeneratedEntity) -> Vec<TokenStream> {
    entity
        .components
        .iter()
        .map(|component| {
            let component_command =
                component_command_ident(&group.ident, &entity.ident, &component.ident);
            let component_enum =
                component_subcommand_enum_ident(&group.ident, &entity.ident, &component.ident);
            let mut component_path = vec![kebab_case(&group.name), kebab_case(&component.id)];
            if group.entities.len() > 1 {
                component_path.insert(1, kebab_case(&entity.id));
            }
            let component_help_attr = help_override_attr(&component_path);
            let component_struct = quote! {
                #[derive(Debug, ::clap::Args)]
                #component_help_attr
                pub struct #component_command {
                    #[clap(subcommand)]
                    sub_command: #component_enum,
                }
            };
            let variants = operation_variants(&component.operations);
            let component_enum_tok = quote! {
                #[derive(Debug, ::clap::Subcommand)]
                pub enum #component_enum {
                    #(#variants)*
                }
            };
            quote! {
                #component_struct
                #component_enum_tok
            }
        })
        .collect()
}

fn operation_variants(operations: &[GeneratedOperation]) -> Vec<TokenStream> {
    operations
        .iter()
        .map(|operation| {
            let variant = ident(&operation.variant_name);
            let probe = ident(&operation.probe_ident);
            let doc_attrs = description_doc_attrs(&operation.description);
            quote! {
                #(#doc_attrs)*
                #variant(#probe),
            }
        })
        .collect()
}

/// Config description, or the given fallback.
fn help_text(description: &Option<String>, fallback: impl FnOnce() -> String) -> String {
    description.clone().unwrap_or_else(fallback)
}

/// Doc-comment attributes for an operation description. clap derives `about`
/// from the first paragraph (up to the first blank line) and `long_about`
/// from everything; a blank line after the first sentence makes the single
/// sentence the short help.
fn description_doc_attrs(description: &str) -> Vec<TokenStream> {
    let (about, rest) = split_first_sentence(description);
    let mut attrs = vec![quote! { #[doc = #about] }];
    if let Some(rest) = rest {
        attrs.push(quote! { #[doc = ""] });
        attrs.push(quote! { #[doc = #rest] });
    }
    attrs
}

/// `description_doc_attrs` plus a trailing blank line and a help note.
fn description_doc_attrs_with_note(description: &str, note: String) -> Vec<TokenStream> {
    let mut attrs = description_doc_attrs(description);
    attrs.push(quote! { #[doc = ""] });
    attrs.push(quote! { #[doc = #note] });
    attrs
}

/// `(first sentence, remaining description)` split at the first `.` followed
/// by whitespace or end-of-string. Single sentence descriptions come back
/// whole with no remainder.
fn split_first_sentence(description: &str) -> (String, Option<String>) {
    let text = description.trim();
    for (i, c) in text.char_indices() {
        if c != '.' {
            continue;
        }
        let after = i + 1;
        let sentence_ends = after >= text.len()
            || text[after..].chars().next().is_some_and(|n| n.is_whitespace());
        if sentence_ends {
            let sentence = text[..after].to_owned();
            let rest = text[after..].trim().to_owned();
            return (sentence, (!rest.is_empty()).then_some(rest));
        }
    }
    (text.to_owned(), None)
}

/// The probe + version enum + per-version structs + execute impls.
fn operation(operation: &GeneratedOperation) -> TokenStream {
    let Some(latest) = operation.versions.last() else {
        // `from_spec` never emits a version-less operation; stay safe anyway.
        return TokenStream::new();
    };
    let ret = execute_ret();
    let probe = ident(&operation.probe_ident);
    let version_enum = ident(&operation.version_enum_ident);
    let default = ident(&latest.variant_ident);

    let probe_struct = quote! {
        #[derive(Debug, ::clap::Args)]
        #[command(disable_help_flag = true)]
        pub struct #probe {
            #[arg(long, value_enum, default_value_t = #version_enum::#default)]
            version: #version_enum,

            /// Raw args for the selected version's command, captured verbatim.
            #[arg(allow_hyphen_values = true)]
            rest: Vec<String>,
        }
    };

    let version_variants = operation.versions.iter().map(|version| {
        let variant = ident(&version.variant_ident);
        quote! { #variant, }
    });
    let version_enum_tok = quote! {
        #[derive(Clone, Copy, Debug, ::clap::ValueEnum)]
        pub enum #version_enum {
            #(#version_variants)*
        }
    };

    let dispatch = dispatch_arms(operation, &version_enum);
    let probe_impl = quote! {
        impl #probe {
            /// Re-parse the captured raw args against the version-specific command.
            pub async fn execute(self) -> #ret {
                match self.version {
                    #(#dispatch)*
                }
            }
        }
    };

    let version_structs = operation.versions.iter().map(|version| {
        let struct_ident = ident(&version.struct_ident);
        let variant = ident(&version.variant_ident);
        let fields = operation
            .flags
            .iter()
            .chain(version.body_flags.iter())
            .map(flag_field);
        let help_doc = format!(
            "Flags differ per API version (--version). This help describes version {}.",
            version.variant_ident
        );
        let doc_attrs = description_doc_attrs_with_note(&operation.description, help_doc);
        let op_impl = operation_impl(operation, version);
        if let Some(schema) = &version.body_schema {
            // Versions with a request body take it from --file/stdin XOR the
            // flat body flags; clio::Input is not Clone so these structs lose
            // the Clone derive. The prepared body rides in a clap-skipped
            // field, filled in by execute() before the client runs.
            let schema_lit = LitStr::new(schema, Span::call_site());
            let inserts = version.body_flags.iter().map(body_flag_insert);
            let has_flags = has_flags_expr(&version.body_flags);
            quote! {
                #(#doc_attrs)*
                #[derive(Debug, ::clap::Parser)]
                pub struct #struct_ident {
                    #[arg(long, value_enum, default_value_t = #version_enum::#variant)]
                    version: #version_enum,
                    #(#fields)*

                    /// Read the request body from a JSON file, or `-` for stdin.
                    #[arg(long)]
                    file: Option<::clio::Input>,

                    #[arg(skip)]
                    body: Option<::bytes::Bytes>,
                }

                impl #struct_ident {
                    /// Run the operation through the atlas client and print
                    /// the JSON response; exit non-zero on any error.
                    pub async fn execute(mut self) -> #ret {
                        ::tracing::debug!(args = ?self, "executing atlas operation");
                        self.body = self.prepare_body().map_err(|error| {
                            eprintln!("{error}");
                            ::std::process::ExitCode::FAILURE
                        })?;
                        let client = ::mongodb_atlas_cli::atlas::client::AtlasClient::from_defaults()
                            .map_err(|error| {
                                eprintln!("{error}");
                                ::std::process::ExitCode::FAILURE
                            })?;
                        let response = client.execute(self).await.map_err(|error| {
                            eprintln!("{error}");
                            ::std::process::ExitCode::FAILURE
                        })?;
                        println!(
                            "{}",
                            ::serde_json::to_string_pretty(&response).map_err(|error| {
                                eprintln!("{error}");
                                ::std::process::ExitCode::FAILURE
                            })?
                        );
                        Ok(())
                    }

                    /// Build the request body from `--file`/stdin or from the
                    /// flat body flags (never both), validated against the
                    /// version's JSON Schema.
                    fn prepare_body(&mut self) -> ::std::result::Result<Option<::bytes::Bytes>, String> {
                        const BODY_SCHEMA: &str = #schema_lit;
                        crate::__atlas_cli_body::build_body(
                            self.file.as_mut(),
                            Some(BODY_SCHEMA),
                            #has_flags,
                            |object| {
                                #(#inserts)*
                            },
                        )
                    }
                }

                #op_impl
            }
        } else {
            // No request body: no --file, no prepared-body state.
            quote! {
                #(#doc_attrs)*
                #[derive(Debug, Clone, ::clap::Parser)]
                pub struct #struct_ident {
                    #[arg(long, value_enum, default_value_t = #version_enum::#variant)]
                    version: #version_enum,
                    #(#fields)*
                }

                impl #struct_ident {
                    /// Run the operation through the atlas client and print
                    /// the JSON response; exit non-zero on any error.
                    pub async fn execute(self) -> #ret {
                        ::tracing::debug!(args = ?self, "executing atlas operation");
                        let client = ::mongodb_atlas_cli::atlas::client::AtlasClient::from_defaults()
                            .map_err(|error| {
                                eprintln!("{error}");
                                ::std::process::ExitCode::FAILURE
                            })?;
                        let response = client.execute(self).await.map_err(|error| {
                            eprintln!("{error}");
                            ::std::process::ExitCode::FAILURE
                        })?;
                        println!(
                            "{}",
                            ::serde_json::to_string_pretty(&response).map_err(|error| {
                                eprintln!("{error}");
                                ::std::process::ExitCode::FAILURE
                            })?
                        );
                        Ok(())
                    }
                }

                #op_impl
            }
        }
    });

    quote! {
        #probe_struct
        #version_enum_tok
        #probe_impl
        #(#version_structs)*
    }
}

/// `impl Operation` for one version struct: fixed method + URL template with
/// the struct's flags substituted in + this version's api version. Versions
/// with a request body report the body `execute()` prepared into the skipped
/// `body` field.
fn operation_impl(operation: &GeneratedOperation, version: &GeneratedVersion) -> TokenStream {
    let struct_ident = ident(&version.struct_ident);
    let method = method_value(&operation.method);
    let url = url_method(operation);
    let api_version = version_value(&version.api_version);
    let request_body = if version.body_schema.is_some() {
        quote! {
            fn request_body(&self) -> ::bytes::Bytes {
                self.body.clone().unwrap_or_default()
            }
        }
    } else {
        TokenStream::new()
    };
    quote! {
        impl ::mongodb_atlas_cli::atlas::Operation for #struct_ident {
            type Response = ::serde_json::Value;

            fn method(&self) -> ::http::Method {
                #method
            }

            fn url(&self) -> String {
                #url
            }

            fn version(&self) -> ::mongodb_atlas_cli::atlas::Version {
                #api_version
            }

            #request_body

            fn parse_response(
                bytes: ::bytes::Bytes,
            ) -> ::std::result::Result<Self::Response, ::mongodb_atlas_cli::atlas::OperationError>
            {
                ::serde_json::from_slice(&bytes).map_err(::std::convert::Into::into)
            }
        }
    }
}

fn method_value(method: &str) -> TokenStream {
    match method {
        "GET" => quote! { ::http::Method::GET },
        "POST" => quote! { ::http::Method::POST },
        "PUT" => quote! { ::http::Method::PUT },
        "DELETE" => quote! { ::http::Method::DELETE },
        "PATCH" => quote! { ::http::Method::PATCH },
        "HEAD" => quote! { ::http::Method::HEAD },
        _ => quote! { ::http::Method::from_bytes(#method).expect("configured HTTP method") },
    }
}

fn version_value(api_version: &ApiVersion) -> TokenStream {
    let version = match api_version {
        ApiVersion::Stable(year, month, day) => quote! {
            ::mongodb_atlas_cli::atlas::Version::date(#year as u16, #month as u8, #day as u8)
        },
        ApiVersion::Upcoming(year, month, day) => quote! {
            ::mongodb_atlas_cli::atlas::Version::upcoming(#year as u16, #month as u8, #day as u8)
        },
        ApiVersion::Preview => quote! { ::mongodb_atlas_cli::atlas::Version::preview() },
    };
    quote! {{
        let version = #version;
        ::tracing::debug!(%version, "atlas request version");
        version
    }}
}

/// Build the request URL: substitute `{pathParam}` placeholders from the
/// template, then append `?name=value` for every set query parameter.
fn url_method(operation: &GeneratedOperation) -> TokenStream {
    let method = &operation.method;
    let template = &operation.url_template;
    let replaces = operation.flags.iter().filter(|f| f.location == FlagLocation::Path).map(|f| {
        // The template holds `{parameterName}`; replace the whole placeholder
        // so no empty curly braces reach the wire.
        let placeholder = format!("{{{}}}", f.name);
        let field = ident(&f.ident);
        quote! { url = url.replace(#placeholder, &self.#field); }
    });
    let query_pushes = operation
        .flags
        .iter()
        .filter(|f| f.location == FlagLocation::Query)
        .map(|f| {
            let name = &f.name;
            let field = ident(&f.ident);
            if f.list {
                quote! {
                    for value in &self.#field {
                        params.push(::std::format!("{}={}", #name, value));
                    }
                }
            } else {
                quote! {
                    if let Some(value) = &self.#field {
                        params.push(::std::format!("{}={}", #name, value));
                    }
                }
            }
        });
    quote! {{
        let mut url = #template.to_owned();
        #(#replaces)*
        let mut params = ::std::vec::Vec::new();
        #(#query_pushes)*
        if !params.is_empty() {
            url.push('?');
            url.push_str(&params.join("&"));
        }
        ::tracing::debug!(method = #method, url = %url, "atlas request");
        url
    }}
}

fn dispatch_arms(operation: &GeneratedOperation, version_enum: &Ident) -> Vec<TokenStream> {
    let name = cmd_name(operation);
    operation
        .versions
        .iter()
        .map(|version| {
            let variant = ident(&version.variant_ident);
            let struct_ident = ident(&version.struct_ident);
            quote! {
                #version_enum::#variant => {
                    let args = ::std::iter::once(#name.to_owned()).chain(self.rest);
                    let cmd = #struct_ident::parse_from(args);
                    cmd.execute().await
                }
            }
        })
        .collect()
}

fn cmd_name(operation: &GeneratedOperation) -> String {
    operation.variant_name.to_lowercase()
}

/// `CompliancePolicy` and `Backup Snapshots` -> `compliance-policy`,
/// `backup-snapshots`. Mirrors clap's subcommand naming (`heck::ToKebabCase`),
/// so section names and usage paths match the real subcommand names exactly
/// (e.g. `MongoDBEmployeeAccess` -> `mongo-db-employee-access`).
fn kebab_case(name: &str) -> String {
    use heck::ToKebabCase;
    let out = name.to_kebab_case();
    if out.is_empty() {
        "_".to_owned()
    } else {
        out
    }
}

fn flag_field(flag: &GeneratedFlag) -> TokenStream {
    let field_ident = ident(&flag.ident);
    let description = flag.description.trim();
    let doc_attr = if description.is_empty() {
        TokenStream::new()
    } else {
        quote! { #[doc = #description] }
    };

    // Body flags are always optional with a value type matching the schema
    // leaf, so "flag set" means "user intent" (the `--file` XOR check) and
    // flag-built JSON carries the right scalar types. Required-ness is
    // enforced by the request-body schema at `prepare_body` time.
    if flag.location == FlagLocation::Body {
        let kind_ty = match flag.value_kind {
            FlagValueKind::String => quote! { String },
            FlagValueKind::Integer => quote! { i64 },
            FlagValueKind::Double => quote! { f64 },
            FlagValueKind::Boolean => quote! { bool },
        };
        // clap's derive only infers repeated args (Append action) from a bare
        // `Vec`, not a fully-qualified `::std::vec::Vec`.
        let field_ty = if flag.list {
            quote! { Option<Vec<#kind_ty>> }
        } else {
            quote! { Option<#kind_ty> }
        };
        // Boolean body flags toggle on (`--copy-protection-enabled`); clap's
        // default for Option<bool> is a Set action that demands a value.
        let arg_attr = if flag.value_kind == FlagValueKind::Boolean && !flag.list {
            quote! { #[arg(long, action = ::clap::ArgAction::SetTrue)] }
        } else {
            quote! { #[arg(long)] }
        };
        return quote! {
            #doc_attr
            #arg_attr
            #field_ident: #field_ty,
        };
    }

    let ty = match (flag.required, flag.list) {
        (true, true) => quote! { Vec<String> },
        (true, false) => quote! { String },
        (false, true) => quote! { Option<Vec<String>> },
        (false, false) => quote! { Option<String> },
    };
    quote! {
        #doc_attr
        #[arg(long)]
        #field_ident: #ty,
    }
}

/// Whether any body flag is set: `self.a.is_some() || self.b.is_some()`, or
/// just `false` when the version has no flat flags (an unsupported-shape body
/// that can only be supplied via `--file`/stdin).
///
/// Boolean flags are enable-only toggles (`SetTrue`, clap fills absent ones
/// with `Some(false)`), so only `Some(true)` counts as intent.
fn has_flags_expr(flags: &[GeneratedFlag]) -> TokenStream {
    if flags.is_empty() {
        return quote! { false };
    }
    let exprs = flags.iter().map(|flag| {
        let field = ident(&flag.ident);
        if flag.location == FlagLocation::Body && flag.value_kind == FlagValueKind::Boolean && !flag.list
        {
            quote! { self.#field == Some(true) }
        } else {
            quote! { self.#field.is_some() }
        }
    });
    quote! { #(#exprs)||* }
}

/// Insert one set body flag into the request-body object, keyed by its dotted
/// JSON path (`advancedConfiguration.minimumEnabledTlsProtocol` nests two
/// levels deep). Boolean toggles emit `true` only.
fn body_flag_insert(flag: &GeneratedFlag) -> TokenStream {
    let field = ident(&flag.ident);
    let name = &flag.name;
    if flag.location == FlagLocation::Body && flag.value_kind == FlagValueKind::Boolean && !flag.list
    {
        quote! {
            if self.#field == Some(true) {
                crate::__atlas_cli_body::insert(object, #name, ::serde_json::Value::Bool(true));
            }
        }
    } else {
        quote! {
            if let Some(value) = &self.#field {
                crate::__atlas_cli_body::insert(object, #name, ::serde_json::json!(value));
            }
        }
    }
}

/// The shared request-body builder emitted once per generated CLI, wired into
/// every `prepare_body`. Kept out of the per-version structs so the
/// read/validate/XOR logic exists in exactly one place.
fn body_module() -> TokenStream {
    quote! {
        /// Shared request-body construction for generated version structs:
        /// `--file`/stdin XOR flat body flags, validated against the version's
        /// JSON Schema before the request goes out.
        #[allow(dead_code, reason = "wired into generated prepare_body impls")]
        pub(crate) mod __atlas_cli_body {
            use ::bytes::Bytes;
            use ::serde_json::{Map, Value};

            /// Build the request body for one version struct.
            ///
            /// Exactly one of the file/stdin input or the flat body flags may
            /// be set; providing both is an error. When `schema` is given the
            /// body is validated against it first. Returns `None` for versions
            /// that carry no request body (never reached from generated code).
            pub(crate) fn build_body(
                file: Option<&mut ::clio::Input>,
                schema: Option<&str>,
                has_flags: bool,
                build: impl FnOnce(&mut Map<String, Value>),
            ) -> Result<Option<Bytes>, String> {
                if file.is_some() && has_flags {
                    return Err(
                        "pass either `--file` or request-body flags, not both".to_owned(),
                    );
                }
                if let Some(input) = file {
                    return read_body(input, schema);
                }
                if has_flags {
                    let mut object = Map::new();
                    build(&mut object);
                    let body = Value::Object(object);
                    validate(schema, &body)?;
                    return ::serde_json::to_vec(&body)
                        .map(Bytes::from)
                        .map(Some)
                        .map_err(|error| error.to_string());
                }
                if schema.is_some() {
                    return Err("a request body is required: pass `--file` (or `-` for stdin) or at least one request-body flag"
                        .to_owned());
                }
                Ok(None)
            }

            fn read_body(
                input: &mut ::clio::Input,
                schema: Option<&str>,
            ) -> Result<Option<Bytes>, String> {
                use ::std::io::Read;
                let mut text = String::new();
                input
                    .read_to_string(&mut text)
                    .map_err(|error| format!("cannot read request body: {error}"))?;
                let value: Value = ::serde_json::from_str(&text)
                    .map_err(|error| format!("request body is not valid JSON: {error}"))?;
                validate(schema, &value)?;
                Ok(Some(Bytes::from(text)))
            }

            fn validate(schema: Option<&str>, value: &Value) -> Result<(), String> {
                let Some(schema) = schema else {
                    return Ok(());
                };
                let schema: Value = ::serde_json::from_str(schema)
                    .map_err(|error| format!("invalid request body schema: {error}"))?;
                let validator = ::jsonschema::validator_for(&schema)
                    .map_err(|error| format!("invalid request body schema: {error}"))?;
                validator
                    .validate(value)
                    .map_err(|error| format!("request body is invalid: {error}"))?;
                Ok(())
            }

            /// Insert `value` at the nested `dotted` JSON path, creating
            /// intermediate objects: `insert(map, "person.first_name", ..)`
            /// sets `{"person": {"first_name": ..}}`.
            pub(crate) fn insert(object: &mut Map<String, Value>, dotted: &str, value: Value) {
                let mut parts = dotted.split('.');
                let Some(head) = parts.next() else {
                    return;
                };
                let Some(next) = parts.next() else {
                    object.insert(head.to_owned(), value);
                    return;
                };
                let tail = ::std::iter::once(next).chain(parts).collect::<Vec<_>>().join(".");
                let nested = object
                    .entry(head.to_owned())
                    .or_insert_with(|| Value::Object(Map::new()));
                let Value::Object(nested) = nested else {
                    return;
                };
                insert(nested, &tail, value);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{
        help_about, operation_sections, split_first_sentence, HelpEntry,
    };
    use crate::ir::{GeneratedComponent, GeneratedEntity, GeneratedOperation, OperationKind};

    #[test]
    fn splits_on_first_period_followed_by_space() {
        let (about, rest) = split_first_sentence(
            "Creates one cluster in the specified project. Cluster hosts keep the same data set.",
        );
        assert_eq!(about, "Creates one cluster in the specified project.");
        assert_eq!(rest.as_deref(), Some("Cluster hosts keep the same data set."));
    }

    #[test]
    fn keeps_single_sentence_whole() {
        let (about, rest) = split_first_sentence("Creates one cluster in the specified project.");
        assert_eq!(about, "Creates one cluster in the specified project.");
        assert!(rest.is_none());
    }

    #[test]
    fn does_not_split_on_abbreviations() {
        let (about, rest) = split_first_sentence("Uses the files.example.com host here. Done.");
        assert_eq!(about, "Uses the files.example.com host here.");
        assert_eq!(rest.as_deref(), Some("Done."));
    }

    #[test]
    fn no_period_splits_at_end() {
        let (about, rest) = split_first_sentence("Creates one cluster");
        assert_eq!(about, "Creates one cluster");
        assert!(rest.is_none());
    }

    fn op(name: &str, kind: OperationKind) -> GeneratedOperation {
        GeneratedOperation {
            variant_name: name.to_owned(),
            description: format!("{name} a thing."),
            probe_ident: String::new(),
            version_enum_ident: String::new(),
            versions: Vec::new(),
            flags: Vec::new(),
            kind,
            method: "GET".to_owned(),
            url_template: "/api/atlas/v2/ping".to_owned(),
        }
    }

    fn component(id: &str) -> GeneratedComponent {
        GeneratedComponent {
            id: id.to_owned(),
            ident: id.to_owned(),
            description: Some(format!("Manage {}.", id.to_lowercase())),
            operations: Vec::new(),
        }
    }

    #[test]
    fn help_about_strips_one_trailing_period() {
        assert_eq!(help_about("Creates one cluster."), "Creates one cluster");
        assert_eq!(help_about("Read and update your cluster's advanced configuration."), "Read and update your cluster's advanced configuration");
    }

    #[test]
    fn help_about_keeps_double_periods_and_plain_text() {
        assert_eq!(help_about("Wait..."), "Wait...");
        assert_eq!(help_about("No trailing dot"), "No trailing dot");
    }

    #[test]
    fn operation_sections_splits_crud_from_actions_sorted() {
        let ops = vec![
            op("Read", OperationKind::Crud),
            op("RestartPrimaries", OperationKind::Action),
            op("Create", OperationKind::Crud),
            op("Status", OperationKind::Action),
        ];
        assert_eq!(
            operation_sections(&ops),
            vec![
                ("Operations".to_owned(), vec![
                    HelpEntry { name: "create".to_owned(), about: "Create a thing".to_owned() },
                    HelpEntry { name: "read".to_owned(), about: "Read a thing".to_owned() },
                ]),
                ("Actions".to_owned(), vec![
                    HelpEntry { name: "restart-primaries".to_owned(), about: "RestartPrimaries a thing".to_owned() },
                    HelpEntry { name: "status".to_owned(), about: "Status a thing".to_owned() },
                ]),
            ]
        );
    }

    #[test]
    fn operation_sections_skips_empty_categories() {
        let only_actions = vec![op("Status", OperationKind::Action)];
        assert_eq!(operation_sections(&only_actions).len(), 1);
        assert_eq!(operation_sections(&only_actions)[0].0, "Actions");

        let only_crud = vec![op("Create", OperationKind::Crud)];
        assert_eq!(operation_sections(&only_crud)[0].0, "Operations");
    }

    #[test]
    fn entity_sections_appends_components_and_help() {
        let entity = GeneratedEntity {
            id: "Cluster".to_owned(),
            ident: "Cluster".to_owned(),
            description: Some("Manage clusters.".to_owned()),
            operations: vec![op("Create", OperationKind::Crud)],
            components: vec![component("AdvancedConfigurationOptions")],
        };
        let sections = super::entity_sections(&entity);
        let headings: Vec<&String> = sections.iter().map(|(h, _)| h).collect();
        assert_eq!(headings, vec!["Operations", "Components", "Help"]);
        let components = &sections[1].1;
        assert_eq!(
            components,
            &[HelpEntry { name: "advanced-configuration-options".to_owned(), about: "Manage advancedconfigurationoptions".to_owned() }]
        );
        assert_eq!(sections[2].1[0].name, "help");
    }

    #[test]
    fn group_with_no_components_skips_components_section() {
        let entity = GeneratedEntity {
            id: "CompliancePolicy".to_owned(),
            ident: "CompliancePolicy".to_owned(),
            description: None,
            operations: vec![op("Read", OperationKind::Crud), op("Disable", OperationKind::Action)],
            components: Vec::new(),
        };
        let sections = super::entity_sections(&entity);
        let headings: Vec<&String> = sections.iter().map(|(h, _)| h).collect();
        assert_eq!(headings, vec!["Operations", "Actions", "Help"]);
    }
}
