package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{3, 4, 5, 6, 7}
	// Slice não é um array e sim uma parte de uma array
	s2 := a2[1:3]

	fmt.Println(a2, s2)
	fmt.Println(reflect.TypeOf(a2), reflect.TypeOf(s2))

	s3 := a2[:2]
	fmt.Println(a2, s3)

	s4 := s2[1:]

	fmt.Println(s2, s4)
	fmt.Println(len(s2), len(s4))
}
