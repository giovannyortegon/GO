package models

import (
	"gorom/db"
)

type User struct {
	Id       int64  `json:"id"`       //`yaml:"id`        //`xml:"id"`
	UserName string `json:"username"` //`yaml:"username"` //`xml:"username"`
	Password string `json:"password"` //`yaml:"password"` //`xml:"password"`
	Email    string `json:"email"`    //`yaml:"email"`    //`xml:"email"`
}

type Users []User

func MigrarUser() {
	db.DataBase.AutoMigrate(User{})
}
