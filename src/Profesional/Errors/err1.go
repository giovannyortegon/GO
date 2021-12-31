package main

import (
	"fmt"
	"errors"
)

func division(dividendo int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0")
	} else {
		return dividendo / divisor, nil
	}
}

func main() {

//	result, err := division(4, 0)
	result, err := division(4, 2)

	if err == nil {
		fmt.Println("Resultado: ", result)
	} else {
		fmt.Println(err)
	}
}
