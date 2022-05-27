package main

import "fmt"

func main() {
	x := make([]int, 3, 10)
	var y = make([]float32, 0, 12)

	fmt.Println("x:", x, " len(x):", len(x), " cap(x):", cap(x))
	fmt.Println("y:", y, " len(y):", len(y), " cap(y):", cap(y))
}
