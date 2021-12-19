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
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Funcion Index ")
	c1 := Curso{"Java"}
	c2 := Curso{".Net"}
	c3 := Curso{"Python"}
	c4 := Curso{"Go"}

	cursos := []Curso{c1, c2, c3, c4}
	tmp, err := template.ParseFiles("index.html")

	user := Users{"Giovanny", 36, true, true, cursos}

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
