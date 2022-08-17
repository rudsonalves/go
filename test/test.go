package main

import (
	"fmt"
)

func main() {
	a := [][]int{}
	b := []int{}

	b = append(b, 1, 2, 3, 4)
	a = append(a, b)
	a = append(a, b)
	fmt.Println(a)
}
