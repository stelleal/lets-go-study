package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo de main")
	auxiliar.Escrever()
	// auxiliar.escrever2()  <-- nao funciona por comecar com minusculo

	erro := checkmail.ValidateFormat("meuemail@gmail.com")
	fmt.Println(erro)
}
