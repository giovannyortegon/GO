package main

import "fmt"

func calcular(a int, b int) (co int, re int) {
	co = a / b
	re = a % b
	return
}

func main() {
	var a, b , co, re int

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&b)

	co, re = calcular(a, b)

	fmt.Println("El cociente es:", co)
	fmt.Println("El reciduo es:", re)
}
