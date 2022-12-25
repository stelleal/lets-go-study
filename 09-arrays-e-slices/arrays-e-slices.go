package main

import "fmt"

func main() {
	// Arrays

	var array1 [5]int
	array1[0] = 333
	fmt.Println(array1)

	array2 := [3]int{}
	fmt.Println(array2)

	array3 := [5]string{"Primeiro", "Segundo", "Terceiro", "Quarto", "Quinto"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4}
	fmt.Println(array4)

	// Slices

	slice1 := []int{11, 22, 33, 44}
	slice1 = append(slice1, 55)
	fmt.Println(slice1)

	slice2 := array3[1:3] // [inclusivo : exclusivo]
	fmt.Println(slice2)
	// Caso o Array referência mude, o slice também mudará
	array3[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Arrays Internos
	
	fmt.Println("----------------")
	slice3 := make([]float32, 10, 11) // criando o array referência para o slice3
	fmt.Println(slice3)
	fmt.Println(len(slice3), cap(slice3)) //length, capacity

	slice3 = append(slice3, 1.1)
	fmt.Println(len(slice3), cap(slice3))

	slice3 = append(slice3, 2.2) // cap torna-se o dobro de length
	fmt.Println(len(slice3), cap(slice3))

	slice4 := make([]float32, 5) // valor de cap é opcional passar
	fmt.Println(len(slice4), cap(slice4))
	slice4 = append(slice4, 6.6)
	fmt.Println(len(slice4), cap(slice4))
}
