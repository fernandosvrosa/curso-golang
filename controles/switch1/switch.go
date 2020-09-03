package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "B"
	case 2, 1, 0:
		return "E"
	default:
		return "nota invalida"

	}
}

func main() {
	nota := "Nota "
	fmt.Println(nota, notaParaConceito(9.8))
	fmt.Println(nota, notaParaConceito(6.9))
	fmt.Println(nota, notaParaConceito(2.1))
	fmt.Println(nota, notaParaConceito(11.0))
}
