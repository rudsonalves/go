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
	var c []chan string

	rand.Seed(time.Now().UnixNano())

	for i, server := range serverList {
		c = append(c, make(chan string))
		go connectDBase(server, c[i])
	}

	select {
	case s := <-c[0]:
		fmt.Println(s)
	case s := <-c[1]:
		fmt.Println(s)
	case s := <-c[2]:
		fmt.Println(s)
	case s := <-c[3]:
		fmt.Println(s)
	}
}
