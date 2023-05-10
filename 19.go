// Faça um programa que leia dois arrays de inteiros de tamanho n
// e gere um terceiro array que seja a soma dos dois primeiros.
package main

import "fmt"

func main() {
	var x int
	fmt.Print("Informe o tamanho do array: ")
	fmt.Scan(&x)
	A1 := make([]int, x)
	A2 := make([]int, x)
	A3 := make([]int, x)
	fmt.Print("digite os elementos do array 1: ")
	for i := 0; i < x; i++ {
		fmt.Print("Posição", i+1, ":")
		fmt.Scan(&A1[i])
	}
	fmt.Print("digite os elementos do array 2: ")
	for i := 0; i < x; i++ {
		fmt.Print("Posição", i+1, ": ")
		fmt.Scan(&A2[i])
	}
	for i := 0; i < x; i++ {
		A3[i] = A1[i] + A2[i]
	}
	fmt.Print("somas dos Arrays", A3)

}
