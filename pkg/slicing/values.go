package slicing

import (
	"fmt"
)

func CopyArrayByValue() {
	arr1 := [3]int{1,2,3}
	arr2 := arr1

	arr2[0] = 0

	fmt.Println(arr1)
	fmt.Println(arr2)
}