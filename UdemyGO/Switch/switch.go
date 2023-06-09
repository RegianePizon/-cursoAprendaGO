package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	default:
		return "Numero incorreto"

	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda"
	default:
		diaDaSemana = "Numero invalido"
	}
	return diaDaSemana
}

func main() {
	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}
