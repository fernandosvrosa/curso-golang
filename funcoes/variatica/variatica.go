package main

import "fmt"

func merdia(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	result := merdia(9.0, 8.5, 6.0, 10.0)
	fmt.Println(result)
}
