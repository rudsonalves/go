package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxN, minN := 5, 5
	for i := 0; i < 1000000; i++ {
		j := int(10 * rand.Float64())
		if j > maxN {
			maxN = j
		} else if j < minN {
			minN = j
		}
		// fmt.Println(j)
	}
	fmt.Println("Max:", maxN)
	fmt.Println("Min:", minN)
}
