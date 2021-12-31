package handlers

import (
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Listar todos los Usuarios")
	/*rw.Header().Set("Content-Type", "application/json")
	// rw.Header().Set("Content-Type", "application/xml")

	db.Connect_DB()
	user, _ := models.ListUsers()
	db.Close_DB()
	output, _ := json.Marshal(user)
	// output, _ := xml.Marshal(user)
	// output, _ := yaml.Marshal(user)
	fmt.Fprintln(rw, string(output))*/

	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, users)
	}
}
func GetUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Obtiene un Usuario")
	/*rw.Header().Set("Content-Type", "application/json")

	// obtener id
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])

	db.Connect_DB()
	user, _ := models.GetUser(userId)
	db.Close_DB()
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))*/
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Crea un Usuario")
	// rw.Header().Set("Content-Type", "application/json")
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}
	// obtener id
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// rw.Header().Set("Content-Type", "application/json")
	// fmt.Fprintln(rw, "Actuliza un Usuario")
	// user := models.User{}
	var userId int64

	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
	/*
		if err := decoder.Decode(&user); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {
			db.Connect_DB()
			user.Save()
			db.Close_DB()
		}
		// obtener id
		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))*/
}
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(rw, "Elimina un Usuario")
	/*	rw.Header().Set("Content-Type", "application/json")

		// obtener id
		vars := mux.Vars(r)

		userId, _ := strconv.Atoi(vars["id"])

		db.Connect_DB()
		user, _ := models.GetUser(userId)
		user.Delete()
		db.Close_DB()
		output, _ := json.Marshal(user)
		fmt.Fprintln(rw, string(output))*/
	if user, err := getUserByRequest(r); err != nil {
		models.SendNotFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
