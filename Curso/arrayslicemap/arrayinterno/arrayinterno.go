package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	s1[2] = 7
	fmt.Println(s1, s2)

	s2[5] = 9
	fmt.Println(s1, s2)
}
