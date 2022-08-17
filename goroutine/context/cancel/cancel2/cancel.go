package main

import (
	"context"
	"fmt"
	"time"
)

func BookHotelA(ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel A fail...")
	case <-time.After(time.Second * 3):
		fmt.Println("Successful booking in Hotel A")
		cancel()
	}
}

func BookHotelB(ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel B fail...")
	case <-time.After(time.Second * 5):
		fmt.Println("Successful booking in Hotel B")
		cancel()
	}
}

func main() {
	ctx := context.Background()
	ctx0, cancel := context.WithCancel(ctx)
	// ctx0, cancel := context.WithTimeout(ctx, time.Second*2)

	go BookHotelA(ctx0, cancel)
	go BookHotelB(ctx0, cancel)

	i := 0
	for {
		i++
		time.After(time.Second * 1)
		// fmt.Println(i)
	}
}
