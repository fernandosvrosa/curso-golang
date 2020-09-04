package main

import (
	"fmt"
	"github.com/fernandosvrosa/htmlgo"
	"time"
)

func oMiasRapido(url1, url2, url3 string) string {
	c1 := htmlgo.Titulo(url1)
	c2 := htmlgo.Titulo(url2)
	c3 := htmlgo.Titulo(url3)

	// estrutura de controle para concorrencia
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam"
		//default:
		//	return "sem resposta ainda!"
	}
}

func main() {
	campeao := oMiasRapido(
		"https://www.terra.com.br",
		"https://www.youtube.com",
		"https://www.aws.com",
	)

	fmt.Println("O campeao = ", campeao)
}
