package main

import "fmt"

func main() {

	m := 123
//	m := 234		// raise error it can not be asigned again

	i, k := 3, 4
	j, k := 1, 2

	fmt.Println(i, " ", j, " ", k)
}
