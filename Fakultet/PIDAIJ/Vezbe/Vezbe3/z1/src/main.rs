use std::io;

fn main() {
    let mut sentence = String::new();

    println!("Enter sentence:");
    io::stdin().read_line(&mut sentence).expect("input error");

    let sentence = sentence.trim();

    let mut words: Vec<&str> = Vec::new();

    let mut start = 0;

    for (i, c) in sentence.char_indices() {
        if c == ' ' {
            if start < i {
                words.push(&sentence[start..i]);
            }
            start = i + 1;
        }
    }

    if start < sentence.len() {
        words.push(&sentence[start..]);
    }

    print!("Words: [ ");
    for word in words {
        print!("{} ", word);
    }
    print!("]\n");
}
