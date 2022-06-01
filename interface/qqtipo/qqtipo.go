package main

import "fmt"

func main() {
	var i interface{}

	i = 5
	fmt.Println(i)

	i = "Albert"
	fmt.Println(i)

	i = 3 + 5i
	fmt.Println(i)

	z := func(x int) int {
		return x * x
	}
	i = z
	fmt.Printf("%v '%T'\n", i, i)

	i = struct {
		Nome string
		RG   int
	}{"Albert", 99999999999}
	fmt.Printf("%v '%T'\n", i, i)
}
