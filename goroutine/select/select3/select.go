package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func connectDBase(server string, ch chan<- string, done <-chan struct{}) {
	defer wg.Done()
	fmt.Printf("Try connection in %s...\n", server)
	time.Sleep(time.Microsecond * (20 + time.Duration(rand.Int63n(10))))
	select {
	case ch <- server:
		return
	case <-done:
		fmt.Printf("Server %s close!\n", server)
		return
	}
}

func main() {
	serverList := []string{"Newton", "Einstein", "Plank", "Fermi"}
	ch := make(chan string)
	done := make(chan struct{}, len(serverList))

	rand.Seed(time.Now().UnixNano())

	for _, s := range serverList {
		wg.Add(1)
		go connectDBase(s, ch, done)
	}

	doneAll := func() {
		for i := 0; i < len(serverList); i++ {
			done <- struct{}{}
		}
	}

	server := <-ch
	doneAll()

	fmt.Printf("Server %s connected.\n", server)
	wg.Wait()
	close(ch)
}
