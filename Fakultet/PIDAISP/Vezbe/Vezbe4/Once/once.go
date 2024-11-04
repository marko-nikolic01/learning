package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	increment := func() {
		count++
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(100)

	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			defer waitGroup.Done()

			once.Do(increment)
		}()
	}

	waitGroup.Wait()

	fmt.Printf("Final count: %d\n", count)
}
