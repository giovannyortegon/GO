package main

import "fmt"

func main() {
	var vocal string

	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&vocal)

/*	switch {
		case vocal == "a" || vocal == "A":
			fmt.Println(vocal, "es vocal")
		case vocal == "e" || vocal == "E":
			fmt.Println(vocal, "es vocal")
		case vocal == "i" || vocal == "I":
			fmt.Println(vocal, "es vocal")
		case vocal == "o" || vocal == "O":
			fmt.Println(vocal, "es vocal")
		case vocal == "u" || vocal == "U":
			fmt.Println(vocal, "es vocal")
		default:
			fmt.Println(vocal, "No es vocal")
	}*/


	switch vocal {
		case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
			fmt.Println(vocal, "es vocal")
		default:
			fmt.Println(vocal, "No es vocal")
	}
}
