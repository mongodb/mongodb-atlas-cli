use clap::Parser;
use target_generated_code::*;

fn main() {
    let command = Cli::parse();

    println!("{command:#?}");
}
