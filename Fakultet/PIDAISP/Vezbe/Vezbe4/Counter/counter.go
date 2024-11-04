package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mtx   sync.Mutex
	count int
}

func (counter *Counter) increment() {
	counter.count++
}

func main() {
	counter := Counter{}

	routineCount := 4

	var waitGroup sync.WaitGroup
	waitGroup.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		go incrementRoutine(&counter, &waitGroup)
	}

	waitGroup.Wait()

	fmt.Printf("Final count: %d\n", counter.count)
}

func incrementRoutine(counter *Counter, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for {
		counter.mtx.Lock()
		if counter.count >= 100 {
			counter.mtx.Unlock()
			break
		}
		counter.increment()
		counter.mtx.Unlock()
	}
}
