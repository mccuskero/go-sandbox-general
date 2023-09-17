package event

import (
	"time"

	"github.com/pborman/uuid"

	"github.com/mccuskero/go-sandbox-general/pkg/randomgen"
)

type Event struct {
	Uuid 		string
	Payload		string
	Timestamp 	int64
}

func CreateEvent() *Event {
	event := &Event{
		Uuid: 		uuid.NewRandom().String(),
		Payload:	randomgen.AlphaNumericString(20),
		Timestamp:  time.Now().Unix(),
	}

	return event
}

func CreateEvents(numEvents int) []*Event {

	events := make([]*Event, numEvents)

	for i := range events {
		events[i] = CreateEvent()
	}

	return events
}