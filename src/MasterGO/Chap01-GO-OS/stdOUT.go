package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	argument := os.Args

	if len(argument) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = argument[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
