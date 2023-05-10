// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz.
// Em seguida, imprima a matriz resultante.
package main

import "fmt"

func main() {

	Array := [3][2]int{}
	for l := 0; l < 3; l++ {
		for c := 0; c < 2; c++ {
			fmt.Printf("Informe os valores da linha [%d] coluna [%d]", l, c)
			fmt.Scan(&Array[l][c])
		}
	}
	fmt.Print(Array)
}
