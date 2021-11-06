package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad int
}


func main() {
	p1 := persona{"Bill", "Gates", 65}
	p2 := persona{"Giovanny", "Ortegon", 36}

	fmt.Println("Persona 1:", p1)
	fmt.Println("Persona 2:", p2)
}
