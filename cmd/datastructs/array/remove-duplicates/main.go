package main

import (
	"fmt"
	"strings"
)


func getPrefix(str string) string {
	idx := strings.IndexByte(str,'/')
	if idx < 0 {
		return str
	}

	return str[:idx]
}

func removeDuplicatesWithSamePrefix(strArray []string) []string {
	uniquePrefixes := make(map[string]bool)

	list := []string{}

	for _, str := range strArray {
		prefix := getPrefix(str)
		if _, value := uniquePrefixes[prefix]; !value {
			uniquePrefixes[prefix] = true
			list = append(list, str)
		}
	}

	return list
}

func main() {
	fmt.Println("Remove duplicates starting")

	strSliceValues := []string{"abc","abc/def/ghi","abc/def/ghi/jkl","abc/def/ghijkl/jkl","def"}

	fmt.Println(strSliceValues)

	for _, str := range strSliceValues {
		fmt.Println(getPrefix(str))
	}

	newList := removeDuplicatesWithSamePrefix(strSliceValues)
	fmt.Println(newList)
}