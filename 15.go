// Crie um Array de floats com 10 elementos.
// Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5.
// Imprima o novo Slice.
package main

import "fmt"

func main() {
	Array := []float64{1.2, 3.5, 8.6, 9.8, 3.3, 5.6, 9.7, 12.4, 15.8, 7.6}
	Slice := []float64{}
	for _, i := range Array {
		if i > 5 {
			Slice = append(Slice, i)
		}
	}
	fmt.Print(Slice)
}
