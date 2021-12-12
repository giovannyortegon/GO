package figuras

import (
	"fmt"
)

type geometria interface {
	area() float64
	perimetro() float64
}

func MedidasGeometricas(g geometria) {
	fmt.Println("Area:", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}
