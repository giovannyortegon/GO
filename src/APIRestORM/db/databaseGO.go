package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "xxxx:xxxx@tcp(127.0.0.1:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var DataBase = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion Exitosa.")
		return db
	}
}()

//username:password@tcp(localhost:3306)/database
const url = "xxxx:xxxx@tcp(localhost:3306)/goweb_db"

var db *sql.DB

func Connect_DB() {

	conn, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion Exitosa")

	db = conn
}
func Close_DB() {
	db.Close()
}

func Ping_DB() {

	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)

	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows.Next()
}

func CreateTable_DB(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("[!] La tabla ya Existe !!")
	}
}

func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect_DB()
	result, err := db.Exec(query, args...)
	Close_DB()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect_DB()
	rows, err := db.Query(query, args...)
	Close_DB()

	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
