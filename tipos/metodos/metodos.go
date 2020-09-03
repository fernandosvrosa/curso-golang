package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompreto(nomeCompreto string) {
	partes := strings.Split(nomeCompreto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Fernando", "Rosa"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompreto("Jose Rosa")
	fmt.Println(p1.getNomeCompleto())
}
