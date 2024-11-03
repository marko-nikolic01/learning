package main

import "fmt"

func main() {
	numbers := []int{1, 241, 5, 543, 63412, 3, 25, 345, 0, 9325, 9, -275, 35, 53, 1, 34, 656, 325, 17, -1}

	routineNumber := 6

	minimum := findMinimum(numbers, routineNumber)

	fmt.Printf("MINIMUM: %d\n", minimum)
}

func findMinimum(numbers []int, routineNumber int) int {
	c := make(chan int)

	step := len(numbers) / routineNumber
	start := 0
	for i := 1; i < routineNumber; i++ {
		go findSubMinimum(numbers[start:start+step], c, i)
		start += step
	}
	go findSubMinimum(numbers[start:], c, routineNumber)

	minimum := <-c
	for i := 1; i < routineNumber; i++ {
		newMinimum := <-c
		if newMinimum < minimum {
			minimum = newMinimum
		}
	}

	return minimum
}

func findSubMinimum(numbers []int, c chan int, partNumber int) {
	minimum := numbers[0]

	n := len(numbers)
	for i := 0; i < n; i++ {
		if numbers[i] < minimum {
			minimum = numbers[i]
		}
	}

	fmt.Printf("SUB-MINIMUM %d: %d\n", partNumber, minimum)

	c <- minimum
}
