// Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor
// e verifique se esse valor está presente no Array. Imprima o resultado da busca
package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("Informe um valor para consultar o array : ")
	fmt.Scan(&x)

	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(array); i++ {
		if array[i] == x {
			fmt.Printf("O valor %d esta na lista ", x)
			break
		}
		if array[i] != x {
			fmt.Printf("O valor %d não esta na lista ", x)
			break
		}
	}

}
