// Crie um Array bidimensional de inteiros com 2 linhas e
// 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e
// outro de coluna e imprima o valor armazenado nessa posição da matriz.
package main

import "fmt"

func main() {
	var l, c int
	Array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(Array)
	fmt.Print("Informe um índice de uma linha")
	fmt.Scan(&l)
	fmt.Print("Informe um índice de uma coluna ")
	fmt.Scan(&c)
	x := Array[l][c]
	fmt.Print(x)
}
