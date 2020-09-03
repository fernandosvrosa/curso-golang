package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i <= 10 { // nao tem while em Go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// laco infinito
		fmt.Println("para sempre")
		time.Sleep(time.Second)

	}
}
