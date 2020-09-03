package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Metodo: funcao pcom receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco - p.desconto
}

func main() {
	var p produto
	p = produto{nome: "Televisor",
		preco:    2000.0,
		desconto: 50.0,
	}
	fmt.Println(p, p.precoComDesconto())
}
