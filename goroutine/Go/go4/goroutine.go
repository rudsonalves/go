package main

import (
	"fmt"
	"sync"
	"time"
)

func function(index int, t time.Duration, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		time.Sleep(t)
		fmt.Printf("Function %02d: %d\n", index, i)
	}
	fmt.Println("Done before...")
	wg.Done()
	fmt.Println("Done after...")
}

func main() {
	var wg sync.WaitGroup
	const maxGo = 3

	wg.Add(maxGo)
	fmt.Println("Starting...")

	for i := 0; i < maxGo; i++ {
		go function(i, time.Second*1, &wg)
	}

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Ending...")
}
