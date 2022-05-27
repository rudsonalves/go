package main

import (
	"fmt"
	"math"
)

const (
	pi     = math.Pi
	BRANCO = iota
	VERMELHO
	LARANJA
	AMARELO
	VERDE
	AZUL
	VIOLETA
	PRETO
)

var (
	colorName = []string{"branco", "vermelho", "laranja", "verde", "azul", "violeta", "preto"}
)

type Circulo struct {
	Raio float64
	Cor  int
}

type Retangulo struct {
	Comprimento float64
	Largura     float64
	Cor         int
}

type TriRetangulo struct {
	Base   float64
	Altura float64
	Cor    int
}

// Metodos
func (c Circulo) Area() float64 {
	return pi * math.Pow(c.Raio, 2)
}

func (c *Circulo) MudaCor(cor int) {
	c.Cor = cor
}

func (c Circulo) Perimetro() float64 {
	return 2 * pi * c.Raio
}

func (r Retangulo) Area() float64 {
	return r.Comprimento * r.Largura
}

func (r *Retangulo) MudaCor(cor int) {
	r.Cor = cor
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Comprimento + r.Largura)
}

func (t TriRetangulo) Area() float64 {
	return t.Base * t.Altura / 2
}

func (t *TriRetangulo) MudaCor(cor int) {
	t.Cor = cor
}

func (t TriRetangulo) Perimetro() float64 {
	return t.Base + t.Altura + math.Sqrt(math.Pow(t.Base, 2)+
		math.Pow(t.Altura, 2))
}

type Shaper interface {
	Area() float64
	Perimetro() float64
}

func AreaTotal(shapes ...Shaper) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func PerimetroTotal(shapes ...Shaper) float64 {
	var perimetro float64
	for _, p := range shapes {
		perimetro += p.Perimetro()
	}
	return perimetro
}

func main() {
	r := Retangulo{6, 5, 1}
	c := Circulo{3, 1}
	t := TriRetangulo{5, 3, 5}

	fmt.Printf("Área do retângulo:  %.2f m² (%s)\n", r.Area(), colorName[r.Cor])
	fmt.Printf("Área do círculo:    %.2f m² (%s)\n", c.Area(), colorName[c.Cor])
	fmt.Printf("Área do triângulo:  %.2f m² (%s)\n", t.Area(), colorName[t.Cor])

	fmt.Println("\nÁrea total:       ", AreaTotal(r, c, t))
	fmt.Println("Perímetro total:  ", PerimetroTotal(r, c, t))
}
