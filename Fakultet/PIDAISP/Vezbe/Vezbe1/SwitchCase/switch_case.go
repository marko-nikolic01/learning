package main

import "fmt"

func main() {
	fmt.Println("Input your number:")

	var num int
	fmt.Scanf("%d", &num)

	fmt.Printf("You inputed number: %d.\n", num)

	switch num {
	case 69:
		fmt.Println("Nice.")
	default:
		fmt.Println("Not nice.")
	}
}
