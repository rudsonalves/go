package main

import (
	T "fmt" // Abreviação do fmt pela letra 'T'
	"math"
)

func main() {
	a := 3
	b := 2

	T.Println("Operações Básicas:")
	T.Println("Soma =>", a+b)
	T.Println("Subtração =>", a-b)
	T.Println("Divisão =>", a/b)
	T.Println("Multiplicação =>", a*b)

	// Operações bitwise
	T.Println("\nOperações bitwise")
	T.Println("E =>", a&b)   // 11 & 10 = 10
	T.Println("Ou =>", a|b)  // 11 | 10 = 11
	T.Println("Xor =>", a^b) // 11 ^ 10 = 01

	// Outras operações
	T.Println("Outras operações")
	c := 3.0
	d := 2.0
	T.Println("Maior =>", math.Max(c, d))
	T.Println("Menor =>", math.Min(float64(a), float64(b)))
	T.Println("Exponenciação =>", math.Pow(float64(a), float64(b)))
}
