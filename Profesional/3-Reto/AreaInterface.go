package main

import (
	"fmt"
	"math"
)

const pi float64 = math.Pi

type Area interface {
	area() float64
}

type Cuadrado struct {
	ancho float64
	altura float64
}

type Circulo struct {
	perimetro float64
}

func (cua * Cuadrado) area() float64 {
	return cua.ancho * cua.altura
}

func (cir * Circulo) area() float64 {
	return pi * math.Pow(cir.perimetro, 2.0)
}

func encontrarArea(a Area) {
	fmt.Println(a.area())
}

func main() {

	circulo := Circulo{35.6}
	encontrarArea(&circulo)

	cuadrado := Cuadrado{20.3, 10.6}
	encontrarArea(&cuadrado)
}
