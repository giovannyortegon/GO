//usr/bin/env go run $0 $@; exit
package main

import "fmt"

func main() {
	suma:= 10 + 5
	fmt.Println("Suma:",suma)
	resta:=suma-5
	fmt.Println("Resta:",resta)
	mult:=resta * 2
	fmt.Println("Multiplicacion:",mult)
	division:=mult/3
	fmt.Println("Division:",division)
	modulo:=division%2
	fmt.Println("Modulo:",modulo)

	suma++
	fmt.Println("Suma:",suma)
	suma--
	fmt.Println("Suma:",suma)

	suma+=10
	fmt.Println("Suma:",suma)

	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println()
	fmt.Println(10 >= 5)
	fmt.Println(10 <= 5)
	fmt.Println(10 ==10)
	fmt.Println(10 !=10)
	fmt.Println(10 > 5 && 10 > 4)
	fmt.Println(10 < 5 || 10 > 4)
}
