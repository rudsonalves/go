package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println("Function x:", x)
	}
	return funcao
}

func main() {
	x := 20

	imprimeX := closure()
	imprimeX()

	fmt.Println("Local x:", x)
}
