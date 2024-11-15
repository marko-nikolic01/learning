package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(10000)

	for i := 1; i < 10001; i++ {
		go func() {
			defer waitGroup.Done()
			fmt.Printf("Go routine %d sleeping...\n", i)
			time.Sleep(1)
		}()
	}

	waitGroup.Wait()

	fmt.Println("All Go routines completed")
} 
