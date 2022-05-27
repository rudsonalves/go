package main

import "fmt"

type esportivo interface {
	ligarTurbo()
	desligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual float64
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func (f *ferrari) desligarTurbo() {
	f.turboLigado = false
}

func main() {
	carro1 := ferrari{"F40", false, 124.}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F41", false, 0}
	carro2.ligarTurbo()

	carro3 := ferrari{"F30", true, 64.3}

	fmt.Println(carro1, carro2, carro3)
}
