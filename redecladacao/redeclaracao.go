package main

import (
	"fmt"
)

func main() {
	var c = 0

	a, c := 1, 2
	fmt.Println(a, c)

	b, c := 3, 4
	fmt.Println(a, b, c)

	a, b, c, d := 5, 6, 7, 8
	fmt.Println(a, b, c, d)
}
