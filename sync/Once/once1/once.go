package main

import (
	"fmt"
	"sync"
)

type DbConnection struct{}

var (
	dbConnOnce sync.Once
	conn       *DbConnection
)

func GetConnection(i int) *DbConnection {
	dbConnOnce.Do(func() {
		conn = &DbConnection{}
		fmt.Println("Inside", i)
	})
	fmt.Println("Outside", i)
	return conn
}

func main() {
	for i := 0; i < 5; i++ {
		_ = GetConnection(i)
		/* Result is ...
		   Inside
		   Outside
		   Outside
		   Outside
		   Outside
		   Outside
		*/
	}
}
