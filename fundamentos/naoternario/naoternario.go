package main

import "fmt"

// nao tem ternario no go

func obterResultado(nota float64) string {
	// return nota >= 6 ? "aprovado" : "reprovado"
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovador"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
