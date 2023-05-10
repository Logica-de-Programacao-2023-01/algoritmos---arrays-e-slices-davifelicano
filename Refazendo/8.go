// Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
// Remova todas as ocorrências desse valor no Slice e imprima o resultado.
package main

import "fmt"

func main() {
	var x int
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	fmt.Print("Informe um valor para ser retirado do slice:")
	fmt.Scan(&x)
	Nslice := slice[:0]
	for _, n := range slice {
		if n != x {
			Nslice = append(Nslice, n)
		}
	}
	fmt.Print(Nslice)
}
