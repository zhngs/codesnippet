use std::env;
use std::process;

use minigrep::Config;

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);

    let config = Config::new(&args).unwrap_or_else(|err| {
        println!("parse config err: {}", err);
        process::exit(1);
    });

    println!("Searching for <{}> in file <{}>", config.query, config.filename);

    if let Err(e) = minigrep::run(config) {
        println!("app error: {}", e);
        process::exit(1);
    }
}