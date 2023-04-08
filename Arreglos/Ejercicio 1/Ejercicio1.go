//Escribir una función que reciba un arreglo de enteros como parámetros y devuelva la suma de todos sus elementos

package main

import "fmt"

func main() {

	array := [10]int{9, 8, 5, 6, 99, 8, 5, 4, 2, 1}

	fmt.Println(sumarElementos(array[:]))
}

func sumarElementos(array []int) (b, c int) {

	var j int
	var l int

	//Forma for tradicional
	for i := 1; i < len(array); i += 2 {
		j += array[i] + array[i-1]
	}

	//Forma foreach
	for k := range array {
		l += array[k]
	}

	return j, l
}
