package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	// Concorrencia "= Paralelismo"
	go escrever("Olá Mundo") // goroutine
	escrever("Programando em GO")

}
