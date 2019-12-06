mod intcode;
use std::io::{stdin, stdout, Write};
// the main program will take one argument which is a file to load the program from
// stdin will be used for input and stdout for output
fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() < 1 {
        println!("missing program file name");
    }
    let filename = args[1].clone();
    println!("loading program from: {}", filename);
    // in go I would open the file and read through the contents with something
    // like scanf, but we know these programs will be small, so lets just load it
    // all into memory and split it on the commas.
    let content = std::fs::read_to_string(filename).expect("failed to load program");

    println!("reading code...");

    let mut code: Vec<i32> = Vec::new();
    for value in content.trim().split(",") {
        match value.parse::<i32>() {
            Ok(v) => code.push(v),
            _ => panic!("syntax error in program"),
        }
    }
    println!("code loaded, building interpreter");

    // create the machine
    let mut interpreter = intcode::Computer::new(code);

    // attach the i/o handlers
    // for input, read from stdin
    let input_handler = || -> i32 {
        let mut input = String::new();
        print!("input> ");
        let _ = stdout().flush();
        stdin().read_line(&mut input).expect("failed to read input");
        match input.trim().parse::<i32>() {
            Ok(v) => v,
            _ => panic!("input not i32"),
        }
    };
    // for output print to stdout
    let output_handler = |x: i32| {
        println!("output> {}", x);
    };

    println!("executing program");
    // now run the program
    if !interpreter.run(input_handler, output_handler) {
        panic!("program execution failed")
    }
}
