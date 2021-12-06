package main

import "fmt"

func main() {
	var numeros [5] int
	fmt.Println("Array numeros: ", numeros)

	for i := 0; i < len(numeros); i++ {
		numeros[i] = i  +1
	}

	fmt.Println("Array numeros: ", numeros)

	nombres := [3] string {"Paula", "Giovanny", "Ricardo"}

	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	colores := [...] string {
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}

	fmt.Println("length colors: ", len(colores))
	fmt.Println("colores: ", colores)

	for i := 0; i < len(colores); i++ {
		fmt.Println(colores[i])
	}

	monedas := [...] string {
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euros",
	}

	monedas[2] = "Yuan"
	monedas[4] = "Yen"

	fmt.Println(monedas, len(monedas))

	// subarray
	//subMonedas := monedas[0:3]
	subMonedas := monedas[3:]

	fmt.Println(subMonedas, len(subMonedas))

}
