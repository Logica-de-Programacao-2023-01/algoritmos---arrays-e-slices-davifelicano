package main

import "fmt"

func main() {
	numero := [5]int{1, 2, 3, 4, 5}
	numeros := append(numero[:2], numero[3:]...)
	fmt.Println(numeros)

}
