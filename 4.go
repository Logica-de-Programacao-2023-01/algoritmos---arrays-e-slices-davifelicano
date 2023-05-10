// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice
// e os valores dos elementos.
// Em seguida, imprima o Slice e a soma dos valores armazenados.
package main

import "fmt"

func main() {
	var t int
	fmt.Print("Informe o tamanho do slice ")
	fmt.Scan(&t)
	slice := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Print("Informe o valores do slice")
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice)
	a := 0
	for _, numeros := range slice {
		a += numeros
	}
	fmt.Println(a)
}
