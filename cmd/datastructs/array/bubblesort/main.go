package main

import (
	"fmt"

	"math/rand"
)

func createArray(len int) *[]int {
	array := make([]int, len)

	for i := 0; i < len; i++ {
		array[i] = rand.Intn(len)
	}

	return &array
}

func bubbleSort(array []int) *[]int {

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return &array
}

func main() {
	fmt.Println("Starting bubble sort")
	array := createArray(50)
	fmt.Println(array)

	array = bubbleSort(*array)
	fmt.Println(array)
	fmt.Println("Finishing bubble sort")
}

