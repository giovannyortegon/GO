package main

import "fmt"

func main() {
	// BUCLE INFINITO
/*	for {
		fmt.Println("Bucle Infinito")
	}*/

	// while
	numeros := 12345
	c := 0

	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantidad de digitos", c)

	for i := 0; i < 10; i++ {
		fmt.Println("Bucle for como en C")
	}
}
