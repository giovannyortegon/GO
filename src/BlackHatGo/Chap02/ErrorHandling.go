package main

import "fmt"

type error interface {
	Error() string
}

type MyError string
func (e MyError) Error() string {
	return string(e)
}

func foo() error {
	return error.New("Some Error Occurred")
}

func main() {

	if err := foo(); err != nil {
		fmt.Println("Error!!")
	}
}
