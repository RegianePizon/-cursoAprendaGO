package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	fmt.Println("Herança")
	p := pessoa{"Jose", 42}
	fmt.Println(p)

	e := estudante{p, "Engenharia"}
	fmt.Println(e.nome)
}
