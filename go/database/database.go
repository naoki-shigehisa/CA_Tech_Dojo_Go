package database

import (
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
  )

type User struct {
	gorm.Model `json:"info"`
	Token string `json:"token"`
	Name string `json:"name"`
}

func Create() {
	db := SqlConnect()
	token := "adgahor2ka01"
    name := "Yamada"
    fmt.Println("create user " + name + " with token " + token)
    db.Create(&User{Token: token, Name: name})
    defer db.Close()
}

func Get() []User{
	db := SqlConnect()
    var users []User
    db.Order("created_at asc").Find(&users)
    defer db.Close()

	return users
}

func SqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"
  
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
  
	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
	  for {
		if err == nil {
		  fmt.Println("")
		  break
		}
		fmt.Print(".")
		time.Sleep(time.Second)
		count++
		if count > 180 {
		  fmt.Println("")
		  fmt.Println("DB接続失敗")
		  panic(err)
		}
		db, err = gorm.Open(DBMS, CONNECT)
	  }
	}
	fmt.Println("DB接続成功")
  
	return db
  }