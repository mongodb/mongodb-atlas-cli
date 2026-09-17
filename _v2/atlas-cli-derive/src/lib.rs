//! `#[derive(atlas_cli)]` for the Atlas CLI.
//!
//! Dumb entry point: parse `#[atlas_cli(spec = "...", hierarchy = "...")]`,
//! read both files, hand them to `atlas-cli-codegen`, and append the returned
//! token stream to the derived item. All real logic lives in that crate.
#![cfg_attr(not(test), deny(clippy::unwrap_used, clippy::expect_used))]

use std::{collections::BTreeSet, env, fs, path::PathBuf};

use proc_macro::TokenStream;
use proc_macro2::TokenStream as TokenStream2;
use syn::{
    parse_macro_input,
    punctuated::Punctuated,
    token::Comma,
    Data, DeriveInput, Expr, ExprLit, Lit, Meta,
};

#[proc_macro_derive(atlas_cli, attributes(atlas_cli))]
pub fn derive_atlas_cli(input: TokenStream) -> TokenStream {
    derive_atlas_cli_inner(parse_macro_input!(input as DeriveInput)).into()
}

fn derive_atlas_cli_inner(input: DeriveInput) -> TokenStream2 {
    build(&input).unwrap_or_else(|err| err.to_compile_error())
}

/// Attribute config of `#[atlas_cli(...)]` on the root `Cli` struct.
struct Config {
    spec: String,
    hierarchy: String,
    excluded_operation_ids: BTreeSet<String>,
}

fn build(input: &DeriveInput) -> syn::Result<TokenStream2> {
    if !matches!(input.data, Data::Struct(_)) {
        return Err(syn::Error::new_spanned(
            &input.ident,
            "#[derive(atlas_cli)] expects a struct (the CLI root)",
        ));
    }

    let config = parse_config(input)?;

    let manifest_dir = env::var("CARGO_MANIFEST_DIR")
        .map_err(|_| syn::Error::new_spanned(&input.ident, "CARGO_MANIFEST_DIR is not set"))?;

    let read = |path: &str, what: &str| {
        let full = PathBuf::from(&manifest_dir).join(path);
        fs::read_to_string(&full).map_err(|error| {
            syn::Error::new_spanned(
                &input.ident,
                format!("cannot read {what} `{}`: {error}", full.display()),
            )
        })
    };
    let spec_yaml = read(&config.spec, "spec")?;
    let hierarchy_yaml = read(&config.hierarchy, "hierarchy")?;

    let cli = atlas_cli_codegen::from_config(
        &spec_yaml,
        &hierarchy_yaml,
        &atlas_cli_codegen::CodegenOptions {
            excluded_operation_ids: config.excluded_operation_ids,
        },
    )
    .map_err(|error| syn::Error::new_spanned(&input.ident, format!("{error}")))?;

    // The compiler keeps the item the derive applies to and appends this
    // output; clap expands its own derive on that same preserved item. So we
    // emit only the generated tree here, never a copy of `struct Cli`.
    let generated = atlas_cli_codegen::to_tokens(&cli);
    Ok(generated)
}

fn parse_config(input: &DeriveInput) -> syn::Result<Config> {
    let attribute = input
        .attrs
        .iter()
        .find(|attr| attr.path().is_ident("atlas_cli"))
        .ok_or_else(|| {
            syn::Error::new_spanned(
                &input.ident,
                "#[atlas_cli(spec = \"...\")] attribute is required",
            )
        })?;

    let metas = attribute
        .parse_args_with(Punctuated::<Meta, Comma>::parse_terminated)?;

    let mut spec = None;
    let mut hierarchy = None;
    let mut excluded_operation_ids = BTreeSet::new();

    for meta in metas {
        match meta {
            Meta::NameValue(name_value) if name_value.path.is_ident("spec") => {
                spec = Some(string_literal(&name_value.value, "spec")?);
            }
            Meta::NameValue(name_value) if name_value.path.is_ident("hierarchy") => {
                hierarchy = Some(string_literal(&name_value.value, "hierarchy")?);
            }
            Meta::NameValue(name_value) if name_value.path.is_ident("excluded_operation_ids") => {
                excluded_operation_ids = string_array(&name_value.value)?;
            }
            other => {
                return Err(syn::Error::new_spanned(
                    other,
                    "expected one of `spec`, `hierarchy`, `excluded_operation_ids`",
                ));
            }
        }
    }

    Ok(Config {
        spec: spec.ok_or_else(|| {
            syn::Error::new_spanned(&input.ident, "`spec = \"...\"` is required")
        })?,
        hierarchy: hierarchy.ok_or_else(|| {
            syn::Error::new_spanned(&input.ident, "`hierarchy = \"...\"` is required")
        })?,
        excluded_operation_ids,
    })
}

fn string_literal(expr: &Expr, what: &str) -> syn::Result<String> {
    match expr {
        Expr::Lit(ExprLit {
            lit: Lit::Str(lit), ..
        }) => Ok(lit.value()),
        _ => Err(syn::Error::new_spanned(
            expr,
            format!("`{what}` must be a string literal"),
        )),
    }
}

/// A `["a", "b"]` attribute value.
fn string_array(expr: &Expr) -> syn::Result<BTreeSet<String>> {
    let Expr::Array(array) = expr else {
        return Err(syn::Error::new_spanned(
            expr,
            "expected a string array like `[\"a\", \"b\"]`",
        ));
    };
    let mut values = BTreeSet::new();
    for element in &array.elems {
        values.insert(string_literal(element, "array element")?);
    }
    Ok(values)
}
