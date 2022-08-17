package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func f() {
	fmt.Println("Running...")
}

func main() {
	for i := 1; i < 100; i++ {
		once.Do(f)
	}
}
