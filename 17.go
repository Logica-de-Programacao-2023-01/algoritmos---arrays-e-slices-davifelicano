//Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.
package main

import "fmt"

func main() {
	Array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0

	for i := 0; i < len(Array); i += 2 {
		soma += Array[i]
	}

	fmt.Print(soma)
}
