// Crie um Array de floats com 6 elementos.
// Solicite ao usuário que informe um número que será adicionado em todas as posições do Array.
// Imprima o Array resultante.
package main

import "fmt"

func main() {
	var x float64

	Array := [6]float64{1.2, 1.5, 2.5, 7.8, 9.1, 1.3}
	fmt.Println(Array)

	fmt.Print("Informe um valor para ser adicionado a cada elemteno do Array acima")
	fmt.Scan(&x)

	for i := 0; i < len(Array); i++ {
		Array[i] += x
	}

	fmt.Print(Array)
}
