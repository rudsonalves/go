package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x", x)
	fmt.Println("x[:2]:", y)
	fmt.Println("x[1:]:", z)
	fmt.Println("x[1:3]:", d)
	fmt.Println("x[:]:", e)
}
