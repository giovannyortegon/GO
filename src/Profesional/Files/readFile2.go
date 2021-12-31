package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("El programa no finalizo de forma correcta!")
		}
	}()

	if file, err := os.Open("hola1.txt"); err != nil{
//		fmt.Println("No fue posible leer")
		panic("No fue posible leer")
	} else {

		defer func() {
			fmt.Println("el archivo ha sido cerrado!")
			file.Close()
		}()

		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)
		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)
	}
}
