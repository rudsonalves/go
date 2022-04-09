package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 2)
	num := copy(y, x)
	fmt.Println("x:", x, " y:", y, " num:", num)

	z := make([]int, 2)
	fmt.Println("z:", z)
	copy(z, x[2:])
	fmt.Println("z:", z)
	num = copy(x[:3], x[1:])
	fmt.Println("x:", x)

	d := [4]int{5, 6, 7, 8}
	copy(y, d[:])
	fmt.Println("y:", y)
	copy(d[:], x)
	fmt.Println("d:", d)
}
