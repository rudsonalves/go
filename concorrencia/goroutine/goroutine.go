package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Porque você não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// O comando go faz com que os comandos subsequentes sejam
	// executadas de forma concorrente
	go fale("Maria", "Ei...", 10)
	go fale("João", "Opa...", 5)

	time.Sleep(time.Second * 5)
	fmt.Println("Fim")
	// No entanto, ao alcançar o final da função main, o programa
	// termina. Ou seja, os comandos dentro da chamad do go somente
	// são executados enquanto o programa está em execução.
}
