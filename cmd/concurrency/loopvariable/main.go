package main

import (  
    "fmt"
    "time"
)

//
// for v 1.21RC
//
// running with GOEXPERIMENT=loopvar we can see the output of the gorutines changes from “three three three” to “three two one”.
//
// Go routines are scheduled randomly since Go 1.5. 
//Another potentially breaking change is that the runtime now sets the default number of threads to run simultaneously, 
// defined by GOMAXPROCS, to the number of cores available on the CPU. In prior releases the default was 1. 
// Programs that do not expect to run with multiple cores may break inadvertently. 
// They can be updated by removing the restriction or by setting GOMAXPROCS explicitly. 
// For a more detailed discussion of this change, see the design document.
//
func main() {  
    data := []string{"one","two","three"}

    for _,v := range data {
		v:=v
        go func() {
            fmt.Println(v)
        }()
    }

    time.Sleep(3 * time.Second)
    //goroutines print: three, three, three
}