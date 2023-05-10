// Crie um Array de inteiros com 10 elementos.
// Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array.
// Imprima o resultado da busca.
package main

import "fmt"

func main() {
	var x int
	Array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(Array)
	fmt.Print("Informe um número")
	fmt.Scan(&x)
	for _, numeros := range Array {
		if numeros == x {
			fmt.Printf("O número %d esta contido no Array", x)
			break
		} else if numeros != x {
			fmt.Printf("O número %d não esta contido no Array", x)
			break
		}
	}
}
