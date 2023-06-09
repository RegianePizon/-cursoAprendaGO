package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA

	var v int
	var ponteiro *int

	v = 100
	ponteiro = &v

	fmt.Println(v, *ponteiro) // desreferencição

}
