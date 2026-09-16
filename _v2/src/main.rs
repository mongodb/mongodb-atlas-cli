use std::{borrow::Cow, ops::Deref};

use anyhow::{Context, Result, bail};
use clap::{Command, builder};
use openapiv3::OpenAPI;
use tokio::fs::read_to_string;

pub use new_types::*;

use crate::hierarchy::{Component, Entity, Hierarchy, cli_config};

mod hierarchy;
mod new_types;

const SPEC_PATH: &'static str = "/Users/jeroen.vervaeke/git/github.com/mongodb/mongodb-atlas-cli/tools/internal/specs/spec-with-overlays.yaml";

#[tokio::main]
async fn main() -> Result<()> {
    tracing_subscriber::fmt::init();

    let spec = read_to_string(SPEC_PATH).await.context("open spec file")?;
    let openapi: OpenAPI = serde_yaml::from_str(&spec).context("parse spec")?;
    let config = cli_config(&openapi).context("build cli config")?;

    println!("{config:#?}");
    let cli = build_cli(&config).context("build cli")?;
    let matches = cli.get_matches();

    Ok(())
}

fn build_cli(hierarchy: &Hierarchy) -> Result<Command> {
    let mut command = Command::new("atlas-cli");

    for (name, group) in hierarchy.groups() {
        match group.entities().count() {
            0 => {
                bail!("a group should have at least one entity '{name}' doesn't have a single one")
            }
            1 => {
                let (entity_name, entity) = group.entities().next().unwrap();
                let sub_command = build_entity_subcommand(entity_name, entity)
                    .context("build entity command")?
                    .name(name);

                command = command.subcommand(sub_command);
            }
            2.. => {
                todo!()
            }
        }
    }

    Ok(command)
}

fn build_entity_subcommand(name: &str, entity: &Entity) -> Result<Command> {
    let mut command = Command::new(name.to_string());

    // Crudl
    command = command.subcommand_help_heading("CRUDL");
    command = add_optinal_operation_id_sub_command(command, "create", entity.create.as_ref());
    command = add_optinal_operation_id_sub_command(command, "delete", entity.delete.as_ref());
    command = add_optinal_operation_id_sub_command(command, "get", entity.read.as_ref());
    command = add_optinal_operation_id_sub_command(command, "list", entity.list.as_ref());
    command = add_optinal_operation_id_sub_command(command, "update", entity.update.as_ref());

    // Actions
    if !entity.actions.is_empty() {
        command = command.subcommand(Command::new("".to_string()).about(""));
        command = command.subcommand(Command::new("Actions".to_string()).about(""));

        for (action, operation_id) in entity.actions.iter() {
            command = add_operation_id_sub_command(command, action, operation_id);
        }
    }

    // Components
    if !entity.components.is_empty() {
        command = command.subcommand(Command::new(" ".to_string()).about(""));
        command = command.subcommand(Command::new("Components".to_string()).about(""));

        for (component_name, component) in entity.components.iter() {
            command = add_component_sub_command(command, component_name, component);
        }
    }

    Ok(command)
}

fn add_optinal_operation_id_sub_command(
    command: Command,
    name: &str,
    operation_id: Option<impl AsRef<str>>,
) -> Command {
    let Some(operation_id) = operation_id else {
        return command;
    };

    let operation_id = operation_id.as_ref();

    command
        .subcommand(Command::new(name.to_string()).about(format!("this will call {operation_id}")))
}

fn add_operation_id_sub_command(
    command: Command,
    name: &str,
    operation_id: impl AsRef<str>,
) -> Command {
    add_optinal_operation_id_sub_command(command, name, Some(operation_id))
}

fn add_component_sub_command(command: Command, name: &str, component: &Component) -> Command {
    let mut sub_command = Command::new(name.to_string());

    sub_command = add_optinal_operation_id_sub_command(sub_command, "get", component.read.as_ref());
    sub_command =
        add_optinal_operation_id_sub_command(sub_command, "update", component.update.as_ref());

    for (action_name, operation_id) in component.actions.iter() {
        sub_command = add_operation_id_sub_command(sub_command, action_name, operation_id);
    }

    command.subcommand(sub_command)
}
