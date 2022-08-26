package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}

func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("cancel")
		return
	}
	fmt.Println("not cancel")
}
