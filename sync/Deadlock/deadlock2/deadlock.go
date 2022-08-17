package main

import (
	"fmt"
	"sync"
	"time"
)

var eatWaitGroup sync.WaitGroup

type fork struct {
	sync.Mutex
}

type Philosopher struct {
	name                string
	leftName, rightName string
	leftFork, rightFork *fork
	sleepTime           time.Duration
}

func (p Philosopher) Dining() {
	defer eatWaitGroup.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("(%s) try to eat the forkful %d.\n", p.name, i)
		for {
			p.leftFork.Lock()
			fmt.Printf("(%s) grab (%s) fork.\n", p.name, p.leftName)
			if !p.rightFork.TryLock() {
				fmt.Printf("(%s) drop (%s) fork.\n", p.name, p.leftName)
				p.leftFork.Unlock()
			} else {
				break
			}
			time.Sleep(p.sleepTime)
		}

		fmt.Printf("(%s) start to eat the forkful %d.\n", p.name, i)
		time.Sleep(p.sleepTime)

		p.leftFork.Unlock()
		fmt.Printf("(%s) drop (%s) fork.\n", p.name, p.leftName)
		p.rightFork.Unlock()
		fmt.Printf("(%s) drop (%s) fork.\n", p.name, p.rightName)

		fmt.Printf("(%s) finish to eat the forkful %d.\n", p.name, i)
	}
}

func main() {
	var (
		philNames = []string{
			"A",
			"B",
			"C",
			// "D",
			// "E",
		}
		forks        []*fork
		philosophers []Philosopher
	)

	nPhil := len(philNames)

	for i := 0; i < nPhil; i++ {
		forks = append(forks, new(fork))
	}

	left, right := 0, 0
	for i := 0; i < nPhil; i++ {
		left = i - 1
		if left < 0 {
			left = nPhil - 1
		}
		right = (i + 1) % nPhil

		philosophers = append(philosophers,
			Philosopher{
				name:      philNames[i],
				leftName:  philNames[left],
				rightName: philNames[right],
				leftFork:  forks[i],
				rightFork: forks[right],
				sleepTime: time.Microsecond,
			})

		eatWaitGroup.Add(1)
		go philosophers[i].Dining()
	}

	eatWaitGroup.Wait()
}
