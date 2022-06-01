package main

import (
	"fmt"
	"math"
)

type any = interface{}

type Medida struct {
	Valor     float64
	Incerteza float64
}

func (c Medida) String() string {
	return fmt.Sprintf("%.2f ± 𝚫%.2f", c.Valor, c.Incerteza)
}

func (c Medida) Produto(valor any) Medida {
	r := Medida{0., 0.}
	// switch t := valor.(type) {
	// case int:
	// 	v, _ := valor.(int)
	// 	r.Valor = c.Valor * float64(v)
	// 	r.Incerteza = c.Incerteza * float64(v)
	// case int64:
	// 	v, _ := valor.(int64)
	// 	r.Valor = c.Valor * float64(v)
	// 	r.Incerteza = c.Incerteza * float64(v)
	// case float64:
	// 	v, _ := valor.(float64)
	// 	r.Valor = c.Valor * v
	// 	r.Incerteza = r.Incerteza * v
	// case Medida:
	// 	v, _ := valor.(Medida)
	// 	r.Valor = c.Valor * v.Valor
	// 	r.Incerteza = math.Abs(c.Valor*v.Incerteza + c.Incerteza*v.Valor)
	// default:
	// 	fmt.Printf("Error '%v' não suportado!\n", t)
	// }
	if v, ok := valor.(int); ok {
		r.Valor = c.Valor * float64(v)
		r.Incerteza = c.Incerteza * float64(v)
	} else if v, ok := valor.(int64); ok {
		r.Valor = c.Valor * float64(v)
		r.Incerteza = c.Incerteza * float64(v)
	} else if v, ok := valor.(float64); ok {
		r.Valor = c.Valor * v
		r.Incerteza = c.Incerteza * v
	} else if v, ok := valor.(Medida); ok {
		r.Valor = c.Valor * v.Valor
		r.Incerteza = math.Abs(c.Valor*v.Incerteza + c.Incerteza*v.Valor)
	} else {
		fmt.Printf("Error '%v' não suportado!\n", valor)
	}
	return r
}

// main function
func main() {
	a := Medida{8., .5}
	b := Medida{1, .2}
	fmt.Println("   a:", a)
	fmt.Println("   b:", b)
	fmt.Println(" 2*a:", a.Produto(2))
	fmt.Println(" a*b:", a.Produto(b))
	fmt.Println("a*as:", a.Produto("as"))
}
