package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mtx   sync.Mutex
	count int
}

func (counter *Counter) Increment() {
	counter.mtx.Lock()
	defer counter.mtx.Unlock()

	counter.count++

	fmt.Printf("New counter value: %d\n", counter.count)
}

func main() {
	counter := Counter{}

	for i := 0; i < 10_000; i++ {
		counter.Increment()
	}
}
