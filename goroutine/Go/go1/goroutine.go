package main

import (
	"fmt"
	"time"
)

func function(index int, t time.Duration) {
	for i := 0; i < 10; i++ {
		time.Sleep(t)
		fmt.Printf("Function %02d: %d\n", index, i)
	}
}

func main() {

	for i := 0; i < 5; i++ {
		go function(i, time.Second*1)
	}

	fmt.Println("Starting...")
	time.Sleep(time.Second * 3)
	fmt.Println("Ending...")
}
