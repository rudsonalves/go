package main

import "fmt"

func main() {
	x := []int{}
	y := []int{0}
	z := []int{0, 1, 3}

	fmt.Println("x:", x, " Len(x):", len(x), " Cap(x):", cap(x))
	fmt.Println("y:", y, " Len(y):", len(y), " Cap(y):", cap(y))
	fmt.Println("z:", z, " Len(z):", len(z), " Cap(z):", cap(z))

	fmt.Println("")
	x = append(x, 1)
	y = append(y, 2)
	z = append(z, 4, 5)

	fmt.Println("x:", x, " Len(x):", len(x), " Cap(x):", cap(x))
	fmt.Println("y:", y, " Len(y):", len(y), " Cap(y):", cap(y))
	fmt.Println("z:", z, " Len(z):", len(z), " Cap(z):", cap(z))
}
