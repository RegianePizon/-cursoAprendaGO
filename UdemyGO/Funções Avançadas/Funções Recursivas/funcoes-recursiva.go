package main

import "fmt"

func fribonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fribonacci(posicao-2) + fribonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")
	// 1 1 2 3 5 8 13
	posicao := uint(12)
	for i := uint(0); i < posicao; i++ {
		fmt.Println(fribonacci(i))
	}

}
