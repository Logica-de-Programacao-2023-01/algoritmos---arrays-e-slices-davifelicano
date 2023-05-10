// Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.
// inicio é a posição do primeiro elemento que desejamos incluir no novo Slice;
// fim é a posição do primeiro elemento que não desejamos incluir no novo Slice.
package main

import "fmt"

func main() {
	Slice := []int{1, 2, 3, 4, 5}
	Slice = append(Slice[:2], Slice[3:]...)
	fmt.Print(Slice)
}
