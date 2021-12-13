package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {

	//fmt.Print(servidor)
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + "No esta disponible"
		//fmt.Println(" No esta disponible")
	} else {
		canal <- servidor + "Esta disponible"
		//fmt.Println(" Essta disponible")
	}
}

func main() {

	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.youtube.com/",
		"https://www.facebook.com",
		"https://www.google.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	Fin := time.Since(inicio)
	fmt.Println("tiempo trancurrido:", Fin)
}
