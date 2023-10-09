package main

import (
	"fmt"
)

func checkForSubString(subString string, fullString string) bool {
	var subStrAlpha [26]int
	for _, c := range subString {
		fmt.Printf("%d-%d=%d\n", c, 'a', c-'a')
		subStrAlpha[c-'a']++
	}

//	fmt.Println(s1alpha)

	for i := range fullString {
		j := i
		var n int
		var s2window [26]int

		for j < len(fullString) && n < len(subString) {
			s2window[fullString[j]-'a']++
			j++
			n++
		}


		fmt.Println("subStrAlpha: ", subStrAlpha)
		fmt.Println("s2window:    ", s2window)

		if subStrAlpha == s2window {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(checkForSubString("aa", "aacatdogmouse"))
}
