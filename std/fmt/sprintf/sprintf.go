package main

import "fmt"

type pessoa struct {
	nome      string
	sobreNome string
	idade     int
}

func (p pessoa) String() string {
	return fmt.Sprintf("O %s %s possui %d anos.", p.nome, p.sobreNome, p.idade)
}

func main() {
	a := pessoa{
		nome:      "Alonso",
		sobreNome: "Bortolini",
		idade:     36,
	}

	fmt.Print(a)
}
