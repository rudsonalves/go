package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readNumbers(done chan<- struct{}) {
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
	done <- struct{}{}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	timeOut := make(chan struct{})

	go func() {
		time.Sleep(time.Millisecond * 500)
		timeOut <- struct{}{}
	}()

	done := make(chan struct{})
	go readNumbers(done)

	select {
	case <-timeOut:
		fmt.Println("Finished by Timeout")
		break
	case <-done:
		fmt.Println("Finished by readNumbers")
	}

	fmt.Println("End...")
}
