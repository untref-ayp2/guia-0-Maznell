//Escribir una función que recibe un número entero n, no negativo, y devuelve su factorial

package estructurasrepetitivas

import "fmt"

func main() {
	var i int
	fmt.Println("Introduzca un número natural:")
	fmt.Scanln(&i)
	fmt.Println("Su factorial es: ", factorial(i))
}

func factorial(i int) (l int) {

	k := 1
	for j := 1; j <= i; j++ {
		k *= j
	}

	return k
}
