package main

import (
	"fmt"
	"time"
)

func intCounter(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	n := 1
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				fmt.Println("done!")
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}

func main() {
	done := make(chan struct{})

	for n := range intCounter(done) {
		fmt.Println(n)
		if n == 5 {
			fmt.Println("Run cancel()")
			done <- struct{}{}
		}
	}

	time.Sleep(time.Millisecond * 5)
}
