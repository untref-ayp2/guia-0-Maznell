//Escribir una función que reciba una lista de enteros y devuelva el menor y el mayor de la lista

package main

import "fmt"

func main() {
	var arreglo = [10]int{8, 1, 33, 9, 8, 5, 10, 23, 8, 4}
	buscar(arreglo[:])
}

func buscar(lista []int) {
	var (
		min int = lista[0]
		max int = lista[0]
	)

	for i := 1; i < len(lista); i++ {
		if lista[i] < min {
			min = lista[i]
		}
	}
	for i := 1; i < len(lista); i++ {
		if lista[i] > max {
			max = lista[i]
		}
	}
	fmt.Println("El minimo es: ", min)
	fmt.Println("El maximo es: ", max)

}
