package main

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}

func main() {
	r1, r2 := troca(1, 2)
	println(r1)
	println(r2)
}
