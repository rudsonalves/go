package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 0., 0., 0.

	fmt.Println("Resolve uma equação de 2o grau")
	fmt.Println("ax² + bx + c = 0")
	fmt.Print("\nEntre a, b, c: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Printf("%.2fx²+%.2fx + %.2f = 0.00\n", a, b, c)

	d := math.Pow(b, 2) - 4*a*c

	a2, sqrd := 2*a, math.Sqrt(d)

	x1, x2 := (-b+sqrd)/a2, (-b-sqrd)/a2

	fmt.Printf("x1: %.2f  x2: %.2f\n", x1, x2)
	fmt.Println("Teste x1:", a*math.Pow(x1, 2)+b*x1+c)
	fmt.Println("Teste x2:", a*math.Pow(x2, 2)+b*x2+c)
}
