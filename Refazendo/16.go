//Crie um Array de inteiros com 10 elementos.
//Crie um novo Slice que contenha apenas os elementos pares do Array original.
//Imprima o novo Slice.
package main

import "fmt"

func main() {
	Array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Slice := []int{}
	for _, numeros := range Array {
		if numeros%2 == 0 {
			Slice = append(Slice, numeros)
		}
	}
	fmt.Print(Slice)
}
