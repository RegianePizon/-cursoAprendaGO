package main

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	//Pode interagir com qualquer tipo
	generica("String")
	generica(1)
	generica(true)
}
