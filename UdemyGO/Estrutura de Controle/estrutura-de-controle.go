package main

import "fmt"

func main() {

	fmt.Println("Estrutura de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O numero é maior que 0")
	} else {
		fmt.Println("O numero é menor que 0")
	}

}
