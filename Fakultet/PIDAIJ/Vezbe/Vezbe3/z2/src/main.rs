fn main() {
    let mut vector = vec![1, 2, 3, 4, 5, 6, 7];

    println!("Vector: {:?}", vector);

    swap(&mut vector);

    println!("Vector - swapped: {:?}", vector);
}

fn swap(vector: &mut Vec<i32>) {
    let mut start = 0;
    let mut end = vector.len() - 1;

    while start < end {
        let temp = vector[start];
        vector[start] = vector[end];
        vector[end] = temp;

        start += 1;
        end -= 1;
    }
}
