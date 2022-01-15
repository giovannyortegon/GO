package main

import "fmt"

func main() {
	fmt.Println("Operadores relacionales")
	var r bool

	a := 10
	b := 20

	r = true
	fmt.Println(r)

	r = a == b
	fmt.Printf("%d es igual que %d? %t\n", a, b, r)

	r = a != b
	fmt.Printf("%d es distinto que %d? %t\n", a, b, r)

	r = a > b
	fmt.Printf("%d es mayor que %d? %t\n", a, b, r)

	r = a < b
	fmt.Printf("%d es menos que %d? %t\n", a, b, r)

	r = a >= b
	fmt.Printf("%d es mayor o igual que %d? %t\n", a, b, r)

	r = a <= b
	fmt.Printf("%d en menor o igual que %d? %t\n", a, b, r)
}
