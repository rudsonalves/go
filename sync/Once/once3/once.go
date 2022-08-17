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
	once      sync.Once
)

func openDBConnection(id int) {
	defer wg.Done()

	once.Do(func() {
		// some process...
		time.Sleep(time.Microsecond)
		// finish process
		connecton = &dbConnection{}
		fmt.Printf("Make a database connection (%d)...\n", id)
	})

	fmt.Printf("Done (%d)...\n", id)
}

func main() {
	for id := 0; id < 5; id++ {
		wg.Add(1)
		go openDBConnection(id)
	}

	wg.Wait()
}
