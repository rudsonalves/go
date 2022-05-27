package main

import "fmt"

func soma(nums ...int) (s int) {
	if len(nums) < 1 {
		return
	}
	for _, i := range nums {
		s += i
	}
	return
}

func main() {
	// fmt.Println(soma())
	// fmt.Println(soma(1))
	// fmt.Println(soma(1, 2))
	// fmt.Println(soma(1, 2, 3))
	// fmt.Println(soma(1, 2, 3, 4))
	numeros := []int{1, 2, 3, 4}
	fmt.Println(soma(numeros...))
	fmt.Println(soma(numeros[:3]...))
}
