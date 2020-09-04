package main

import (
	"fmt"
	"github.com/fernandosvrosa/htmlgo"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		htmlgo.Titulo("https://www.terra.com.br", "http://www.google.com.br"),
		htmlgo.Titulo("https://www.youtube.com", "https://www.aws.com"),
	)

	fmt.Println("primeiro = ", <-c)
	fmt.Println("segundo = ", <-c)
	fmt.Println("terceiro = ", <-c)
	fmt.Println("quarto = ", <-c)
}
