//Escribir una función que indique si un número entero, ingresado por un usuario, es primo (devuelve verdadero o falso)

package estructurasrepetitivas

import "fmt"

func main() {
	var i int
	fmt.Println("Introduzca un número entero:")
	fmt.Scanln(&i)
	fmt.Println("Es Primo ? ", esPrimo(i))

}

func esPrimo(a int) (b bool) {

	primo := true

	for i := 2; i < a; i++ {
		if a%i == 0 {
			primo = false
			break
		}
	}

	return primo
}
