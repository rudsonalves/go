package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fibo(ch chan<- int, done <-chan struct{}) {
	defer wg.Done()

	a, b := 0, 1
	for {
		c := a + b
		a, b = b, c
		select {
		case ch <- c:
		case <-done:
			fmt.Println("\nFibo stoped...")
			close(ch)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	wg.Add(1)
	go fibo(ch, done)

	rand.Seed(time.Now().UnixMicro())

	max := 5 + rand.Intn(15)
	fmt.Printf("\nFirst %d Fibonacci numbers:\n0 1", max+2)
	for i := 0; i < max; i++ {
		fb := <-ch
		fmt.Print(" ", fb)
	}

	done <- struct{}{}
	wg.Wait()
}
