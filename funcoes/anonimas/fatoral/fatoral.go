package main

import "fmt"

func main() {
	fat := func(n int) int {
		f := 1
		for i := n; i > 1; i-- {
			f *= i
		}
		return f
	}

	fmt.Println(fat(3))
	fmt.Println(fat(4))
	fmt.Println(fat(5))
	fmt.Println(fat(6))
}
