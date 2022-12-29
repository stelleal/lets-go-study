package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Olá Gopher! ʕ •ᴥ•ʔノ") // não espera a função terminar para continuar com o programa
	escrever("Programando em Go! ʕ •ᴥ•ʔっ💻")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
