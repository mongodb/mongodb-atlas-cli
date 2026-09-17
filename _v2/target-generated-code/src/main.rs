use std::process::ExitCode;

use clap::Parser;
use target_generated_code::*;

#[tokio::main]
async fn main() -> Result<(), ExitCode> {
    let command = Cli::parse();

    command.execute().await
}
