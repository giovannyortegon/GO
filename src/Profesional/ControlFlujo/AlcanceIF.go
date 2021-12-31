package main

import "fmt"

func main() {

	if nombre, edad := "Alex", 26; (nombre == "Alex") {
		fmt.Println("Hola,", nombre);
	} else {
		fmt.Printf("Nombre %s - Edad: %d\n", nombre, edad);
	}

	// Obtener valor de map
	users := make(map[string] string)
	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	if correo, ok := users["Juan"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor.")
	}
}
