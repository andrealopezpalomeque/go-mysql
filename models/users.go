package models

import (
	"gomysql/db"
)


type User struct {
	Id int64
	Username string
	Password string
	Email string
}


const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`


//constructor de un nuevo usuario
func NewUser(username string, password string, email string) *User{
	user := &User{
		Username: username,
		Password: password,
		Email: email,
	}
	return user
}

//crear usuario e insertar en la base de datos
func CreateUser(username, password, email string) *User{
	user := NewUser(username, password, email)
	user.insert()
	return user
}

//insertar registro
func(user *User) insert(){
	sql := "INSERT users SET username=?, password=?, email=?"

	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)

	user.Id, _ = result.LastInsertId() //devuelve el id del ultimo registro insertado

}