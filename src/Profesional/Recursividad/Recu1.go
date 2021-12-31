package main

import "fmt"

func bucInc(n int) int {
	if n == 0 {
		return 0
	} else {
		fmt.Println(n)
		return bucInc(n - 1) + n
	}
}
func main() {

	suma := bucInc(10)

	fmt.Println("La suma es: ", suma)
}
