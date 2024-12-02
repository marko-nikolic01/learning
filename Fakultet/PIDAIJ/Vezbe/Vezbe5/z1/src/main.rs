struct Pair<'a, 'b> {
    ref1: &'a str,
    ref2: &'b i32
}

impl<'a, 'b> Pair<'a, 'b> {
    fn get_first(&self) -> &'a str {
        return self.ref1;
    }

    fn get_second(&self) -> &'b i32 {
        return self.ref2;
    }
}

fn main() {
    let pair = Pair {
        ref1: "REF1",
        ref2: &2
    };

    println!("get_first(): {}", pair.get_first());
    println!("get_second(): {}", pair.get_second());
}
