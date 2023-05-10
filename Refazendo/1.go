// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
package main

import "fmt"

func main() {
	var soma int
	Array := [3]int{1, 2, 3}
	for _, a := range Array {
		soma += a

	}
	fmt.Print(soma)
}
