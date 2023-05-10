//Crie um Array de inteiros com 5 elementos.
//Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3.
//Imprima o novo Slice.
package main

import "fmt"

func main() {
	Array := [5]int{1, 2, 3, 15, 27}
	Slice := []int{}
	for _, numeros := range Array {
		if numeros%3 == 0 {
			Slice = append(Slice, numeros)
		}
	}
	fmt.Print(Slice)
}
