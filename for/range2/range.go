package main

import "fmt"

func main() {
	primos := []int{2, 3, 5, 7}

	for _, v := range primos {
		v += 10
		fmt.Println(v)
	}

	fmt.Println(primos)
}
