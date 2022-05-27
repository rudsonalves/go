package main

import (
	"fmt"
)

func main() {
	var (
		x       = []int{0, 1}
		y       = []int{}
		cap_old = cap(x)
		cap_new = cap_old
		strF    = "  %6v   %5d %5d  %5d\n"
	)

	fmt.Println("Últimos 3 elementos de x    Comp.  Cap.  Incr.")
	for i := 2; i < 10000; i++ {
		x = append(x, i) // adiciona novo elemento a slice x
		cap_new = cap(x)
		// Compara a nova capacidade com a anterior
		if cap_new > cap_old {
			// pega um slice de apenas os últimos 3 elemnetos de x
			y = x[len(x)-3:]
			// imprime as informações
			fmt.Printf(strF, y, len(x), cap(x), cap_new-cap_old)
			// atualiza a capacidade anterior com a nova capacidade
			cap_old = cap_new
		}
	}
}
