package main

import (
	"errors"
	"fmt"
)

func main() {

	//INTEIROS (default = 0)
	// int -> int8 int16 int32 int64
	var numero1 int64 = -1111111
	// uint -> unasigned int, numeros positivos apenas
	var numero2 uint32 = 2222
	// alias
	// int32 = rune
	var numero3 rune = 3333
	// int8 = byte
	var numero4 byte = 4

	// REAIS (default = 0)
	// float -> float32 float64
	var numReal1 float64 = 111.1

	// STRING -> aspas duplas "" (default = "")
	var str string = "Texto"
	str2 := "Texto2"
	// aspas simples converte o caractere para ASCII, vira rune
	char := 'A'

	// BOOLEAN (default = nil)
	var boolean bool = true

	// ERROR (default = nil)
	var erro error = errors.New("Meu Erro")

	fmt.Println(numero1, numero2, numero3, numero4, numReal1, str, str2, char, boolean, erro)
}
