package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"

	"github.com/mccuskero/sandbox/pkg/concurrency"
)

func doSimpleMessage() {
	go concurrency.SimpleDisplayMessage()
	time.Sleep(2 * time.Second)
}

func doSimpleChannelMessage() {
	// make the channel, then startup the go routine that reads it.
	ch := make(chan string)
	defer close(ch)

	go concurrency.DisplayMessageWithChannel(ch)
	// send a string to the channel. NOTE, if you do this before the go routine start you will have
	// and deadlock, in that no go routine is ready to read the channel
	ch <- "This is my message"
	fmt.Println("Sent message to display.")
	// time.Sleep(2 * time.Second)
}

func doLongRunningWriteToIntegerChannel() {
	ch := make(chan int)

	go concurrency.WriteToIntegerChannel(ch, 20)
	for i := range ch {
		fmt.Println(i)
	}
}

func doLongRunningReadWriteToIntegerChannel() {
	ch := make(chan int)
	numTimes := 20

	go concurrency.WriteToIntegerChannel(ch, numTimes)
	go concurrency.ReadFromIntegerChannel(ch)

	time.Sleep(2 * time.Second)
}

func doReadWriteUnbufferedChannel() {
	ch := make(chan int, 3)

	go concurrency.WriteToUnbufferedIntChannel(ch, 5)
	concurrency.ReadFromUnbufferedIntChannel(ch)

	// if you thread this off, main will exit before it prints...
	//	go concurrency.ReadFromUnbufferedIntChannel(ch)
	//	time.Sleep(2 * time.Second)
}

func doReadOnlyWriteOnlyChannel() {
	ch := make(chan int, 3)

	go concurrency.SendWriteOnlyChannel(ch, 10)
	go concurrency.RecvReadOnlyChannel(ch)

	time.Sleep(2 * time.Second)
}

func doReadOnlyWriteOnlySelectChannel() {
	ch := make(chan int, 3)
	exit := make(chan bool)

	go concurrency.SendWriteOnlyChannel(ch, 1000)

	go func() {
		time.Sleep(time.Microsecond * 200)
		exit <- true
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-exit:
			return
		}
	}
}

func main() {
	fmt.Println("Starting up the concurrency sandbox.")
	fmt.Printf("The number logical CPU for concurreny is: %d\n", runtime.NumCPU())
	fmt.Printf("Main:System Process Id:%d\n", syscall.Getppid())

	doSimpleMessage()
	doSimpleChannelMessage()
	doLongRunningWriteToIntegerChannel()
	doLongRunningReadWriteToIntegerChannel()
	doReadWriteUnbufferedChannel()
	doReadOnlyWriteOnlyChannel()
	doReadOnlyWriteOnlySelectChannel()
}
