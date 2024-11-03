package main

import "fmt"

func main() {
	fmt.Println("Input first number:")

	var num1 int
	fmt.Scanf("%d", &num1)

	fmt.Println("Input second number:")

	var num2 int
	fmt.Scanf("%d", &num2)

	sum := add(num1, num2)
	fmt.Printf("Sum of numer %d and number %d is number %d.\n", num1, num2, sum)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
