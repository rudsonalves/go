package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	Pow   = math.Pow
	Sqrt  = math.Sqrt
	Print = fmt.Println
)

func segGrau(a, b, c float64) (x1, x2 float64, err error) {
	delta := Pow(b, 2) - 4*a*c
	a2 := 2. * a
	if delta >= 0 {
		x1 = (-b + Sqrt(delta)) / a2
		x2 = (-b - Sqrt(delta)) / a2
		return
	}
	err = errors.New("raiz de número negativo")
	return
}

func main() {
	a := segGrau
	fmt.Printf("Tipo de a: %T\n", a)
	x1, x2, _ := a(2, 5, 3)
	Print(x1, x2)
}
