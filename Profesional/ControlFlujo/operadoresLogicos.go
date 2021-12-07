package main

import (
	"fmt"
)

func main() {

	//NOT
	fmt.Println("Operador NOT")
	fmt.Println(!true)

	// AND
	fmt.Println("Operador AND")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	// OR
	fmt.Println("Operador OR")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
}
