package main

import (
	"fmt"
	"sync"
)

/*- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

var wg sync.WaitGroup //WaitGroup serve para esperar que uma coleção de goroutines termine sua execução

func main() {

	wg.Add(2) //"Quantas goroutines?"

	go Goroutine1()
	go Goroutine2()

	wg.Wait() //"Espera todo mundo terminar."

}

func Goroutine1() {
	fmt.Println("Eu sou o Goroutine numero 1")

	wg.Done() //"Deu!"
}

func Goroutine2() {
	fmt.Println("Eu sou o Goroutine numero 2")

	wg.Done() //"Deu!"
}
