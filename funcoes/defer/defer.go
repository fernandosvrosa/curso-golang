package main

import "fmt"

func opterValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero >= 5000 {
		fmt.Println("valor alto...")
		return 5000
	}
	fmt.Println("valor baixo...")
	return numero

}

func main() {
	fmt.Println(opterValorAprovado(6000))
	fmt.Println(opterValorAprovado(3000))
}
