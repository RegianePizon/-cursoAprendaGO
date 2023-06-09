package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1

}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeoInvertido := inverterSinal(numero)
	fmt.Println(numeoInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
