package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G":{
			"Gabriel" : 13250.00,
			"Guga Pereira" : 12000.00,
		},
		"J"  : {
			"Jose joao" : 10000.00,
		},
		"P":{
			"Pedro Junior" : 1200.0,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra, funcs)
		for nome, valor := range funcs{
			fmt.Println(nome, valor)
		}
	}

}
