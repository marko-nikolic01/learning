package main

import "fmt"

func main() {
	fmt.Println("Input your number:")

	var num int
	fmt.Scanf("%d", &num)

	fmt.Printf("You inputed number: %d.\n", num)

	if num < 0 {
		fmt.Println("Your number is negative.")
	} else if num > 0 {
		fmt.Println("Your number is positive.")
	} else {
		fmt.Println("Your number is 0.")
	}
}
