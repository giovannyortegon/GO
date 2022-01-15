//usr/bin/env go run $0 $@;exit
package main

import "fmt"

func main() {
	var entero int
	var decimal float64
	var texto string
	var booleano bool

	entero = 5
	decimal = 10.4
	texto = "Bienvenido al curso"
	booleano = true

	fmt.Println(entero, decimal, texto, booleano)

	num:=10

	fmt.Println(num)

	// comentarios
	mas_texto:="Hola a todos"
	fmt.Println(mas_texto)
}
