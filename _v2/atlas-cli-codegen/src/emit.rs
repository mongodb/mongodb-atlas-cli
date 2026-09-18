//! [`GeneratedCli`] -> `TokenStream`: dumb, positional translation.
//!
//! All logic worth testing lives in `from_spec.rs`; this module only maps the
//! intermediate representation onto clap types. The CLI tree is
//! `Cli -> group -> entity -> {direct ops, component -> ops}`; a single-entity
//! group is flattened onto the group command.

use proc_macro2::{Span, TokenStream};
use quote::quote;
use syn::Ident;

use crate::ir::*;

/// Fully-qualified return type shared by every `execute`.
fn execute_ret() -> TokenStream {
    quote! { ::std::result::Result<(), ::std::process::ExitCode> }
}

pub fn to_tokens(cli: &GeneratedCli) -> TokenStream {
    let cli_subcommands = cli_subcommands(cli);
    let cli_execute = cli_execute(cli);
    let groups = cli.groups.iter().map(group);

    let operations = cli
        .groups
        .iter()
        .flat_map(|g| g.entities.iter())
        .flat_map(|e| {
            e.operations
                .iter()
                .chain(e.components.iter().flat_map(|c| c.operations.iter()))
        })
        .map(operation);

    quote! {
        #cli_subcommands
        #cli_execute
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

    let group_struct = quote! {
        #[derive(Debug, ::clap::Args)]
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
            let entity_struct = quote! {
                #[derive(Debug, ::clap::Args)]
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
            let component_struct = quote! {
                #[derive(Debug, ::clap::Args)]
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
        quote! {
            #(#doc_attrs)*
            #[derive(Debug, ::clap::Parser)]
            pub struct #struct_ident {
                #[arg(long, value_enum, default_value_t = #version_enum::#variant)]
                version: #version_enum,
                #(#fields)*
            }

            impl #struct_ident {
                pub async fn execute(&self) -> #ret {
                    todo!()
                }
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

/// `CompliancePolicy` and `Backup Snapshots` -> `compliance-policy`, `backup-snapshots` (help text only).
fn kebab_case(name: &str) -> String {
    let mut raw = String::new();
    for (i, c) in name.chars().enumerate() {
        if c.is_ascii_alphanumeric() {
            if c.is_ascii_uppercase() && i > 0 {
                raw.push('-');
            }
            raw.push(c.to_ascii_lowercase());
        } else {
            raw.push('-');
        }
    }
    let mut out = String::new();
    let mut previous_dash = false;
    for c in raw.chars() {
        if c == '-' && previous_dash {
            continue;
        }
        previous_dash = c == '-';
        out.push(c);
    }
    out.trim_matches('-').to_owned()
}

fn flag_field(flag: &GeneratedFlag) -> TokenStream {
    let field_ident = ident(&flag.ident);
    let description = flag.description.trim();
    let doc_attr = if description.is_empty() {
        TokenStream::new()
    } else {
        quote! { #[doc = #description] }
    };
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

#[cfg(test)]
mod tests {
    use super::split_first_sentence;

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
}
