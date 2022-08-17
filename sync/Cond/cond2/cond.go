package main

import (
	"fmt"
	"sync"
)

type Item = int

type Queue struct {
	items     []Item
	itemAdded sync.Cond
}

func NewQueue() *Queue {
	q := new(Queue)
	q.itemAdded.L = &sync.Mutex{}
	return q
}

func (q *Queue) Put(item Item) {
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()
	q.items = append(q.items, item)
	q.itemAdded.Signal()
}

func (q *Queue) GetMany(n int) []Item {
	q.itemAdded.L.Lock()
	defer q.itemAdded.L.Unlock()
	for len(q.items) < n {
		q.itemAdded.Wait()
	}
	items := q.items[:n:n]
	q.items = q.items[n:]
	return items
}

func main() {
	q := NewQueue()

	var wg sync.WaitGroup
	for n := 10; n > 0; n-- {
		wg.Add(1)
		go func(n int) {
			items := q.GetMany(n)
			fmt.Printf("%2d: %2d\n", n, items)
			wg.Done()
		}(n)
	}

	for i := 0; i < 100; i++ {
		q.Put(i)
	}

	wg.Wait()
	fmt.Println(q)
}
