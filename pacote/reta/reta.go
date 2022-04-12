package main

import "math"

// Inicializando com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)!
// Inicializando com letra minúscula é PRIVADO (visível apenas dentro do pacote)!

// Exemplo:
// Ponto - é público
// ponto ou _Ponto - é privado

// Ponto representa uma coordenada no plano
type Ponto struct {
	x float64
	y float64
}

func catedos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distância entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catedos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
