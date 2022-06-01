package main

import "fmt"

func main() {
	var a []int
	var b interface{}

	fmt.Printf("a: %v '%T'\n", a, a)
	fmt.Printf("b: %v '%T'\n", b, b)
	fmt.Println("a == nil: ", a == nil)
	fmt.Println("b == nil: ", b == nil)

	b = a
	fmt.Println("\nb = a: ")
	fmt.Println("b == nil: ", b == nil)
	fmt.Printf("b: %v '%T'\n", b, b)
}
