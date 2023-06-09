package main

import "fmt"

func main() {
	// ARITMETICA
	// +, -, /, *, %

	soma := 1 + 2
	subtração := 1 - 2
	divisão := 10 / 4
	multiplicação := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtração, divisão, multiplicação, restoDivisao)

	//ATRIBUIÇÃO = :=

	var variavel string = "String"
	variavel2 := "String2"

	fmt.Println(variavel, variavel2)

	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println("/n", verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNARIOS

	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

}
