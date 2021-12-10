package main

import (
	"fmt"
	"strings"
)

func reversar(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertido := make([] string, 0)

	for i := len(arrayCadena) -1; i >= 0; i-- {
		arrayInvertido = append(arrayInvertido, arrayCadena[i])
	}

	return strings.Join(arrayInvertido, "")
}

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	//palabra = strings.Replace(palabra, " ", "*", 1)
	palabra = strings.ReplaceAll(palabra, " ", "")

	palabraInvertida := reversar(palabra)

	return palabra == palabraInvertida
}

func main() {
	if esPalindromo("Luz Azul") {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}
