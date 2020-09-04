package main

import (
	"fmt"
	"time"
)

func rorina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("so depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer
	go rorina(c)

	fmt.Println(<-c) // operacao bloqueante
	fmt.Println("foi lido")
	fmt.Println(<-c) //deradlock
	fmt.Println("fim")
}
