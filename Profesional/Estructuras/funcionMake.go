package main

import "fmt"

func main() {
	numeros := make([]int, 0, 3)
	fmt.Println("Slice", numeros, len(numeros), cap(numeros))
}
