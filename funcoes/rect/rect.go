package main

import "fmt"

type Retangulo struct {
	larg int
	comp int
}

func (r Retangulo) Area() int {
	return r.comp * r.larg
}

func (r Retangulo) Perimetro() int {
	return 2 * (r.larg + r.comp)
}

func (r *Retangulo) SetRet(larg, comp int) {
	r.larg = larg
	r.comp = comp
}

func main() {
	r1 := Retangulo{4, 3}
	fmt.Println(r1)
	fmt.Println("Área: ", r1.Area())
	fmt.Println("Perímetro: ", r1.Perimetro())

	r1.SetRet(5, 8)
	fmt.Println(r1)
	fmt.Println("Área: ", r1.Area())
	fmt.Println("Perímetro: ", r1.Perimetro())
}
