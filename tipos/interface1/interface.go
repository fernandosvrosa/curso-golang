package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobreNome string
}

type produto struct {
	nome  string
	preco float64
}

// interface sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobreNome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - RS %2.f", p.nome, p.preco)
}

func imprimnir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa  imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimnir(coisa)

	coisa = produto{"Marterlo", 25.5}
	fmt.Println(coisa.toString())
	imprimnir(coisa)

	imprimnir(produto{"Marterlo", 25.5})
}
