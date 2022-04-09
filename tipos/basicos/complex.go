package main

import (
	"fmt"
	"math"
)

func main() {
	z1 := complex(2.5, 3.1)
	z2 := complex(10.2, 2)
	z3 := z1 + z2
	const Pi = math.Pi

	fmt.Println(z3)
	fmt.Printf("z1: %v\n", z3)
	fmt.Println("Real z1:", real(z3))
	fmt.Println("Imag z1:", imag(z3))
	fmt.Println("Pi:", Pi)
}
