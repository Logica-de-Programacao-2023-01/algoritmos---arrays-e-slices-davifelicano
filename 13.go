// Crie um Array de inteiros com 7 elementos.
// Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array.
// Imprima o Array resultante.
package main

import "fmt"

func main() {
	var n int
	Array := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Array)
	fmt.Println("Informe um número que sera adicionado ao primeiro e ao ultimo termo do array acima")
	fmt.Scan(&n)
	Array[0] += n
	Array[6] += n
	fmt.Println(Array)
}
