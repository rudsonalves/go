package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64

	// forma reduzida

	area := PI * math.Pow(raio, 2)

	fmt.Println("Area da circunferência é ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Print(a, b, c, d)

	// inicializar várias variáveis em forma curta (semelhante ao python)
	var e, f bool = true, false
	g, h, i := 2, false, "epa"

	fmt.Println(e, f)
	fmt.Println(g, h, i)
}
