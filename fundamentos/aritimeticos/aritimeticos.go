package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("soma => ", a+b)
	fmt.Println("subtracao => ", a-b)
	fmt.Println("divisao => ", a/b)
	fmt.Println("multiplicacao => ", a*b)
	fmt.Println("Modulo => ", a%b)

	// bitwise
	fmt.Println("E => ", a&b)
	fmt.Println("ou => ", a|b)
	fmt.Println("xor => ", a^b)

	// outra operacoes usando math
	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponencial =>", math.Pow(c, d))
}
