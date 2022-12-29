package main

import "fmt"

func funcaoClosure() func() {
	texto := "Dentro de funcaoClosure"

	// respeita o ESCOPO, usando o valor de texto acima
	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro de main"
	fmt.Println(texto)

	funcaoNova := funcaoClosure()
	funcaoNova()
}
