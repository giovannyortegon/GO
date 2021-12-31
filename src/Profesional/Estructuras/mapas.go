package main

import "fmt"

func main() {
	dias := make(map[int] string)

	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"

	fmt.Println("Mapa:", dias)

	delete(dias, 7)
	fmt.Println("Mapa:", dias)

	// Otra clase de Mapa

	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{1, 2, 3}
	estudiantes["Roel"] = []int{4, 5, 6}

	fmt.Println(estudiantes["Alex"][2])
}
