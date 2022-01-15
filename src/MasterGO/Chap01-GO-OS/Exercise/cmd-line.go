package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main() {

	var sum int64
	var errs error = errors.New("Error: Not float number!")
	arguments := os.Args

	if len(arguments) ==1 {
		fmt.Println("Please give one or more floats\n")
		os.Exit(1)
	}

	for i := 1; i < len(arguments); i++ {
		num, err := strconv.ParseInt(arguments[i], 10, 64)
		if err != nil {
			fmt.Println(errs)
			return
		}
		sum += num
	}

	fmt.Println("La suma es: ", sum)
}
