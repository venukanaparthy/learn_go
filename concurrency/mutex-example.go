package concurrency

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

var accessCount int

func mutexExample() {

	var wg sync.WaitGroup
	accessCount = 0
	loop := 500
	wg.Add(loop)

	for i := 1; i <= loop; i++ {
		go func(c int) {

			defer wg.Done()
			incr()
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count = ", accessCount)
}

func incr() {
	mu.Lock()
	fmt.Println(&mu)
	defer mu.Unlock()
	accessCount = accessCount + 1
}
