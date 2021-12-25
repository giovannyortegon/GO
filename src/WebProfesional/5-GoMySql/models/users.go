package models

import (
	"gomysql/db"
)

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

//constructor usuario
func NewUser(username string, password string, email string) *User {
	user := &User{UserName: username, Password: password, Email: email}

	return user
}

// crear usuario e insertar a db
func CreateUser(username string, password string, email string) *User {
	user := NewUser(username, password, email)
	user.Save()

	return user
}

// Crear usuario e insert
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

// Obtener un registro
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}

	return user
}

func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
