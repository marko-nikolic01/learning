package main

import "fmt"

func main() {
	var words = [2]string{"Car", "Fish"}
	fmt.Printf("Static array (length: %d):\n", len(words))
	for i, word := range words {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

	wordsSlice := make([]string, 0, 3)
	fmt.Printf("Dynamic array/slice (length: %d, cap: %d):\n", len(wordsSlice), cap(wordsSlice))
	for i, word := range wordsSlice {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

	wordsSlice = append(wordsSlice, "Apple")
	wordsSlice = append(wordsSlice, "Banana")
	fmt.Printf("Dynamic array/slice (length: %d, cap: %d):\n", len(wordsSlice), cap(wordsSlice))
	for i, word := range wordsSlice {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

	wordsSlice = append(wordsSlice, "Dog")
	wordsSlice = append(wordsSlice, "Cat")
	fmt.Printf("Dynamic array/slice (length: %d, cap: %d):\n", len(wordsSlice), cap(wordsSlice))
	for i, word := range wordsSlice {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

	wordsSlice = wordsSlice[1:]
	fmt.Printf("Dynamic array/slice (length: %d, cap: %d):\n", len(wordsSlice), cap(wordsSlice))
	for i, word := range wordsSlice {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

	wordsSlice = wordsSlice[0 : len(wordsSlice)-1]
	fmt.Printf("Dynamic array/slice (length: %d, cap: %d):\n", len(wordsSlice), cap(wordsSlice))
	for i, word := range wordsSlice {
		fmt.Printf("\t- %d. %s\n", i, word)
	}

}
