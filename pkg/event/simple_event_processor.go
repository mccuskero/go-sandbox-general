package event

import (
	"errors"
	"fmt"
	"sync"
)

// handles events by printing them out
func handleEvent(event *Event) {
	fmt.Printf("%v\n", event)
}

// create a event processing loop using a go routine, signal to the caller when done
func EventLoop(eventChan <-chan *Event, wg *sync.WaitGroup, numEventsProcessed *int) {
	var run = true
	go func() {
		fmt.Println("Starting event loop")
		for run {
			event, ok := <-eventChan
			if ok {
				handleEvent(event)
				*numEventsProcessed = *numEventsProcessed + 1
				//					fmt.Printf("%d\n", *numEventsProcessed)
			} else {
				fmt.Println("channel is closed exiting")
				run = false
			}
		}

		// NOTE: the wg.Done() cannot be deferred, a race condition will kick in
		// add EventLoop will not process channel is closed
		wg.Done()
		fmt.Println("Finished event loop")
	}()
}

func ProcessEvents(numEvents int) error {
	numEventsProcessed := 0

	// create event Chan
	eventChan := make(chan *Event, numEvents)

	// create one wait group representing the one eventloop
	var wg sync.WaitGroup
	wg.Add(1)

	// startup event processing loop
	EventLoop(eventChan, &wg, &numEventsProcessed)

	// add events to loop
	// create events
	events := CreateEvents(numEvents)

	// send events to eventloop to consume
	for _, event := range events {
		eventChan <- event
	}

	// close channel, and wait
	close(eventChan)

	// check that all events have been processed
	wg.Wait()

	fmt.Printf("numEvernts: %d, numEventsProcessed: %d\n", numEvents, numEventsProcessed)
	if numEventsProcessed != numEvents {
		return errors.New("event processed count does not match total events")
	}

	return nil
}
