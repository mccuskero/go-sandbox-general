package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 1111
	describe(i)

	i = "This is a test"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
