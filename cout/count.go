package main

import (
	"fmt"
	"strings"
)

func main() {
	var contador [3]int

	// frase de https://pt.wikipedia.org/wiki/Albert_Einstein
	frase := `Nascido em uma família de judeus alemães, mudou-se 
	para a Suíça ainda jovem e iniciou seus estudos na Escola 
	Politécnica de Zurique. Após dois anos procurando emprego, 
	obteve um cargo no escritório de patentes suíço enquanto ingressava 
	no curso de doutorado da Universidade de Zurique. Em 1905, publicou 
	uma série de artigos acadêmicos revolucionários. Uma de suas obras 
	era o desenvolvimento da teoria da relatividade especial. Percebeu, 
	no entanto, que o princípio da relatividade também poderia ser 
	estendido para campos gravitacionais, e com a sua posterior teoria 
	da gravitação, de 1916, publicou um artigo sobre a teoria da 
	relatividade geral. `

	// remove \n, \t, . e ,
	frase = strings.Replace(frase, "\n", " ", -1)
	frase = strings.Replace(frase, "\t", " ", -1)
	frase = strings.Replace(frase, ",", " ", -1)
	frase = strings.Replace(frase, ".", " ", -1)
	// remove 2 espaços conssecutivos
	frase = strings.Replace(frase, "  ", " ", -1)
	frase = strings.Replace(frase, "  ", " ", -1)

	// remove espaços e divide a frase
	frase = strings.TrimSpace(frase)
	split := strings.Split(frase, " ")

	for _, word := range split {
		if len(word) <= 5 {
			contador[0]++
		} else if len(word) <= 10 {
			contador[1]++
		} else {
			contador[2]++
		}
	}

	fmt.Printf("%2d palavras com até 5 letras\n", contador[0])
	fmt.Printf("%2d palavras com 6 a 10 letras\n", contador[1])
	fmt.Printf("%2d palavras com mais de 10 letras", contador[2])
}
