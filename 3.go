// Crie um Array de floats com 4 elementos
// e calcule o produto dos valores armazenados no Array.
package main

import "fmt"

func main() {
	var a float64
	Array := [4]float64{1.5, 1.6, 1.9, 2.5}
	a = 1
	for i := 0; i < len(Array); i++ {
		a *= Array[i]
	}
	fmt.Printf("%.2f", a)
}
