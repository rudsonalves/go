package main

import (
	"context"
	"fmt"
	"time"
)

func intCounter(ctx context.Context) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done()!")
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for n := range intCounter(ctx) {
		fmt.Println(n)
		if n == 5 {
			fmt.Println("Run cancel()")
			cancel()
		}
	}

	time.Sleep(time.Millisecond * 5)
}
