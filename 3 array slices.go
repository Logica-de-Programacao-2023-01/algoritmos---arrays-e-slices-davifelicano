// Crie um Array de floats com 4 elementos
// e calcule o produto dos valores armazenados no Array.
package main

import "fmt"

func main() {
	Array := [4]float64{2, 3, 4, 5}
	a := Array[0] * Array[1] * Array[2] * Array[3]
	fmt.Print("Multiplicação entre os Array ", a)
}
