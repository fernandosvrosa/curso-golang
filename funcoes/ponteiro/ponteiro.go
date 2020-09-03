package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao: um ponteira é representado por um *

func inc2(n *int) {
	//revisao: * e usado para desreferenciar, ou seja,
	// ter acessso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n) // passando valor

	// revisao: & usado para obeter o endereco da variavel
	inc2(&n) // passando referencia
	fmt.Println(n)

}
