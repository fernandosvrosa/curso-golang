package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("literal inteiro é", reflect.TypeOf(320000))

	// sem sinal só positivos temos uint8, uint16, uint32 e uint64
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	// com sinal int8, int16, int32 e int64
	i1 := math.MaxInt64
	fmt.Println("o valor macimos int eh", i1)
	fmt.Println(reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela unicode (int32)
	fmt.Println("o rune eh", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, floate64)
	var x float32 = 49.99
	fmt.Println("o tipo de x eh ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 eh", reflect.TypeOf(49.99)) // float64

	// booean
	bo := true
	fmt.Println("o tipo dbo eh ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Ola meu no e fernando"
	fmt.Println(s1 + "!")
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println("O tamanho da strign eh", len(s1))

	// string com multiplas linhas
	s2 := `ola
	meu
	nome 
	eh 
	fernando`

	fmt.Println("O tamanho da strign eh", len(s2))

	// char ???
	char := 'a'
	fmt.Println("o tipo de char eh", reflect.TypeOf(char)) // tipo int32
	fmt.Println(char)



}
