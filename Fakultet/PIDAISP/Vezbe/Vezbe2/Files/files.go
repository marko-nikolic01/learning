package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./Files/text.txt", os.O_RDONLY, os.ModeExclusive)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error while opening file...")
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		fmt.Printf("%d: "+scanner.Text()+"\n", i)
		i++
	}
}
