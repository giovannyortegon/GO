package main

import (
	"fmt"
	"packages/figuras"
	"packages/models"

	"github.com/donvito/hellomod"
)

func main() {

	cir := figuras.Circulo{3.5}
	figuras.MedidasGeometricas(&cir)

	cua := figuras.Cuadrado{5.6}
	figuras.MedidasGeometricas(&cua)
	hellomod.SayHello()

	per1 := models.Persona{}
	per1.Constructor("Giovanny", 37)
	per1.Imprimir()
	fmt.Println(per1.GetNombre())
	per1.SetNombre("Alex")
	fmt.Println(per1.GetNombre())
	per1.Imprimir()
	per1.SetEdad(45)
	per1.Imprimir()
}
