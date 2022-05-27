package main

import (
	"fmt"

	"github.com/rudsonalves67/mathlib"
)

func main() {
	a := mathlib.Vector{X: -1, Y: -1, Z: 0}
	b := mathlib.Vector{3, -2, 1}

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("a+b:", a.Add(b))
	fmt.Println("a-b:", a.Sub(b))
	fmt.Println("a.b:", a.Scalar(b))
	fmt.Println("axb:", a.Cross(b))
	fmt.Println("|a|:", a.Module())
	fmt.Println("|b|:", b.Module())
	fmt.Println("a-a:", a.Sub(a))
	c := a.Cross(b)
	fmt.Println("c = axb:", c)
	fmt.Println("(axb).a = 0:", c.Scalar(a))
	fmt.Println("(axb).b = 0:", c.Scalar(b))
}
