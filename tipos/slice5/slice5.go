package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	y := x[:2:2]
	z := x[3:5:5]

	fmt.Println("a. x:", x, "y:", y, "z:", z)

	fmt.Println("\nb. Append [30, 40] a y")
	y = append(y, 30, 40)
	fmt.Println("c. x:", x, "y:", y, "z:", z)

	fmt.Println("\nd. Append [70] a z")
	z = append(z, 70)
	fmt.Println("e. x:", x, "y:", y, "z:", z)

	fmt.Println("\nf. Append [77] a a")
	x = append(x, 77)
	fmt.Println("g. x:", x, "y:", y, "z:", z)

	fmt.Println("\nh. Append [80] a z")
	z = append(z, 80)
	fmt.Println("i. x:", x, "y:", y, "z:", z)

	fmt.Println("\nj. z[1] = 50")
	z[1] = 50
	fmt.Println("k. x:", x, "y:", y, "z:", z)
}
