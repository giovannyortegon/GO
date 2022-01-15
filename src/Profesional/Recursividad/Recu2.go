package main

import "fmt"

func bucInc(n int, max int) int {
	if n >= max {
		fmt.Println(n)
		return n
	} else {
		fmt.Println(n)
		n ++
		return bucInc(n * 2 , max) + 2
	}
}
func main() {

	suma := bucInc(1, 20)

	fmt.Println("La suma es: ", suma)
}
