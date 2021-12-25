package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "skill:xxxxx@tcp(localhost:3306)/goweb_db"

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
	result, err := db.Exec(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}
