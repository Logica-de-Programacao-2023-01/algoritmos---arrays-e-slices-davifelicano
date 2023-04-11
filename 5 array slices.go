// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas.
// Solicite ao usuário que informe os valores de cada elemento da matriz.
// Em seguida, imprima a matriz resultante.
package main

import "fmt"

func main() {

	matriz := [3][2]int{}

	for l := 0; l < 3; l++ {
		for c := 0; c < 2; c++ {
			fmt.Printf("digite os valores da matriz [%d][%d]", l, c)
			fmt.Scan(&matriz[l][c])
		}
	}

	fmt.Println("matrizresultante")
	for l := 0; l < 3; l++ {
		for c := 0; c < 2; c++ {
			fmt.Printf("%d", matriz[l][c])

		}
		fmt.Println()

	}
	fmt.Println("Essas são as atribuiçoes das matrizes ")
}
