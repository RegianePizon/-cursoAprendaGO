package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Ola Mundo!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
