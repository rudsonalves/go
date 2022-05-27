package main

import "fmt"

var x int = 10

func printX() {
	fmt.Println(x)
}

func main() {
	fmt.Println(x)
	x := 20
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		x := 8
		fmt.Print(x, " ")
	}
	// x := 12
	fmt.Println("\n", x)
	printX()
}
