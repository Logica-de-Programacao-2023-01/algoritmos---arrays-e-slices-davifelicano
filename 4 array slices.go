// Crie um Slice de inteiros e
// solicite ao usuário que informe o tamanho do Slice e
// os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados
package main

import "fmt"

func main() {
	var x, y, soma int
	Slice := []int{}
	fmt.Print("Informe o tamanho do Slice: ")
	fmt.Scan(&x)

	for {
		fmt.Print("Informe um valor : ")
		fmt.Scan(&y)
		Slice = append(Slice, y)
		if len(Slice) == x {
			break
		}

	}

	for _, v := range Slice {
		soma += v
	}

	fmt.Print(Slice, soma)

}
