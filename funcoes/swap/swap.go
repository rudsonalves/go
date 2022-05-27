package main

import "fmt"

// swap troca os valores das duas variáveis passadas
func swap(x, y *int) {
	z := *x
	*x = *y
	*y = z
}

func main() {
	i, j := 2, 5

	fmt.Println("i:", i, " j:", j)
	swap(&i, &j)
	fmt.Println("i:", i, " j:", j)
}
