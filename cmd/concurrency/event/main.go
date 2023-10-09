package main

import (
	"fmt"

	"github.com/mccuskero/go-sandbox-general/pkg/event"
)

// This test creates 100 events and then go routines to process them
// using a channel.
// within the ProcessEvents func, a wait group is created to sync up finishing of worker threads.
// The calling main thread syncs on closing the channel.
func main() {
	fmt.Println("Starting event processing sandbox test")

	event.ProcessEvents(100)

	fmt.Println("Finished event processing sandbox test")
}
