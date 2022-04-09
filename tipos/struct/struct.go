package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metodo: função com receiver (receptor)
func (p produto) precoConDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto

	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: .05,
	}

	fmt.Println(produto1.precoConDesconto())

	produto2 := produto{"Notebook", 2000., 0.1}

	fmt.Println(produto2.precoConDesconto())
}
