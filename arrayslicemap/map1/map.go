package main

import "fmt"

func main() {
	// map deve ser init
	aprovador := make(map[int]string)

	aprovador[12345678978] = "Maria"
	aprovador[12345678977] = "Ana"
	aprovador[12345678976] = "Pedro"
	fmt.Println(aprovador)

	for cpf, nome := range aprovador {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovador[12345678977])
	delete(aprovador, 12345678976)
	fmt.Println(aprovador)

}
