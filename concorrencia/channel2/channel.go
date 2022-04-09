package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// Channel é um tipo

func contador(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)
	go contador(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println("B")

	fmt.Println(<-c)
	fmt.Println("C")
}
