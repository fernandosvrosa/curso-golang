package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String", "banana"  == "banana")
	fmt.Println("!=", 3  != 2)
	fmt.Println(">", 3  > 2)
	fmt.Println("<", 3  < 2)
	fmt.Println("<=", 3  <= 2)
	fmt.Println(">=", 3  >= 2)

	d1 := time.Unix(0,0)
	d2 :=time.Unix(0,0)

	fmt.Println("datas :", d1 == d2)
	fmt.Println("datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"fernando"}
	p2 := Pessoa{"fernando"}
	fmt.Println("Pessoas: ", p1 == p2)

}
