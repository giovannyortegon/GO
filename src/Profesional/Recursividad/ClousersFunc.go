package main

import (
	"fmt"
	"strings"
)

func repetir(i int) func(cadena string) string {
	return func(cadena string) string {
//		return fmt.Sprintf("Hola %s\n", cadena)
		return strings.Repeat(cadena, i)
	}
}

func main() {

	rep := repetir(3)
	fmt.Println(rep("Giovanny\t"))
}
