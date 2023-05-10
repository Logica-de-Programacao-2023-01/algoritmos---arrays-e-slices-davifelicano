// Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5.
// Remova o terceiro elemento e imprima o Slice resultante.
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice[:2], slice[3:]...)
	fmt.Print(slice)
}
