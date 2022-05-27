package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:4]

	x[0] = 11
	z[0] = 33

	fmt.Println("x:", x, " cap:", cap(x))
	fmt.Println("y:", y, " cap:", cap(y))
	fmt.Println("z:", z, " cap:", cap(z))

	y = append(y, 34, 44, 55)
	x = append(x, 66)
	z = append(z, 77)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
