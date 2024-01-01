package main

import "gomysql/db"


func main(){
	db.Connect()
	db.Ping() //verifica si la conexion esta activa
	db.Close()
}