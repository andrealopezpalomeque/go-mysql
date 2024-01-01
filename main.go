package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)


func main(){
	db.Connect()
	//db.Ping() //verifica si la conexion esta activa

	//fmt.Println(db.ExistsTable("users"))
	
	//db.CreateTable(models.UserSchema, "users")

	user := models.CreateUser("jose", "123456", "jose@gmail.com")
	fmt.Println(user)

	//db.TruncateTable("users")

	db.Close()
}