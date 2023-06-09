package main

import "fmt"

// Função INIT execulta primeiro que o main
func init() {
	fmt.Println("Função init")
}

func main() {
	fmt.Println("Função main")
}
