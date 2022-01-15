package main

import "fmt"

func main() {
	numeros := [] int {1, 2, 3, 4}
	fmt.Println("slicing: ", numeros, len(numeros))
	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	fmt.Println("slicing: ", numeros, len(numeros))
	numeros = numeros[:3]
	fmt.Println("slicing: ", numeros, len(numeros))

	meses := [] string {"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len: %d, capacidad: %d, referencia: %p\n",
				len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %d, capacidad: %d, referencia: %p\n",
				len(meses), cap(meses), meses)
}
