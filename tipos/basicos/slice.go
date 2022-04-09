package main

import "fmt"

func main() {
	var x = []int{1, 2, 3}
	var y = []int{1, 5: 5, 10: 1, 15}

	fmt.Println(x)
	fmt.Println(y)

	y[3] = 4
	fmt.Println(y)

	var z []int
	fmt.Println(z)
	fmt.Println(z == nil)

	fmt.Println("\nSizes:")
	fmt.Println("len(y):", len(y))
	fmt.Println("len(z):", len(z))

	fmt.Println("\nAppend:")
	y = append(y, 1, 2, 3)
	fmt.Println("y = append(1,2,3):", y)
	x = append(x, y...)
	fmt.Println("x = append(1,2,3):", x)

	fmt.Println("\ncap (capacidade) function:")
	var x0 []int
	fmt.Println(x0, len(x0), cap(x0))
	x0 = append(x0, 10)
	fmt.Println(x0, len(x0), cap(x0))
	x0 = append(x0, 20)
	fmt.Println(x0, len(x0), cap(x0))
	x0 = append(x0, 30)
	fmt.Println(x0, len(x0), cap(x0))
	x0 = append(x0, 40)
	fmt.Println(x0, len(x0), cap(x0))
	x0 = append(x0, 50)
	fmt.Println(x0, len(x0), cap(x0))

	fmt.Println("\nStart a slice with a correct initial capacity with a make function")
	x1 := make([]int, 0, 10)
	fmt.Println("x1 := make([]int, 0, 10)", x1)
	fmt.Printf("x1 is a slice with len(x1) %d and capacity %d1\n", len(x1), cap(x1))

	fmt.Println("\nDeclarete a slice:")
	fmt.Println("var a = []int{} // empty slice")
	var a = []int{}
	fmt.Println("var b = []int{1, 2, 3}")
	var b = []int{1, 2, 3}
	fmt.Println("c := []int{4, 5, 6}")
	c := []int{4, 5, 6}
	fmt.Println("a, b, c =", a, b, c)
}
