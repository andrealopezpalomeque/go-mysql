package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3307)/database
const url = "root@tcp(localhost:3307)/goweb_db"
//guarda la conexion
var db *sql.DB

//realizar la conexion
func Connect(){
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = connection

}
//cerrar la conexion
func Close(){
	db.Close()
}

//verificar la conexion
func Ping(){
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Pong")
}


