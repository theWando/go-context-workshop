package main

import (
	"context"
	"fmt"
)

func gen(_ context.Context) <-chan int {
	out := make(chan int)
	var i int
	go func() {
		for {
			select {
			case out <- i:
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
}
