package main

import "fmt"

type Persona struct {
	nombre string
	edad int
}

func (p * Persona) setNombre(nNom string) {
	p.nombre = nNom
}

func (p * Persona) getNombre() string {
	return p.nombre
}

func (p * Persona) imprimir() {
	fmt.Printf("Nombre: %s Edad: %d\n", p.nombre, p.edad)
}

// Herencia
type Empleado struct {
	Persona
}

func main() {

	p1 := Persona {"Alex", 26}
	p1.imprimir()
//	fmt.Println("Estructura 1: ", p1)
	p1.setNombre("Roel")
	fmt.Println(p1.getNombre())
	p1.imprimir()
//	fmt.Println("Estructura 1: ", p1)
	p2 := Persona {
		nombre: "Juan",
		edad: 32,
	}
	p2.imprimir()
//	fmt.Println("Estructura 1: ", p2)
	emp1 := Empleado{}

	emp1.nombre = "Carlos"
	emp1.edad = 40
	emp1.imprimir()

}
