package main

import "fmt"

func main() {
	x := []int{2, 4, 6, 8, 10, 12}
	y := x[2:5]
	z := x[:3]
	w := x[4:]
	k := x[2:3]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("w:", w)
	fmt.Println("k:", k)

	fmt.Println("\nTrocando elementos")
	y[0] = 122
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("w:", w)
	fmt.Println("k:", k)
}
