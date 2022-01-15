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

// var tmp = template.Must(template.New("T").ParseFiles("index.html", "base.html"))
var tmp = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func handlError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Println("Ejecutando")
	errorTemplate.Execute(rw, nil)
}

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	err := tmp.ExecuteTemplate(rw, name, data)

	if err != nil {
		//		panic(err)
		//http.Error(rw, "No es posible retornar el template", http.StatusInternalServerError)
		handlError(rw, http.StatusInternalServerError)
	}
}

func Index(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Funcion Index ")

	// tmp := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	user := Users{"Giovanny", 36}

	renderTemplate(rw, "index.html", user)
	// tmp.Execute(rw, user)
}

func Registro(rw http.ResponseWriter, r *http.Request) {

	renderTemplate(rw, "registro.html", nil)
}

func main() {

	staticFile := http.FileServer(http.Dir("static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	fmt.Println("El servidor esta corriendo en el puerto 1234")
	fmt.Println("Run server: http://localhost:1234/")
	server := &http.Server{
		Addr:    "localhost:1234",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

}
