package slicing

import (
	"fmt"
)

func SimpleSliceExample() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]

	fmt.Println("Array: ", array)
	fmt.Println("Slice: ", slice)
}
