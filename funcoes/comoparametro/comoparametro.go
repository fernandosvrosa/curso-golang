package main

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int , p1 , p2 int) int {
	return funcao(p1,p2)
}

func main() {
	println(exec(multiplicacao, 2,2))
}
