package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	task := make(chan int)
	fmt.Println(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("...loading")
			}
			// update time 0.3seconds
			time.Sleep(300 * time.Millisecond)
		}
	}()
	fmt.Println(2)
	time.Sleep(3 * time.Second)
	fmt.Println(3)
	for i := 0; i < 5; i++ {
		task <- i
	}
	fmt.Println(4)
	cancelFunc()
}
