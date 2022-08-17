package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Stress(value int) {
	mutex.Lock()
	defer mutex.Unlock()

	b.Balance += value
	time.Sleep(time.Microsecond)
	b.Balance -= value

	balance := b.Balance
	fmt.Printf("Balance: %6d\n", balance)
	if balance < 0 {
		log.Fatalf("Balance shoud not be less that zero: %d\n", balance)
	}
}

func main() {
	account := BankAccount{Balance: 100}

	for i := 0; i < 10000; i++ {
		go account.Stress(1000)
	}
}
