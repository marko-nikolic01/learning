package main

import "fmt"

func main() {
	n := 10

	c := make(chan int)
	done := make(chan bool)

	go fib(n, c)
	go printFib(c, done)

	<-done

	fmt.Println("Fibonacci finished")

	return
}

func fib(n int, c chan int) {
	defer close(c)

	f1 := 0
	f2 := 1

	for n > 0 {
		c <- f1

		fTemp := f2
		f2 += f1
		f1 = fTemp

		n--
	}
}

func printFib(c chan int, done chan bool) {
	i := 1

	for {
		select {
		case fibNum, ok := <-c:
			if !ok {
				done <- true
				return
			}
			fmt.Printf("Fibonacci %d: %d\n", i, fibNum)
			i++
		}
	}
}
