package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	instituicao string
	curso       string
}

// declarando assim é possível acessar estudante.nome, diferente de estudante.pessoa.nome

func main() {
	p1 := pessoa{"Pedro", "Souza", 25}
	e1 := estudante{p1, "UFF", "Letras"}

	e2 := estudante{pessoa{"José", "Santos", 20}, "IF", "Matematica"}

	fmt.Println(e1.nome)
	fmt.Println(e2)
}
