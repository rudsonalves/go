package main

import "fmt"

func main() {
	x := 1
	y := 2

	// apenas posfix
	fmt.Println("Posfix only:")
	x++
	y--
	fmt.Printf("x: %v, y: %v", x, y)
}
