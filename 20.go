// Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.
package main

import "fmt"

func main() {
	var x int
	A1 := make([]int, x)
	fmt.Print("Informe o tamanho do array:")
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Print("número", i+1)
		fmt.Scan(&A1[i])
	}
	ordenado := true
	for i := 1; i < x; i++ {
		if A1[i] < A1[i-1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println(" ordenado em ordem crescente.")
	} else {
		fmt.Println(" não está ordenado em ordem crescente.")
	}
}
