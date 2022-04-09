package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numéricos inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uinir8, uinit16, unint32, uinit64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))
	// com sinal... init8, init16, init32, init64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do init64 é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	// mapeamento da tabela Unicode (init32)
	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32 e float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("Valor de x é", x)
	fmt.Println("O tipo do 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("O valor de bo é", bo)
	fmt.Println("O valor degado (operador !) de bo é", !bo)

	//string
	s1 := "Olá mundo"
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))
	fmt.Println("O valor de s1 é", s1)
	fmt.Println("String s1 (" + s1 + ") pode concatenar com operador soma (+)")
	fmt.Println("Função len(), retorna o comprimento da string: len(s1) =", len(s1))

	// string com várias linhas com crase
	s2 := `Olá
	meu 
	mome 
	é
	Rudson`
	fmt.Println(s2)
}
