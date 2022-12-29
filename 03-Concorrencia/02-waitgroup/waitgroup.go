package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//				tipo ↴
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // numero de go routines dentro do grupo de espera

	go func() {
		escrever("Olá Gopher! ʕ •ᴥ•ʔノ")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em Go! ʕ •ᴥ•ʔっ💻")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	fmt.Println("-------")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
