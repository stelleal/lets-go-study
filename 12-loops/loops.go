package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 3 {
		i++
		fmt.Println("i =", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 5; j += 2 {
		fmt.Println("j =", j)
		time.Sleep(time.Second)
	}

	// Iteração
	//
	nomes := [3]string{"Ana", "Maria", "Pedro"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for index, letra := range "PALAVRA" {
		fmt.Println(index, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Matheus",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// For Infinito
	// for{
	// 	fmt.Println("for loop Infinito")
	// 	time.Sleep(time.Second)
	// }
}
