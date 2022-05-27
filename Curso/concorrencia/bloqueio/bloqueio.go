package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	fmt.Println("Iniciou rotina")
	time.Sleep(time.Second * 2)
	c <- 1 // operação bloqueio
	fmt.Println("Só depois que foi lido")
	c <- 2
}

func main() {
	c := make(chan int) // canel sem buffer

	go rotina(c)
	time.Sleep(time.Second * 10)
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
