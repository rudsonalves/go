package main

import (
	"fmt"
)

func routine(c1 <-chan int, c2 chan<- int) {
	c2 <- 5
	v := <-c1
	fmt.Println("go routine:", v)
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go routine(c1, c2)

	c1 <- 10
	v := <-c2
	fmt.Println("main:", v)
}
