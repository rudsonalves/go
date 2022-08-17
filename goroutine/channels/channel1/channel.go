package main

import (
	"fmt"
)

func sayHello(id int, ch chan string) {
	name := <-ch
	fmt.Printf("(%d) Hello %s!\n", id, name)
}

func main() {
	names := []string{"Michael", "Pam", "Tim", "Dwight", "Andy", "Kelly"}
	ch := make(chan string)

	for i := 0; i < len(names); i++ {
		go sayHello(i, ch)
	}

	for _, name := range names {
		ch <- name
	}
}
