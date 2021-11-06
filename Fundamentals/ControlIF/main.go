package main

import  "fmt"

func main() {
	var num int

	num = 65
	result:= 10

	fmt.Println("Resultado:", num + result)
	fmt.Println("Resultado:", num != result)

	if num == 5 {
		fmt.Println("IF num", num)
	} else if num >= 50 {
		fmt.Println("ELSE IF num", num)
	} else {
		fmt.Println("ELSE num", num)
	}
}
