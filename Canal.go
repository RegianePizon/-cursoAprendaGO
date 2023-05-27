//  - Chans par, ímpar, quit
//  - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//  - Func receive é um select entre os três canais, encerra no quit

package main

import "fmt"

func send(par, impar chan int, quit chan bool){
total:= 50
for i := 0 ; i < total; i++{
	if i%2 == 0{
par <-i
	}else{
		impar <-i
	}
}
close(par)
close(impar)
quit <- true
}

func receive(par,impar chan int, quit chan bool){
for {
select{
case r := <-par:
	fmt.Println("Par recebeu:\t", r)
case r := <-impar:
	fmt.Println("Impar recebeu:\t", r)
case r, ok := <-quit:
	if !ok{
	fmt.Println("Deu ruim!!")	
	} else{
		fmt.Println("Concluido! Recebemos:", r)	
}
return
}


}

}


func main(){

par:= make(chan int)
impar:=make(chan int)
quit:=make(chan bool)

go send(par, impar,quit)
receive(par,impar,quit)

}