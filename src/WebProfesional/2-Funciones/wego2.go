package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una Funcion."
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("In Index")

	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	tmp, err := template.New("index.html").Funcs(funciones).ParseFiles("index.html")

	if err != nil {
		panic(err)
	} else {
		tmp.Execute(rw, nil)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Index)

	fmt.Println("El servidor esta corriendo en el puerto 4040")
	fmt.Println("Run Server: http://localhost:4040")
	server := &http.Server{
		Addr:    "localhost:4040",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

}
