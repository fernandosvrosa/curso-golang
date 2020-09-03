package main

import "fmt"

func main() {
	funcsSalarios := map[string]float64{
		"Jose Joao":      11325.23,
		"Gabriela silva": 10325.23,
		"Pedro Junior":   9325.23,
	}

	funcsSalarios["Rafael Filho"] = 1350.0
	delete(funcsSalarios, "nao existe") // nao da erro

	for nome, salario := range funcsSalarios{
		fmt.Println(nome, salario)
	}
}
