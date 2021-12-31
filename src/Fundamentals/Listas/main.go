package main

import "fmt"

func main() {

	frutas:= [] string {"Pera", "Manzana", "Naranja"}

	fmt.Println(frutas[0])

	frutas = append(frutas, "Melocoton", "Melon")

	for i:=0; i < len(frutas); i++ {
		fmt.Println(frutas[i])
	}

	for i:=0; i < len(frutas); i++ {
		if frutas[i] == "Limo" {
			fmt.Println(frutas[i] + " existe.")
			break
		} else {
			fmt.Println("No se encuentra")
		}
	}

}
