package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "3"
		variavel4 int
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Morango", "Copo"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"

	// trocando valores de variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
