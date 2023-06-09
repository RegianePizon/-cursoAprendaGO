package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ola Mundo")
	}() // parametro que ela precisa

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido --> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)

}
