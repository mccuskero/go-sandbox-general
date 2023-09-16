package concurrency

import (
	"fmt"
)

func SendWriteOnlyChannel(ch chan<- int, numVals int) {
	for i := 0; i < numVals; i++ {
		ch <- i
	}
}

func RecvReadOnlyChannel(ch <-chan int) {
	for {
		val, ok := <-ch

		if !ok {
			fmt.Println("Channel has closed")
			return
		}
		fmt.Println(val)
	}
}

// to the exit channel
func Exit(ch chan<- bool) {
	ch <- true
}
