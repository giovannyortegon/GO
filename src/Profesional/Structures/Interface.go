package main

import "fmt"

type Animal interface {
	mover() string
}

type Perro struct {}
type Pez struct {}
type Pajaro struct {}

func (* Perro) mover() string {
	return "Soy perro y camino"
}

func (* Pez) mover() string {
	return "Soy pez y nado"
}

func (* Pajaro) mover() string {
	return "Soy pajaro y vuelo"
}

func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}

func main() {

	perro := Perro{}
	moverAnimal(&perro)

	pez := Pez{}
	moverAnimal(&pez)

	pajaro := Pajaro{}
	moverAnimal(&pajaro)
}
