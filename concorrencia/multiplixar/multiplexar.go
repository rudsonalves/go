package main

import (
	"fmt"

	"githut.com/GoLang/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - juntar messagens num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	t1 := []string{"https://diolinux.com.br", "https://dbikeshop.com.br",
		"https://www.cod3r.com.br", "https://google.com"}
	t2 := []string{"https://github.com", "https://youtube.com",
		"https://www.udemy.com/"}

	c := juntar(html.Titulo(t1...), html.Titulo(t2...))

	fmt.Println("Em ordem de resposta:")
	for i := 0; i < len(t1)+len(t2); i++ {
		fmt.Printf("%d: ", i+1)
		fmt.Print(<-c, "\n")
	}
}
