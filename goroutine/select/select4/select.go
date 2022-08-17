package main

import (
	"fmt"
	"math/rand"
	"time"
)

func connectDBase(server string, ch chan<- string) {
	fmt.Printf("Try connection in %s...\n", server)
	time.Sleep(time.Microsecond * (20 + time.Duration(rand.Int63n(10))))
	ch <- fmt.Sprintf("%s connected!", server)
}

func main() {
	serverList := []string{"Newton", "Einstein", "Plank", "Fermi"}
	c := make(chan string)

	rand.Seed(time.Now().UnixNano())

	for _, server := range serverList {
		go connectDBase(server, c)
	}

	s := <-c
	fmt.Println(s)
}
