package main

import "fmt"

func main() {
	numeros := [3]int{1, 2, 3}
	soma := 0
	for _, numeros := range numeros {
		soma += numeros

	}
	fmt.Print("A soma dos arrays é ", soma)

}
