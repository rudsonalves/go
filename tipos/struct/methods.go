package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.Firstname, p.Lastname, p.Age)
}

func (p *Person) SetName(name string) {
	p.Firstname = name
}

func main() {
	p1 := Person{"Rudson", "Ribeiro", 49}
	p2 := Person{"Juliana", "Boasquevisque", 30}
	fmt.Println(p1.String())
	fmt.Println(p2.String())

	p1.SetName("Rosilene")
	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
