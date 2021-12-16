package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es ", r.Method)
	fmt.Fprintln(rw, "Hola Mundo de GO Web")
}

func PageNotFound(rw http.ResponseWriter, rq *http.Request) {
	http.NotFound(rw, rq)
}

func PagError(rw http.ResponseWriter, rq *http.Request) {
	http.Error(rw, "Este es un Error", http.StatusConflict)
}
func Saludar(rw http.ResponseWriter, rq *http.Request) {
	fmt.Println(rq.URL.RawQuery)
	fmt.Println(rq.URL.Query())

	name := rq.URL.Query().Get("name")
	age := rq.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola %s tu edad es %s!!", name, age)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PageNotFound)
	mux.HandleFunc("/error", PagError)
	mux.HandleFunc("/saludar", Saludar)

	// Router
	/*	http.HandleFunc("/", Hola)
		http.HandleFunc("/page", PageNotFound)
		http.HandleFunc("/error", PagError)
		http.HandleFunc("/saludar", Saludar)*/

	// crear servidor
	fmt.Println("El servidor esta corriendo en el puerto 1234")
	fmt.Println("Run server: http://localhost:1234/")
	//	sock := http.ListenAndServe("localhost:1234", nil)				// sin mux
	//sock := http.ListenAndServe("localhost:1234", mux) // using mux
	//log.Fatalln(sock)

	// mejorando servidor
	server := &http.Server{
		Addr:    "localhost:1234",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())

}
