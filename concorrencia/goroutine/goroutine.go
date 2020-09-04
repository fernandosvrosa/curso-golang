package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int)  {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "porque vc nao fala comigo", 3)
	// fale("Joao", "so posso falar de pois de vc", 1)

	// go fale("maria","Ei...", 500)
	// go fale("joao","Opa...", 500)


	go fale("maria","Entendi...", 10)
	fale("joao","Parabens...", 5)
}
