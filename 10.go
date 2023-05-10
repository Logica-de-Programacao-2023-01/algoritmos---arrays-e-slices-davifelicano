//Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
package main

import "fmt"

func main() {
	Slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	maior := 0
	menor := 0
	for _, numeros := range Slice {
		if numeros < menor {
			menor = numeros
		}
		if numeros > maior {
			maior = numeros
		}
	
	}
	fmt.Printf("O maior número é %d e o menor número é %d ", maior, menor)
}
