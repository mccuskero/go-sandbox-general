package geometry

import (
	"fmt"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())

	fmt.Println(&g)
}
