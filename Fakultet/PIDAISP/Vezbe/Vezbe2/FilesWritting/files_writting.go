package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./FilesWritting/text.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	defer file.Close()

	if err != nil {
		fmt.Println("Error while opening file...")
	}

	writer := bufio.NewWriter(file)

	length, err := writer.WriteString("\nFile opened")

	if err != nil {
		fmt.Println("Error while writting in file...")
	}
	err = writer.Flush()

	fmt.Printf("Successfuly wrote %d characters in file.", length)
}
