package main

import "fmt"

type any = interface{}

func main() {
	var x any

	x = 10
	i := x.(int)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", x, x)
	x = 10.0
	j, ok := x.(int)
	if ok {
		fmt.Printf("j: %v %T\n", j, j)
	}
	k, ok := x.(float64)
	if ok {
		fmt.Printf("k: %v %T\n", k, k)
	}
	// v := x.(type)
	// fmt.Println(v == x)
}
