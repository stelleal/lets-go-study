package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint16
}

func main() {
	fmt.Println("Arquivo structs.go")

	var u usuario
	u.nome = "Ana"
	u.idade = 20
	fmt.Println(u)

	enderecoPrincipal := endereco{"Rua Essa Aí", 42}

	u2 := usuario{"Davi", 15, enderecoPrincipal}
	fmt.Println(u2)

	u3 := usuario{idade: 29} //caso não tenha todos os valores para a criação do usuario
	fmt.Println(u3)
}
