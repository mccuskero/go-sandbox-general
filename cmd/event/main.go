package main

import (
	"fmt"

	"github.com/mccuskero/go-sandbox-general/pkg/event"
)

func main() {
	fmt.Println("Starting event processing sandbox test")

	event.ProcessEvents(100)

	fmt.Println("Finished event processing sandbox test")
}
