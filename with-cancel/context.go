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
	ctx, cancel := context.WithCancel(ctx)
	for i := range gen(ctx) {
		fmt.Println("Got:", i)
		if i == 10 {
			cancel()
			break
		}
	}
	time.Sleep(time.Second)
	fmt.Println("bye!")
}
