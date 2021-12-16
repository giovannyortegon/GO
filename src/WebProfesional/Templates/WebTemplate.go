package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Users struct {
	UserName string
	Edad     int
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Hola ")
	tmp, err := template.ParseFiles("index.html")

	user := Users{"Giovanny", 36}

	if err != nil {
		panic(err)
	} else {
		tmp.Execute(rw, user)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	fmt.Println("El servidor esta corriendo en el puerto 1234")
	fmt.Println("Run server: http://localhost:1234/")
	server := &http.Server{
		Addr:    "localhost:1234",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

}
