use anyhow::{Context, Result};
use openapiv3::OpenAPI;
use tokio::fs::read_to_string;

pub use new_types::*;

use crate::hierarchy::{Hierarchy, cli_config};

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
    build_cli(&config);

    Ok(())
}

fn build_cli(hierarchy: &Hierarchy) {
    //let mut cli = command!();
    todo!()
}
