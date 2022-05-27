package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// start a new random number
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var count = [10]int{}
	const max_count = 10000

	for i := 0; i < max_count; i++ {
		index := rand.Intn(10)
		count[index]++
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Ocorrências de %d: %1.1f%%\n", i,
			float64(max_count)/float64(count[i]))
	}
}
