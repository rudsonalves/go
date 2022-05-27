package main

import "fmt"

func main() {
	x := []int{1, 2, 3}

	y := append(x, 44, 55)

	var z = []int{}
	z = append(x)
	z[0] = 11
	fmt.Println("x:", x, " y:", y, " z:", z)
}
