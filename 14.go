// Crie um Slice de inteiros com tamanho 8 e
// solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição.
// Imprima o Slice resultante.
package main

import "fmt"

func main() {
	var d1, d2 int
	var troca string
	Slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Slice)
	fmt.Println("Informe um indices  : ")
	fmt.Scan(&d1)
	fmt.Println("Informe um indices  : ")
	fmt.Scan(&d2)
	if d1 != d2 {
		troca = Slice[d1]
		Slice[d1] = Slice[d2]
		Slice[d2] = troca
	}
	fmt.Println(troca)
}
