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

	//user := models.CreateUser("hipolito", "8989", "hipolito.coronel@telco.com.ar")
	//fmt.Println(user)

	users := models.ListUsers()
	fmt.Println(users)

	//db.TruncateTable("users")



	db.Close()
}