package concurrency

import (
	"fmt"
)

func DisplayMessageWithChannel(ch chan string) {
	fmt.Println("This is the message " + <-ch + "!")
}

func WriteToIntegerChannel(ch chan int, iterations int) {
	defer close(ch)
	for i := 0; i < iterations; i++ {
		ch <- i
	}
}

func ReadFromIntegerChannel(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

/*
// reading from a channnel returns data, and a boolean to check if closed.
data, ok := <-ch

	if !ok {
	    fmt.Println("channel closed")
	}
*/
func ReadFromUnbufferedIntChannel(ch chan int) {
	for val := range ch {
		fmt.Println(val, len(ch), cap(ch))
	}
}

func WriteToUnbufferedIntChannel(ch chan int, numIters int) {
	defer close(ch)
	for i := 0; i < numIters; i++ {
		ch <- i
	}
}
