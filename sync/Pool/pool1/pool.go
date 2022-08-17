package main

import (
	"fmt"
	"sync"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = new(Person)
			p.Age = 23
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			p = personPool.Get().(*Person)
			p.Age = 23
			personPool.Put(p)
		}
	}
}

func main() {
	newPerson := personPool.Get().(*Person)

	defer personPool.Put(newPerson)

	newPerson.Name = "Jack"

	fmt.Println(newPerson)
}
