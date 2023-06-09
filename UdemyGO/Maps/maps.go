package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println("Nome:", usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Jose",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println("Nome:", usuario2["nome"])
	delete(usuario2, "nome")
	fmt.Println("Sem campo nome:\n", usuario2)

	usuario2["faculdade"] = map[string]string{
		"nome": "Anhanguera",
	}
	fmt.Println("Add o campo Faculdade\n", usuario2)
}
