// Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
// Remova todas as ocorrências desse valor no Slice e imprima o resultado.
package main

import "fmt"

func main() {
	var x int
	Slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Slice)
	fmt.Print("Informe um valor para ser retirado do slice")
	fmt.Scan(&x)
	New := Slice[:0]
	for _, numeros := range Slice {
		if numeros != x {
			New = append(New, numeros)
		}
	}
	fmt.Println(New)
}
