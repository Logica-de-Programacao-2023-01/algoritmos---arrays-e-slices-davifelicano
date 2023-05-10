// Crie um Array de inteiros com 5 elementos.
// Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3.
// Imprima o novo Slice.
package main

import "fmt"

func main() {
	Array := [5]int{1, 10, 15, 2, 30}
	slice := Array[:0]
	for _, numeros := range Array {
		if numeros%3 == 0 {
			slice = append(slice, numeros)
		}
	}
	fmt.Print(slice)
}
