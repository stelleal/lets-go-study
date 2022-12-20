package auxiliar

import "fmt"

// Escrever printa uma mensagem
func Escrever() {
	fmt.Println("Escrevendo de auxiliar")
	escrever2() // em minusculo funciona por pertencer ao mesmo package
}
