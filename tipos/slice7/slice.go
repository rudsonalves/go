package main

import "fmt"

func main() {
	var x []int
	y := []int{}

	fmt.Println(x == nil)
	fmt.Println(y == nil)

	x = append(x, 0)
	y = append(y, 1)

	fmt.Println(x == nil)
	fmt.Println(y == nil)
}
