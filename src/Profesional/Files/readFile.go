package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("hola.txt")

	if err != nil{
		fmt.Println("No fue posible leer")
	} else {
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}

	file.Close()
}
