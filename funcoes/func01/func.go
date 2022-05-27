package main

import (
	"fmt"
	"math"
)

func areaQuad(aresta int) int {
	return int(math.Pow(float64(aresta), 2))
}

func areaRet(base, largura int) (area int) {
	area = base * largura
	return
}

func segGrau(a, b, c float64) (x1, x2 float64) {
	delta := math.Pow(b, 2) - 4*a*c
	a2 := 2. * a
	if delta >= 0 {
		x1 = (-b + math.Sqrt(delta)) / a2
		x2 = (-b - math.Sqrt(delta)) / a2
		return
	}
	return
}

func main() {
	a, b := 5, 7

	fmt.Printf("Área de quadrado de aresta %d: %d\n", a, areaQuad(a))
	fmt.Printf("Área de retângulo de %dx%d: %d\n", a, b, areaRet(a, b))

	x1, x2 := segGrau(2, 5, 3)
	fmt.Printf("2x²+5x+3=0  x1:%.2f  x2: %.2f\n", x1, x2)
	x1, x2 = segGrau(2, -5, 3)
	fmt.Printf("2x²-5x+3=0  x1:%.2f  x2: %.2f\n", x1, x2)
}
