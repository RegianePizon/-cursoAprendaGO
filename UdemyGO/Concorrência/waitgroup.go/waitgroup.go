package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Programando em GO")
		waitGroup.Done()
	}()
	waitGroup.Wait()

	//	escrever("Ola mundo")
	//	escrever("Programando em GO!")
}
