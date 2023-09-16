package slicing

import (
	"fmt"
	"runtime"
	"time"
)

func PrintMemStatus(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)

	fmt.Println("-------------------------------------------")
	fmt.Println("Memory Statistics Reporting: ", time.Now())
	fmt.Println("-------------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
}
