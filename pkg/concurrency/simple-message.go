package concurrency

import (
	"fmt"
	"runtime"
	"time"
)

func DisplayMessage(msg string) {
	fmt.Println(msg)
}

func SimpleDisplayMessage() {
	fmt.Println("Hello")

	fmt.Println("This is the concurrency sandbox.")
	fmt.Printf("The number logical CPU for concurreny is: %d\n", runtime.NumCPU())
	go DisplayMessage("Hi")
	go DisplayMessage("Bye")
	time.Sleep(2 * time.Second)
}

/*
func SimpleDisplayMessageWithChannel(chan ) {
	fmt.Println("Hello")

	fmt.Println("This is the concurrency sandbox.")
	fmt.Printf("The number logical CPU for concurreny is: %d\n", runtime.NumCPU())
	go DisplayMessage("Hi")
	go DisplayMessage("Bye")
	time.Sleep(2 * time.Second)
}
*/
