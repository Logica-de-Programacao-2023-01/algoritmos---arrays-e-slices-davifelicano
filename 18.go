// Escreva um programa que leia um número inteiro positivo n e
// gere um array com os n primeiros números primos.
package main

import "fmt"

func main() {
	var x int
	fmt.Print("informe um número: ")
	fmt.Scan(&x)
	if x < 0 {
		fmt.Print("Error")
	}
	w := 2
	p := make([]int, 0, x)
	for len(p) < x {
		ép := true
		for k := 2; k*k <= w; k++ {
			if w%k == 0 {
				ép = false
			}
		}
		if ép {
			p = append(p, w)
		}
		w++
	}
	fmt.Print(p)
}
