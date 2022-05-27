package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func segGrau(a, b, c float64) (x1, x2 float64, err error) {
	delta := math.Pow(b, 2) - 4*a*c
	a2 := 2. * a
	if delta >= 0 {
		x1 = (-b + math.Sqrt(delta)) / a2
		x2 = (-b - math.Sqrt(delta)) / a2
		return
	}
	err = errors.New("raiz de número negativo")
	return
}

func main() {
	coef := [3][3]int{{2, 5, 3}, {2, -5, 3}, {5, 2, 3}}

	for _, c := range coef {
		fmt.Printf("%dx²+%dx+%d = 0\n", c[0], c[1], c[2])
		x1, x2, err := segGrau(float64(c[0]), float64(c[1]), float64(c[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("x1:%.2f  x2: %.2f\n\n", x1, x2)
	}
}
