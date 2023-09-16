package context

import (
	"context"
	"fmt"
	"time"
)

func ExecuteWithTimeout(ctx context.Context) error {
	task1 := make(chan struct{}, 1)
	task2 := make(chan struct{}, 1)

	go func() {
		time.Sleep(1 * time.Second)
		task1 <- struct{}{}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		task2 <- struct{}{}
	}()

	task1Done := false
	task2Done := false

	for i := 0; i < 3; i++ {
		//	run := true
		//	for run {
		select {
		case <-task1:
			fmt.Println("Task 1 is done.")
			task1Done = true
		case <-task2:
			fmt.Println("Task 2 is done.")
			task2Done = true
		case <-ctx.Done():
			if task1Done && task2Done {
				return nil
			}
			return ctx.Err()
		}
		if task1Done && task2Done {
			return nil
		}
	}

	return nil
}
