package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context, cnt int) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("counter：%v，receive: %v, finish\n", cnt, v)
			return
		default:
			fmt.Printf("counter：%v，finish\n", cnt)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 3; i++ {
		go task(ctx, i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("cancel send")
	cancel()
	time.Sleep(2 * time.Second)
}
