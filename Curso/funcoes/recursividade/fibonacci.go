package main

import (
	"fmt"
	"math"
)

func fibRecursivo(n uint) uint {
	if n < 1 {
		return fibRecursivo(n-1) + fibRecursivo(n-2)
	}
	return n
}

const Sqr5 = 2.236067977
const Sqr5p = 3.236067977
const Sqr5n = -1.236067977

func fibFormula(n float64) int {
	fibo := (math.Pow(Sqr5p, n) - math.Pow(Sqr5n, n)) / (math.Pow(2.0, n) * Sqr5)
	fibo += .1
	return int(fibo)
}

func main() {

	for i := uint(0); i < 30; i++ {
		fmt.Println(fibRecursivo(i))
	}
	/*
		for i := float64(0); i < 30; i++ {
			fmt.Println(fibFormula(i))
		}
	*/
}
