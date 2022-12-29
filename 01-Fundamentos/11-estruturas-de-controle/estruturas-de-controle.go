package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SWITCH
func diaDaSemana(numero int) string {
	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
	case 2:
		diaDaSemana = "Segunda-feira"
	case 3:
		diaDaSemana = "Terça-feira"
	case 4:
		diaDaSemana = "Quarta-feira"
	case 5:
		diaDaSemana = "Quinta-feira"
	case 6:
		diaDaSemana = "Sexta-feira"
		// fallthrough <- caso numero igual a 6, com o fallthrough o resultado seria o próximo, mesmo se número não fosse igual a 7
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana

	// Outra forma de fazer:
	//
	// switch {
	// case numero == 1:
	// 	return "Domingo"
	// case numero == 2:
	// 	return "Segunda-feira"
	// default:
	// 	return "..."
	// }
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// IF / ELSE

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

	// SWITCH em uso
	numeroDia := rand.Intn(6) + 1
	dia := diaDaSemana(numeroDia)
	fmt.Println("Hoje é " + dia)

}
