package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // a variávei i está definida epenas dentro do if
		fmt.Println("Ganhou", i)
	} else {
		fmt.Println("Perdeu", i)
	}
}
