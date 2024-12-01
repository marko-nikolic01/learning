use std::{collections::HashMap, io};
use rand::Rng;

fn main() {
    let vector = make_vector();
    write_vector(&vector);
    write_vector_reverse(&vector);
    write_vector_element(&vector);
    write_vector_every_index_divisible_by_3(&vector);
    count_even_odd_indexes(&vector);
}

fn make_vector() -> Vec<i32> {
    println!("Enter vector size:");

    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Input error.");

    let n: usize = match input.trim().parse() {
        Ok(number) => number,
        Err(_) => {
            println!("Array size must be a valid number.");
            return make_vector();
        }
    };

    let mut vector = Vec::with_capacity(n);

    let mut rng = rand::thread_rng();

    for _ in 0..n {
        vector.push(rng.gen_range(0..100));
    }

    return vector;
}

fn write_vector(vector: &Vec<i32>) {
    print!("Vector: [ ");
    for x in vector {
        print!("{} ", x);
    }
    print!("]\n");
}

fn write_vector_reverse(vector: &Vec<i32>) {
    print!("Vector - reversed: [ ");
    for x in vector.iter().rev() {
        print!("{} ", x);
    }
    print!("]\n");
}

fn write_vector_element(vector: &Vec<i32>) {
    println!("Enter index:");

    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Input error.");

    match input.trim().parse::<usize>() {
        Ok(i) => {
            match vector.get(i) {
                Some(element) => println!("Element at index {} is: {}", i, element),
                None => println!("Index {} is out of bounds", i)
            }
        },
        Err(_) => {
            println!("Index must be a valid number.");
            write_vector_element(&vector);
        }
    };
}

fn write_vector_every_index_divisible_by_3(vector: &Vec<i32>) {
    print!("Vector - indexes divisible by 3: [ ");
    for (i, &val) in vector.iter().enumerate() {
        if i % 3 == 0 {
            print!("{} ", val);
        }
    }
    print!("]\n");
}

fn count_even_odd_indexes(vector: &Vec<i32>) {
    let mut indexes: HashMap<&str, i32> = HashMap::new();
    indexes.insert("event", 0);
    indexes.insert("odd", 0);

    print!("Vector - indexes divisible by 3: [ ");
    for (i, _) in vector.iter().enumerate() {
        if i % 2 == 0 {
            let even = indexes.entry("even").or_insert(0);
            *even += 1;
        } else {
            let odd = indexes.entry("odd").or_insert(0);
            *odd += 1;
        }
    }

    let even = indexes.entry("even").or_insert(0);
    println!("Even indexes: {}", *even);

    let odd = indexes.entry("odd").or_insert(0);
    println!("Odd indexes: {}", *odd);
}
