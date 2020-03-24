package main

import (
	"context"
	"fmt"
	"time"
)

func delay(ctx context.Context, t int) {
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("running ... ")
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("delay() Done:", ctx.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("delay():", r)
	}
	return
}

func main() {
	ctx := context.Background()
	deadline := time.Now().Add(time.Duration(10) * time.Second)
	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	delay(ctx, 5)
	time.Sleep(10 * time.Second)
	cancel()
}
