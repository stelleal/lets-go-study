package main

import "fmt"

func main() {
	// Aritmeticos + - / * %
	calculosDoidos := 2 + 3 - 1/4*10%2
	fmt.Println(calculosDoidos)

	var num1 int16 = 10
	var num2 int16 = 15 // se fosse outro tipo de int daria erro, exemplo abaixo
	// var num2 int = 15
	soma := num1 + num2
	fmt.Println(soma)

	// Atribuição = :=
	var var1 string = "primeiro"
	var2 := "segundo"
	fmt.Println(var1, var2)

	// Relacionais > >= < <= == != retornam um bool
	fmt.Println(1 >= 2)

	// Lógicos && || !
	fmt.Println((1 >= 2) || (1 <= 2))

	//Unários ++ -- += -= /= *= %=
	num := 10
	num++ // não existe ++num --num
	num /= 2
	fmt.Println(num)
}
