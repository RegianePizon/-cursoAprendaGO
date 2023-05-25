/*
- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
	## - Utilize mutex para consertar a condição de corrida do exercício anterior.
	*/

	package main 

	import (
		"fmt"
		"runtime"
		"sync"
	)
	
	var contador int
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	const quantidade = 500
	
	func criarGo(i int) {
	
		wg.Add(i)
		for j := 0; j < i; j++ {
			go func() {
				mu.Lock()
				v := contador
				runtime.Gosched()
				v++
				contador = v
				mu.Unlock()
				wg.Done()
				
			}()
		}
	
	}
	
	func main() {
		//contador := 0
		criarGo(quantidade)
		wg.Wait()
		fmt.Println("Valor do Goroutine:", quantidade, "Valor do contador:", contador)
	
	}
	