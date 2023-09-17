package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/mccuskero/go-sandbox-general/pkg/slicing"
)

func doMemoryStats() {
	var ms runtime.MemStats

	fmt.Println("Running memstats program...")
	slicing.PrintMemStatus(ms)

	// create some memory
	intArr := make([]int, 90000)
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Int()
	}

	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		slicing.PrintMemStatus(ms)

		// clear is integrated with 1.21, as is maps
		//		if i==4 {
		//			fmt.Println("clearing intArr")
		//			clear(intArr)
		//		}
	}
}

func main() {
	fmt.Println("Starting slicing example.")

	slicing.SimpleSliceExample()
	slicing.SimpleMapExample()
	slicing.SimpleMapComparisonExample()
	slicing.CopyArrayByValue()
	doMemoryStats()
}
