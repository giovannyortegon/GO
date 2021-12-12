package main

import "fmt"

type User struct {
	nombre string
	email string
	activo bool
}

type Student struct {
	user User
	codigo string
}

func main() {
	alex := User {
		nombre: "Alex",
		email: "alex@hotmail.com",
		activo: true,
	}

	roel := User {
		nombre: "Roel",
		email: "roel@hotmail.com",
		activo: true,
	}

	alexStudent := Student{
		user: alex,
		codigo: "RS23132",
	}

	fmt.Println(alex)
	fmt.Println(roel)
	fmt.Println(alexStudent)
	fmt.Println(alexStudent.user.nombre)

}
