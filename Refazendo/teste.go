package main

import "fmt"

func main() {
	array := [10]float64{5.5, 1.2, 1.5, 1.8, 2.8, 3.2, 4.3, 6.5, 7.0, 2.9}
	fmt.Print(array)
	slice := []float64{}
	fmt.Print(slice)
	for _, numeros := range array {
		if numeros > 5 {
			slice = append(slice, numeros)
		}
	}
	fmt.Print(slice)
}
