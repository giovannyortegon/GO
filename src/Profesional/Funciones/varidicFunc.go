package main

import "fmt"

func sumar(nombre string, numeros ... int) (string, int) {

	var	resultado int

	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)

//	fmt.Print("Los numeros a sumar son: ")
/*	for _, numero := range numeros {
		fmt.Print(numero, " ")
	}

	fmt.Print("\n")
*/
	for i := 0; i < len(numeros); i++ {
		resultado += numeros[i]
	}

	return mensaje, resultado
}

func main() {
	msg, result := sumar("Alice", 4, 5, 2)
	fmt.Printf("%s %d", msg, result)
}
