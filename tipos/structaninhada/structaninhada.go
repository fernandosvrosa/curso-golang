package main

import "fmt"

type item struct {
	produdoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{
				produdoID:  1,
				quantidade: 11,
				preco:      10.0,
			},
			{
				produdoID:  2,
				quantidade: 20,
				preco:      49.0,
			},
			{
				produdoID:  3,
				quantidade: 5,
				preco:      101.0,
			},
		},
	}

	fmt.Printf("valor total do pedido eh R$ %.2f\n", pedido.valorTotal())
}
