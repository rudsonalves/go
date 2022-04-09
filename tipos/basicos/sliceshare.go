package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	z := x[1:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	x[0] = 10
	y[1] = 20
	z[1] = 30

	fmt.Println("\nChange values:")
	fmt.Printf("x[0] = 10\ny[1] = 20\nz[1] = 30\n")

	fmt.Println("\nSlices share storange sometimes:")
	fmt.Println("When you take a slice from a slice, you are not making")
	fmt.Println("a copy of the data. Instead, you nw have two vaiables")
	fmt.Println("that are sharing memory.")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
