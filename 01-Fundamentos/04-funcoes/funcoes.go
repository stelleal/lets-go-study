package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// funcoes podem ter mais de 1 retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(1, 2)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("Minha função f")

	// usando _ para declarar uma variável não utilizada
	_, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao)
}
