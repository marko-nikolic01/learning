package main

import "fmt"

func main() {
	fmt.Println("Input range start:")

	var start int
	fmt.Scanf("%d", &start)

	fmt.Println("Input range end:")

	var end int
	fmt.Scanf("%d", &end)

	for end < start {
		fmt.Println("Range end can't be smaller than range start. Input range end:")

		fmt.Scanf("%d", &end)
	}

	fmt.Println("Range:")
	for i := start; i <= end; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
}
