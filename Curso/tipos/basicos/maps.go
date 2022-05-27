package main

import "fmt"

func main() {
	var m1 map[string]int
	m2 := map[string]int{}
	m3 := map[string][]int{
		"Orcas":   []int{1, 2, 3},
		"Lions":   []int{4, 5},
		"Kittens": []int{6, 7, 8, 9},
	}

	fmt.Println(m1, m2)
	fmt.Println(m3)
}
