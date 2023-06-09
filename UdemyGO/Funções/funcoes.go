package main

import "fmt"

func matematica(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtração := n1 - n2
	return soma, subtração
}

func main() {
	//soma, subtração := matematica(10, 20)
	//fmt.Println("Resultado da soma:", soma, "Resultado da subtração:", subtração)
	_, subtração := matematica(10, 20)
	fmt.Println("Resultado apenas da subtração:", subtração)

	var f = func(txt string) string {
		fmt.Println("\nDentro da função:", txt)
		return txt
	}
	Resultado := f("Outra forma de função")
	fmt.Println("Fora da função:", Resultado)
}
