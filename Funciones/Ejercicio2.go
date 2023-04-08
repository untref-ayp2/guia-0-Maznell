// 2. Formar un menú con 4 opciones (como se muestra debajo) y al elegir una de ellas mostrar un cartel diciendo qué opción se eligió o si fue una opción incorrecta.
// - Opción 1
// - Opción 2
// - Opción 3
// - Opción 4

package main

import "fmt"

func main() {
	i := 0
	fmt.Println("Ingrese un valor entre 1 y 4")
	fmt.Scan(&i)
	menu(i)

}

func menu(i int) {
	switch i {
	case 1:
		fmt.Println("Opcion elegida: 1")
	case 2:
		fmt.Println("Opcion elegida: 2")
	case 3:
		fmt.Println("Opcion elegida: 3")
	case 4:
		fmt.Println("Opcion elegida: 4")
	default:
		fmt.Println("Opcion elegida es inválida")
	}
}
