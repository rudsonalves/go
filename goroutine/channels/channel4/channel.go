package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func print(ch <-chan int) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Printf("Print %02[2]d (len: %[1]d/%[3]d, loop: %[4]d)\n", len(ch), <-ch, len(ch), i)
	}
}

func main() {
	// ch := make(chan int)
	ch := make(chan int, 3)

	wg.Add(1)
	go print(ch)

	for i := 0; i < 6; i++ {
		value := i * 10
		fmt.Printf("Write %02d\n", value)
		ch <- value
	}
	fmt.Println("Wait...")
	wg.Wait()
}
