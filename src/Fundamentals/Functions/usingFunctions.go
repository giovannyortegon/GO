package main

import "fmt"

func main() {

	saludar("Giovanny")
	fmt.Println("Ejecucion main")
	mi_funcion()

	fmt.Println(suma(10, 15))

	resultado:= suma(3, 5)

	fmt.Println(resultado)
}

func mi_funcion() {
	num:=1

	for num != 10 {
		fmt.Println("Ejecuntando Funcion")
		num++
	}
}

func saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func suma(numero1 int, numero2 int) int {
	suma:= numero1 + numero2

	return suma
}
