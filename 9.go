// Crie um Array de floats com 6 elementos.
// Solicite ao usuário que informe um número que será adicionado em todas as posições do Array.
// Imprima o Array resultante.
package main

import "fmt"

func main() {
	var x float64
	Array := [6]float64{1.3, 1.5, 1.9, 4.3, 5.5, 6.7}
	fmt.Print(Array)
	fmt.Print("Informe um número que sera adicinonado a cada posição do array acima")
	fmt.Scan(&x)
	for i := 0; i < len(Array); i++ {
		Array[i] += x
	}
	fmt.Print(Array)
}
