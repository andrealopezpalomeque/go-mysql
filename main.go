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

	//===========INSERTAR UN REGISTRO================
	user := models.CreateUser("ANDREA", "8989", "ANDREA.LOPEZ@telco.com.ar")
	fmt.Println(user)

	//===========TODOS LOS REGISTROS================
	// users := models.ListUsers()
	// fmt.Println(users)

	//===========UN REGISTRO================
	 //user := models.GetUser(2)
	//fmt.Println(user)

	//===========ACTUALIZAR UN REGISTRO================
	// user.Username = "hipolito CARLOS"
	// user.Save()

	// fmt.Println(models.ListUsers())

	//===========ELIMINAR UN REGISTRO================	

	//user.Delete()

	//db.TruncateTable("users")



	db.Close()
}