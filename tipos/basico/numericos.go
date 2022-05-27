package main

import (
	"fmt"
)

func main() {
	var (
		i, j = 10, 20
		x    = 2.5
		z    = 6.8 + 4.9i
	)
	// i, j, x, z := 10, 20, 2.5, 6.8+4.9i
	fmt.Printf("i: %v  tipo: %T\n", i, i)
	fmt.Printf("j: %v  tipo: %T\n", j, j)
	fmt.Printf("x: %v  tipo: %T\n", x, x)
	fmt.Printf("z: %v  tipo: %T\n", z, z)
}
