package main

import "fmt"

func main() {
	var ch chan int
	fmt.Printf("ch é nil: %v\nTipo: (%T)\nValor: %v\n\n", ch == nil, ch, ch)

	ch = make(chan int)
	fmt.Printf("ch é nil: %v\nTipo: (%T)\nValor: %v", ch == nil, ch, ch)

	go func() {
		ch <- 0
	}()

	n := <-ch

	fmt.Println(n)
}
