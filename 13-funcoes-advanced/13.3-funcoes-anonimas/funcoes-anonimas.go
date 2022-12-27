package main

import (
	"fmt"

)

func main() {

	func() {
		fmt.Println("Olá GO!")
	}()

	meuRetorno := func(texto string) string{
		return fmt.Sprintf("Recebendo -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(meuRetorno)
}
