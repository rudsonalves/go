package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func BookHotel(hotel string, ch chan<- string, done <-chan struct{}) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(200)))

	select {
	case <-done:
		fmt.Printf("Hotel %s fail...\n", hotel)
	case ch <- hotel:
	}
}

func main() {
	hotels := []string{
		"Park Lane",
		"WestHouse",
		"Royalton Park Avenue",
		"Union Square",
		"Thompson",
	}
	done := make(chan struct{}, len(hotels)-1)
	ch := make(chan string)

	rand.Seed(time.Now().Local().UnixMicro())

	for _, hotel := range hotels {
		wg.Add(1)
		go BookHotel(hotel, ch, done)
	}

	hotel := <-ch
	for i := 0; i < len(hotels)-1; i++ {
		done <- struct{}{}
	}
	fmt.Printf("Successful booking in Hotel %s\n", hotel)

	wg.Wait()
}
