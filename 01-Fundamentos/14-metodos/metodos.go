package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint
}

// u semelhante a palavra chave this
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados de USER %s no db\n", u.nome)
}

func (u usuario) eMaiorDeIdade() (maiorDeIdade bool) {
	maiorDeIdade = u.idade >= 18
	return
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Pedro", 22}
	usuario1.salvar()

	usuario2 := usuario{"Ana", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.eMaiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
