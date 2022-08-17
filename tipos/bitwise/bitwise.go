package main

import "fmt"

func main() {
	a := 14
	b := 2
	c := ^a << b
	d := ^(a << b)
	e := (^a) << b

	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", d)
	fmt.Printf("%b\n", e)
}
