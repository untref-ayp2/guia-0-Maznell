//Escribir una función que recibe dos enteros: a y b y devuelve el producto. Para el cálculo, utilizar sumas sucesivas

package main

import "fmt"

func main() {

	fmt.Println(producto(4, 3))
	fmt.Println(producto(4, 10))
	fmt.Println(producto(8, 7))
	fmt.Println(producto(5, 9))

}

func producto(a, b int) (c int) {

	j := a
	for i := 1; i < b; i++ {
		j += a
	}
	return j
}
