use std::io;

fn main() {
    println!("Input N:");

    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Input error.");

    match input.trim().parse::<i32>() {
        Ok(n) => {
            if n < 1 {
                println!("Input must be greater than 0.");
                main(); 
            } else {
                let f = fib(n);
                println!("{}. Fibonacci number is: {}.", n, f);
            }
        },
        Err(_) => {
            println!("Input must be a valid number.");
            main();
        }
    }
}

fn fib(mut n: i32) -> i32 {
    if n == 1 {
        return 0;
    } else if n == 2 {
        return 1;
    }

    let mut f1 = 0;
    let mut f2 = 1;

    n -= 2;
    while n > 0 {
        let f3 = f1 + f2;
        f1 = f2;
        f2 = f3;
        n -= 1;
    }

    return f2;
}