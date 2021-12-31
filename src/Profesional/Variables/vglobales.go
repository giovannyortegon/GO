package main

import "fmt"

var mensaje string

func func1() {
	mensaje = "func 1"
	fmt.Println(mensaje)
}

func func2() {
	mensaje = "func 2"
	fmt.Println(mensaje)
}

func func3() {
	mensaje = "func 3"
	fmt.Println(mensaje)
}

func main() {
	const PI = 3.1415

	fmt.Println(PI)

	mensaje = "Func Main"
	defer func1()
	func2()
	func3()
}
