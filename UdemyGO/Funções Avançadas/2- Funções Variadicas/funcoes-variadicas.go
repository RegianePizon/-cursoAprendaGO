package main

import "fmt"

// Variadico por obrigação tem que ser o Ultimo
func soma(n ...int) int {
	total := 0
	for _, numero := range n {
		total += numero
	}
	return total
}

func escrever(texto string, n ...int) {
	for _, numero := range n {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 200, 422)
	fmt.Println(totalDaSoma)

	escrever("Ola Mundo", 1, 2, 3, 4, 5)
}
