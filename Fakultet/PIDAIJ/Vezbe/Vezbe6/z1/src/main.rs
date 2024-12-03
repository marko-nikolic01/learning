struct ListNode {
    val: i32,
    next: Option<Box<ListNode>>
}

struct LinkedList {
    head: Option<Box<ListNode>>
}

impl LinkedList {
    fn new() -> Self {
        return LinkedList{
            head: None
        }
    }

    fn is_empty(&self) -> bool {
        return self.head.is_none();
    }

    fn insert(&mut self, val: i32, index: i32) {
        if index == 0 {
            self.head = Some(Box::new(ListNode{
                val: val,
                next: self.head.take()
            }));
            return;
        }

        let mut current = match self.head.as_mut() {
            Some(node) => node,
            None => {
                self.head = Some(Box::new(ListNode{
                    val: val,
                    next: self.head.take()
                }));
                return;
            }
        };

        for i in 1..index {
            current = match current.next.as_mut() {
                Some(node) => node,
                None => {
                    panic!("Index {} out of bounds!", i);
                }
            };
        }

        current.next = Some(Box::new(ListNode{
            val: val,
            next: current.next.take()
        }));
    }

    fn print(&mut self) {
        let mut current = self.head.as_mut();

        print!("Linked list: ");
        loop {
            match current {
                Some(node) => {
                    print!("{} -> ", node.val);
                    current = node.next.as_mut();
                },
                None => break
            };
        }
        print!("None\n");
    }
}


fn main() {
    let mut linked_list = LinkedList::new();

    println!("Is empty: {}", linked_list.is_empty());

    linked_list.insert(1, 0);
    linked_list.insert(2, 1);
    linked_list.insert(3, 2);
    linked_list.insert(4, 3);
    linked_list.insert(5, 4);
    linked_list.insert(6, 5);

    linked_list.print();

    println!("Is empty: {}", linked_list.is_empty());
}
