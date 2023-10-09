package main

import (
    "fmt"
    "sort"
)

func mergeIntIntervals(intervals [][]int) [][]int {
    const start, end = 0, 1

    fmt.Println("starting merge...")    

    var merged [][]int

    // We need to sort the interval before we merge them
    if len(intervals) > 1 {
        sort.Slice(intervals, 
            func(i, j int) bool {
                return intervals[i][start] < intervals[j][start]
            })
    }

    // the intervals must be sorted for the following to work
    for _, interval := range intervals {
        last := len(merged) - 1
        // first time through... 
        if last < 0  {
            fmt.Println("if: last ", last)
            merged = append(merged,
                []int{start: interval[start], end: interval[end]},
            )
        // add interval to merged list, 
        } else if interval[start] > merged[last][end] {
            fmt.Println("else if: last: ", last, " interval[start]: ",  interval[start], ">", merged[last][end])
            merged = append(merged,
                []int{start: interval[start], end: interval[end]},
            )
        // extend the last interval added
        } else if interval[end] > merged[last][end] {
            fmt.Println(interval, " else if 2: last: ", last, " interval[end]", interval[end], ">", merged[last][end])
            merged[last][end] = interval[end]
        }

        fmt.Println(merged)
    }

    fmt.Println("finished merge.")    


    return merged[:len(merged):len(merged)]
}

func main() {
    // note that the test intervals, start is always less then, or equal to, end
    test_intervals := []struct {
        intervals [][]int
    }{
        {[][]int{{6, 6}, {1, 3}, {8, 10}, {11, 18}}},
        {[][]int{{3, 6}, {1, 3}, {8, 10}, {11, 18}}},
        {[][]int{{1, 2}}},
        {[][]int{}},
        {[][]int{{6, 10}, {3, 4}, {2, 4}}},
    }

    for _, tt := range test_intervals {
        fmt.Print(tt.intervals)
        fmt.Println(" ->", mergeIntIntervals(tt.intervals))
    }
}