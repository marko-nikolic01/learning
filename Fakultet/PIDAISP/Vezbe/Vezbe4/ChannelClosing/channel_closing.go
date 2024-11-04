package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer close(c)
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for i := range c {
		fmt.Printf("Number: %d\n", i)
	}
}
