package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad int
}

func (p persona) saludar (saludo string) {
	fmt.Println(saludo + ", " + p.nombre + " " + p.apellido)
}
func (p persona) cumple () int {
	return p.edad + 1
}


func main() {
	var edad int

	p1 := persona{"Bill", "Gates", 65}
	p2 := persona{"Giovanny", "Ortegon", 36}

	fmt.Println("Persona 1:", p1)
	fmt.Println("Persona 2:", p2)

	p1.saludar("Hello")
	edad = p1.cumple()
	fmt.Println(edad)

	p2.saludar("Hola")
	edad = p2.cumple()
	fmt.Println(edad)
}
