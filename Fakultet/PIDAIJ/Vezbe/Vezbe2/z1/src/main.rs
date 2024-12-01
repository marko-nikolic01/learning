use std::io;

fn main() {
    f_to_c();
}

fn f_to_c() {
    println!("Insert temperature in Farenheits (F):");
    
    let mut input : String = String::new();

    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line.");

    match input.trim().parse::<f64>() {
        Ok(f) => {
            let c = (f - 32.0) * 5.0 / 9.0;
            let c = (c * 100.0).round() / 100.0;
            println!("{:.2}°F is {:.2}°C", f, c);
        },
        Err(_) => {
            println!("Invalid input. Please enter a valid floating-point number.");
        }
    }
}
