// Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que
// informe dois índices de elementos que deverão ser trocados de posição.
// Imprima o Slice resultante.
package main

import "fmt"

func main() {
	var x int
	var y int
	var trocado string
	slice := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	fmt.Print("Informeu indice para ser trocado: ")
	fmt.Scan(&x)
	fmt.Print("Informe outro indice")
	fmt.Scan(&y)

	if x != y {
		trocado := slice[x]
		slice[x] = slice[y]
		slice[y] = trocado
	}

	fmt.Print(trocado)
}
