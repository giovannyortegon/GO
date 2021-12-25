package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect_DB()
	// db.Ping_DB()
	// fmt.Println(db.ExistsTable("users"))
	// db.CreateTable_DB(models.UserSchema, "users")
	// db.TruncateTable("users")
	// user := models.CreateUser("Esteban", "Es1234", "jesteban@email.com")
	user := models.CreateUser("Ana", "Ana1234", "qana@email.com")
	fmt.Println(user)
	// users := models.ListUsers()
	// users := models.GetUser(2)
	// fmt.Println(users)
	// users.UserName = "Juan"
	// users.Password = "J12345"
	// users.Email = "juan@email.com"
	// users.Save()
	// fmt.Println(users)
	// users.Delete()
	fmt.Print(models.ListUsers())

	db.Close_DB()
}
