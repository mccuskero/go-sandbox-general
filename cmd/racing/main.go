package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	counter := 0
	// To play, add/remove the mutex...
	//	var mu sync.Mutex

	const num = 15
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func() {
			// To play, add/remove the mutex...
			// mu.Lock()
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			// mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
