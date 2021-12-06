package main

import "fmt"

func Saludar (nombre string) {
	fmt.Println("Hola", nombre)
}

func concatenar(a string, b string) string {
	return a + b
}

func sumar(a int, b int) int {
	return a + b
}

func main() {
	Saludar("Giovanny")
	fmt.Println("La suma es:",sumar(5, 6))
	fmt.Println(concatenar("Clean", "Up"))
}
