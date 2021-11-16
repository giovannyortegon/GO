package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main() {

	var average float64
	var n_args int
	var errs error = errors.New("Error: Not float number!")

	n_args = len(os.Args) - 1
	arguments := os.Args

	if n_args == 0 {
		fmt.Println("Error: Not arguments!")
		return
	}

	for i := 1; i <= n_args; i++ {
		num, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println(errs)
			return
		}
		average += num
	}

	fmt.Println("The average is: ", average / float64(n_args))
}
