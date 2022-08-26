package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		fmt.Println("go routine")
	}()
	fmt.Println("stop")
	<-ctx.Done()
	fmt.Println("start")
}
