package main

import (
	"fmt"
	"time"
)

func selectCase(ch1, ch2, ch3 chan<- int, done <-chan struct{}) {
	for {
		select {
		case ch1 <- 1:
			ch1 = nil
			continue
		case ch2 <- 2:
			ch2 = nil
			continue
		case ch3 <- 3:
			ch3 = nil
			continue
		case <-done:
			fmt.Println("Goroutine finised...")
			return
		}
		// more code here...
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	done := make(chan struct{})

	go selectCase(ch1, ch2, ch3, done)

	fmt.Printf("ch%d is nil\n", <-ch1)
	fmt.Printf("ch%d is nil\n", <-ch3)
	fmt.Printf("ch%d is nil\n", <-ch2)

	done <- struct{}{}
	time.Sleep(time.Microsecond * 100)
}
