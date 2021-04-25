package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
  )

type User struct {
	gorm.Model `json:"info"`
	Token string `json:"token"`
	Name string `json:"name"`
}

func CreateUser() {
	db := sqlConnect()
	token := "adgahor2ka01"
    name := "Yamada"
    fmt.Println("create user " + name + " with token " + token)
    db.Create(&User{Token: token, Name: name})
    defer db.Close()
}

func GetUsers() []User{
	db := sqlConnect()
    var users []User
    db.Order("created_at asc").Find(&users)
    defer db.Close()

	return users
}