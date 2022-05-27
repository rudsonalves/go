package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // não precisa passar o retorno pois está na definição do retorno
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
