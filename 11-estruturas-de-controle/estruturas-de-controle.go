package main

import "fmt"

func main() {
	num := 10

	if num > 15 {
		fmt.Println(num, "é maior que 15")
	} else if num == 15 {
		fmt.Println(num, "é igual a 15")
	} else {
		fmt.Println(num, "é menor que 15")
	}

	// if init -> inicia variavel apenas no escopo do if
	if outroNum := num; outroNum > 0 {
		fmt.Println("Numero maior que zero")
	}

	// fmt.Println(outroNum) -> daria erro
}
