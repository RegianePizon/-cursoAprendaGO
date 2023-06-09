package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola Mundo", canal)

	fmt.Println("Depois da função ser executada")
	/*
		for {
			mensagem, aberto := <-canal
			if !aberto {
				break
			}
			fmt.Println(mensagem)
		}
	*/

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do Programa!!")
	//	mensagem := <-canal
	//	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)

}
