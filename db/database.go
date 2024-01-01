package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3307)/database
const url = "root:1234@tcp(localhost:3307)/goweb_db"

var db *sql.DB

func Connect(){
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = connection

}

func Close(){
	db.Close()
}


