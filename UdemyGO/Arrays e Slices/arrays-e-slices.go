package main

import "fmt"

func main() {

	var array1 [5]int
	fmt.Println(array1)

	var array2 [5]string
	array2[0] = "Posição 1"
	fmt.Println(array2)

	array3 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	//Arrays não são muito utilizadas, devido ser engessado

	//SLICES
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 9)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	// Arrays Interno

	slice3 := make([]float32, 10, 15)
	fmt.Println("\n", slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

}
