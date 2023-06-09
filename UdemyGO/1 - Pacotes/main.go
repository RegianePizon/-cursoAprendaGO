package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Sou o Inicio")
	auxiliar.Escrever()

	//validar email
	erro := checkmail.ValidateFormat("tyane.ps@gmail.com")
	fmt.Println(erro)
}
