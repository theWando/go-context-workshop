package main

import (
	"context"
	"fmt"
	"time"
)

func gen(_ context.Context) <-chan int {
	out := make(chan int)
	var i int
	go func() {
		defer fmt.Println("closing go routine")
		for {
			select {
			case out <- i:
				fmt.Println("sending ", i)
				i++
			}
		}
	}()
	return out
}

func main() {
	ctx := context.Background()
	for i := range gen(ctx) {
		fmt.Println("Got:", i)
		if i == 10 {
			break
		}
	}
	time.Sleep(time.Second)
	fmt.Println("bye!")
}
