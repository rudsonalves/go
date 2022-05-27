package main

import "fmt"

type detalhes struct {
	segmento string
	marca    string
	tipo     string
}

type revista struct {
	nome    string
	autores string
	preco   float32
	detalhes
}

func main() {
	dcComics := detalhes{"Comics", "DC", "Livros"}

	revista0 := revista{
		nome:     "Batman: Ano Um",
		autores:  "David Mazzucchelli, Frank Miller",
		preco:    66.90,
		detalhes: dcComics,
	}
	revista1 := revista{"Watchmen", "Alan Moore, Dave Gibbons", 184.90, dcComics}

	fmt.Println(revista0)
	fmt.Println(revista1)

	fmt.Println(revista0.segmento)

	revista2 := revista{
		nome:    "Shazam!",
		autores: "Clayton Henry, Tim Sheridan",
		preco:   24.90,
		detalhes: detalhes{
			segmento: "Comics",
			marca:    "DC",
			tipo:     "Revista",
		},
	}
	fmt.Println(revista2)
}
