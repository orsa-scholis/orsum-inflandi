extern crate man;

use std::fs;
use man::prelude::*;

fn main() {
    generate_main_man_page();
    generate_start_man_page();
}

fn generate_start_man_page() {
    let main_man = Manual::new("cli/app")
        .about("Orsum-Inflandi's CLI")
        .arg(Arg::new("command"))
        .example(
            Example::new()
                .text("print help of command")
                .command("cli/app help")
                .output("<<Help page>>"),
        )
        .author(Author::new("Lukas Bischof").email("polinderis@gmail.com"))
        .author(Author::new("Philipp Fehr").email("philipp@thefehr.me"))
        .render();

    render_file("app.man", main_man);
}

fn generate_main_man_page() {
    let start_man = Manual::new("cli/app start")
        .about("Orsum-Inflandi's Start CLI Command")
        .flag(
            Flag::new()
                .long("--backend")
                .help("Only start backend")
        )
        .flag(
            Flag::new()
                .long("--frontend")
                .help("Only start frontend")
        )
        .flag(
            Flag::new()
                .long("--dev")
                .help("Enables HMR for backend")
        )
        .flag(
            Flag::new()
                .long("--dual-frontend")
                .help("Starts two frontends for debugging and testing")
        )
        .author(Author::new("Lukas Bischof").email("polinderis@gmail.com"))
        .author(Author::new("Philipp Fehr").email("philipp@thefehr.me"))
        .render();
    render_file("start.man", start_man);
}


fn render_file(file_name: &str, content: String) {
    fs::write(format!("./{}", file_name), content).expect("Could not write");
}
