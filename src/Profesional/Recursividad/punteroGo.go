package main

import "fmt"

func passPointer(cadena * string) {
	fmt.Println(cadena)
	fmt.Println(*cadena)
}

func main() {
	var p * string

	cadena := "Hola a todos"

	p = &cadena

	fmt.Println("Referencia ", p)
	fmt.Println("Referencia ", &cadena)
	fmt.Println("Referencia ", * p)

	passPointer(&cadena)
	passPointer(p)
}
