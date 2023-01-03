package main

import "fmt"

func main() {
	//					buffer ↴ 
	canal := make(chan string, 2)
	canal <- "Olá Mundo"
	canal <- "Programando em Go" 
	// canal <- "Causaria um deadlock"

	mensagem := <- canal
	mensagem2 := <- canal
	// mensagem3 := <- canal // outro deadlock

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}