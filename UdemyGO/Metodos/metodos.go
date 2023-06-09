package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuario %s no banco de dados\n", u.nome)
}
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniverssario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Davi", 17}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	//	fmt.Println("É maior de idade:", maiorDeIdade)

	if maiorDeIdade == true {
		fmt.Println("É maior de idade")
	} else {
		fmt.Println("É menor de Idade")
	}
	//	fmt.Println("É maior de idade:", maiorDeIdade)

	usuario2.fazerAniverssario()
	fmt.Println("Agora completou:", usuario2.idade)
	//	fmt.Println("É maior de idade:", usuario2.maiorDeIdade())

}
