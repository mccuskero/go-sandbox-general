package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Test struct {
	value int
}

type Tests []Test

func (test *Test) show(wg *sync.WaitGroup) {
	fmt.Println(test.value)
	wg.Done()
}

func main() {
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	numTests := 10

	var tests Tests

	for i := 0; i < numTests; i++ {
		test := Test{
			value: i,
		}
		tests = append(tests, test)
	}

	var wg sync.WaitGroup
	wg.Add(numTests)

	for _, test := range tests {
		go func(t Test) {
			t.show(&wg)
		}(test)
	}
	wg.Wait()
}
