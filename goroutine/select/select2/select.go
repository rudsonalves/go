package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dbServer struct {
	name string
	ch   chan string
}

func connectDBase(server dbServer) {
	fmt.Printf("Try connection in %s...\n", server.name)
	time.Sleep(time.Microsecond * (20 + time.Duration(rand.Int63n(10))))
	server.ch <- fmt.Sprintf("%s connected!", server.name)
}

func main() {
	servers := []dbServer{
		{"Newton", make(chan string)},
		{"Einstein", make(chan string)},
		{"Plank", make(chan string)},
		{"Fermi", make(chan string)},
	}

	rand.Seed(time.Now().UnixNano())

	for _, server := range servers {
		go connectDBase(server)
	}

	select {
	case s := <-servers[0].ch:
		fmt.Println(s)
	case s := <-servers[1].ch:
		fmt.Println(s)
	case s := <-servers[2].ch:
		fmt.Println(s)
	case s := <-servers[3].ch:
		fmt.Println(s)
	default:
		fmt.Println("Out...")
	}
}
