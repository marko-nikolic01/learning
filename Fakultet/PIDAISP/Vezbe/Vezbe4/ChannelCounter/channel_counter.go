package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mtx   sync.Mutex
	count int
}

func main() {
	counter := Counter{}

	c := make(chan bool)

	for i := 0; i < 4; i++ {
		go counter.increment(c)
	}

	<-c

	fmt.Printf("Final count: %d\n", counter.count)
}

func (counter *Counter) increment(c chan bool) {
	for {
		counter.mtx.Lock()

		if counter.count > 99 {
			counter.mtx.Unlock()
			c <- true
			return
		}

		counter.count++
		counter.mtx.Unlock()
	}
}
