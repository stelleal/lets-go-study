package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}
func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	// Defer -> respeita o escopo. Caso a func tenha retorno, o que tiver com defer vai até o último momento ANTES do retorno
	defer fmt.Println("Média calculada! Resultado será retornado...")
	fmt.Println("Verificando se aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	// Adia a execução do código com defer
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 2))
}
