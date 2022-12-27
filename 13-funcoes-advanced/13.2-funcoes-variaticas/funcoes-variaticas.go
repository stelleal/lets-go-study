package main

import "fmt"

// cria um slice com os parametros recebidos
func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

// pode receber outros parametros, mas apenas um parametro variático por func, o qual precisa ser o último
func olaNumero(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(5, 5, 5)
	fmt.Println(totalDaSoma)

	olaNumero("Olá Mundo", 1, 4, 11, 234, 556)
}
