package main

import (
	"fmt"

	"github.com/mccuskero/go-sandbox-general/pkg/geometry"
)

type rectangle struct {
	width float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
} 

func (r rectangle) Perimeter() float64 {
	return 2*r.height + 2*r.width
}

// if you look at the memory address, printed here, and 
// the memory address printed in the g.Measure, you will see 
// that the interface Method set is pass by value. 
// event if you de-reference the location. 
func main() {
	fmt.Println("Starting geometry sandbox to test interfaces")

	r := rectangle{10,10}

	fmt.Printf("%p",&r)

	geometry.Measure(&r)
}