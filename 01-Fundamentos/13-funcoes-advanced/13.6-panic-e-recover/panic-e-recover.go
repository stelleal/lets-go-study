package main

import "fmt"

func recuperarExecucao() {
	fmt.Println("Recuperando Execução...")
	// se a aplicação não estiver no panic, r == nil
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	fmt.Println("Verificando se aluno está aprovado")
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// panic não é um erro/err/error
	panic("A MÉDIA É EXATAMENTE 6!!")
}

func main() {
	// caso media == 6 e com o recover ativo, o retorno da func será o valor padrão do tipo do retorno, neste caso bool -> false
	fmt.Println(alunoEstaAprovado(6, 6))

	fmt.Println("Final da Aplicação")
}
