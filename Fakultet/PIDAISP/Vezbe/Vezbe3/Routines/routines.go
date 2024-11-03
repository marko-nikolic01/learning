package main

import "fmt"

func main() {
	for i := 1; i < 1001; i++ {
		go hello(i)
	}
}

func hello(number int) {
	fmt.Printf("Hello from %d!\n", number)
}
