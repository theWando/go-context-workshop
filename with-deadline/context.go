package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	var i int
	go func() {
		defer fmt.Println("closing go routine")
		for {
			select {
			case out <- i:
				fmt.Println("sending ", i)
				time.Sleep(200 * time.Millisecond)
				i++
			case <-ctx.Done():
				fmt.Println("err:", ctx.Err())
				close(out)
				return
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second))
	for i := range gen(ctx) {
		fmt.Println("Got:", i)
		if i == 10 {
			break
		}
	}
	time.Sleep(time.Second)
	fmt.Println("bye!")
}
