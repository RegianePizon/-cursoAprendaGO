package main

import (
	"fmt"
)

func main() {
	fmt.Println("loopS")
	/*	i := 0
		for i < 10 {
			i++
			fmt.Println(i, "- Incrementando . . .")
			time.Sleep(time.Second)
		}
		fmt.Println(i)

		for j := 0; j < 10; j += 2 {
			fmt.Println("Incrementando j", j)
			time.Sleep(time.Second)
		}
	*/
	nomes := [3]string{"João", "Maria", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for i, letra := range "PALAVRA" {
		fmt.Println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Jose",
		"sobrenome": "Couve",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não dá para usar em Struct
}
