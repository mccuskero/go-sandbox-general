package measure

//
// This packge provides a method to track the execution time of a function:
// example:
//        defer measure.Duration(measure.Track("myFunc"))
//

import (
	"fmt"
	"time"
) 

func Track(measureName string) (string, time.Time) {
	return measureName, time.Now()
}

func Duration(measureName string, beginTime time.Time) {
	fmt.Printf("%v:%v\n", measureName, time.Since(beginTime))
}

