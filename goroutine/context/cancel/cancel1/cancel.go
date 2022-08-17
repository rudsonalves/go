package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func readNumbers(done chan<- int) {
	ch := make(chan int, 3)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("Write %d in channel\n", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	for v := range ch {
		fmt.Printf("read %d in channel\n", v)
	}
	done <- 0
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * 500)
		cancel()
	}()

	done := make(chan int)
	go readNumbers(done)

	select {
	case <-ctx.Done():
		fmt.Println("Finished by Timeout")
		break
	case <-done:
		fmt.Println("Finished by readNumbers")
	}

	fmt.Println("End...")
}
