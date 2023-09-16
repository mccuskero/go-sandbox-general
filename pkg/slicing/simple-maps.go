package slicing

import (
	"fmt"
	//	"maps"

	"golang.org/x/exp/maps"
)

func SimpleMapExample() {
	mp := make(map[string]int)

	mp["k1"] = 7
	mp["k2"] = 12

	fmt.Println("map: ", mp)
}

func SimpleMapComparisonExample() {
	mp := make(map[string]int)
	mp2 := make(map[string]int)

	mp["k1"] = 1
	mp2["k1"] = 1

	if maps.Equal(mp, mp2) {
		fmt.Println("Maps are equal.")
	}

}
