// Crie um Slice de inteiros com o tamanho 5.
// Em seguida, solicite ao usuário que informe um número inteiro.
// Adicione esse número ao Slice apenas se ele não estiver presente.
// Imprima o Slice resultante.
package main

import "fmt"

func main() {
	var x int
	Slice := []int{1, 2, 3, 4, 5}
	fmt.Print("Informe um valor ")
	fmt.Scan(&x)
	for i := 0; i < len(Slice); i++ {
		if Slice[i] == x {
			fmt.Printf("o número %d já esta no array", x)
			break
		} else if Slice[i] != x {
			fmt.Println("O novo slice é : ")
			Slice = append(Slice, x)
			fmt.Print(Slice)
			break
		}
	}

}
