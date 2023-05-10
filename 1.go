// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
package main

import "fmt"

func main() {
	Array := [3]int{1, 2, 3}
	a := 0
	for _, numeros := range Array {
		a = a + numeros
	}
	fmt.Print(a)
}
