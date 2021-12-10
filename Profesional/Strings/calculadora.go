package main

import (
	"fmt"
	"strings"
	"strconv"
)

func sumar(expresion string) int {
	var result int

	valores := strings.Split(expresion, "+")

	for _, valor := range valores {
		num, err := strconv.Atoi(valor)

		if err != nil {
			continue
		}
		result += num
	}

	return (result)
}

func main() {
	var expresion string
	var result int

	fmt.Print("$ ")
	fmt.Scanln(&expresion)
	result = sumar(expresion)
	fmt.Println(result)
}
