package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Luz",
	}
	fmt.Println(usuario["nome"])

	estudante := map[string]map[string]string{
		"dados": {
			"nome":  "Ana",
			"email": "ana@email.com",
		},
		"universidade": {
			"nome": "AAP",
		},
	}
	fmt.Println(estudante["dados"]["nome"])

	fmt.Println("----------")

	delete(estudante, "universidade")
	fmt.Println("delete de universidade", estudante)

	fmt.Println("----------")

	estudante["projetos"] = map[string]string{
		"artes": "pintura a oleo",
	}

	fmt.Println("add de projetos", estudante)
	fmt.Println(estudante["projetos"]["artes"])

}
