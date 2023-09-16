package main

import (
	"context"
	"fmt"
	"log"
	"time"

	contextPkg "github.com/mccuskero/sandbox/pkg/context"
)

func main() {
	fmt.Println("Starting context timeout test.")

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := contextPkg.ExecuteWithTimeout(ctx); err != nil {
		log.Fatalf("error: %#v\n", err)
	}

	//	time.Sleep(5*time.Second)

	log.Println("Success to process in time")
}
