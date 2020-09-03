package main

import "math"

// inicializando com letra maiuscula eh publico (visibilidade dentro e fora do pacote)!
// inicializando com letra minuscula eh privado (visibilidade apenas dentro do pacote)!

// por exemplo
// Ponto - gerara algo publico
// ponto ou _Ponto - gera algo provado

// Ponto representa uma cordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distacia eh reposnsavel por calcular a distancia linear entre dois pontos
func Distacia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
