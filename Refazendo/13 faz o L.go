// Crie um Array de inteiros com 7 elementos.
// Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array.
// Imprima o Array resultante.
package main

import "fmt"

func main() {
	var x int
	Array := []int{}
	for len(Array) < 7 {
		fmt.Print("Informe um número: ")
		fmt.Scan(&x)
		Array = append(Array, x)
	}
	fmt.Print(Array)
}
