package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var hotel_reservation = time.After

func BookHotel(hotel string, ctx context.Context, cancel context.CancelFunc) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Printf("Hotel %s fail...\n", hotel)
	case <-hotel_reservation(time.Millisecond * time.Duration(rand.Int63n(200))):
		cancel()
		fmt.Printf("Successful booking in Hotel %s\n", hotel)
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
	ctx, cancel := context.WithCancel(context.Background())

	rand.Seed(time.Now().Local().UnixMicro())

	for _, hotel := range hotels {
		wg.Add(1)
		go BookHotel(hotel, ctx, cancel)
	}

	wg.Wait()
}
