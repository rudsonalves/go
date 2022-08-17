package main

import (
	"fmt"
	"sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()

	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-done:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	done := make(chan struct{}, 7)
	in1 := gen(done, 1, 2, 3)
	in2 := gen(done, 4, 5, 6, 7)
	in3 := gen(done, 8, 9)

	c1 := sq(done, in1)
	c2 := sq(done, in2)
	c3 := sq(done, in3)

	for n := range merge(done, c1, c2, c3) {
		fmt.Println(n)
		if n == 4 {
			for i := 0; i < cap(done); i++ {
				done <- struct{}{}
				// fmt.Println("Done:", i)
			}
			break
		}
	}
}
