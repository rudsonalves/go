package main

import (
	"fmt"
	"sync"
	"time"
)

type dbConnection struct{}

var (
	connecton *dbConnection
	wg        sync.WaitGroup
	mutex     sync.Mutex
)

func openDBConnection(id int) {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	if connecton == nil {
		// some process...
		time.Sleep(time.Millisecond)
		// finish process
		connecton = &dbConnection{}
		fmt.Printf("Make a database connection (%d)...\n", id)
	}

	fmt.Printf("Done (%d)...\n", id)
}

func main() {
	for id := 0; id < 5; id++ {
		wg.Add(1)
		go openDBConnection(id)
	}

	wg.Wait()
}
