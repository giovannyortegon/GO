package main

import "fmt"

func sumaDosNumeros(a int, b int) int {
	return a + b
}

func main() {
	var a, b int

	fmt.Print("Ingrese el primer numero: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scan(&b)
	fmt.Println("La suma es:", sumaDosNumeros(a, b))
}
