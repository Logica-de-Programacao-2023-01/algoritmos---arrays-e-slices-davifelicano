// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e
// o valor máximo armazenados no Slice.
package main

import "fmt"

func main() {
	Slice := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	minimo := 0
	maximo := 0
	for _, numeros := range Slice {
		if numeros < minimo {
			minimo = numeros
		}
		if numeros > maximo {
			maximo = numeros
		}
	}
	fmt.Printf("O maior %d e o menor número do slice %d", maximo, minimo)
}
