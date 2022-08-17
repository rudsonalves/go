package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func function(index int, t time.Duration) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(t)
		fmt.Printf("Function %02d: %d\n", index, i)
	}
}

func main() {
	const maxGo = 3

	wg.Add(maxGo)
	fmt.Println("Starting...")

	for i := 0; i < maxGo; i++ {
		go function(i, time.Second*1)
	}

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Ending...")
}
