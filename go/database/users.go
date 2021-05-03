package database

import (
	"fmt"
	"crypto/rand"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
  )

type User struct {
	gorm.Model `json:"info"`
	Token string `json:"token"`
	Name string `json:"name"`
}

func makeRandomStr(digit uint32) (string, error) {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // 乱数を生成
    b := make([]byte, digit)
    if _, err := rand.Read(b); err != nil {
        return "", errors.New("unexpected error...")
    }

    // letters からランダムに取り出して文字列を生成
    var result string
    for _, v := range b {
        // index が letters の長さに収まるように調整
        result += string(letters[int(v)%len(letters)])
    }
    return result, nil
}

func CreateUser(name string) string{
	token, _ := makeRandomStr(10)
	db := sqlConnect()
    fmt.Println("create user " + name + " with token " + token)
    db.Create(&User{Token: token, Name: name})
    defer db.Close()

	return token
}

func UpdateUser(token string, name string){
	db := sqlConnect()

    var userBefore User
    userBefore.Token = token
	userAfter := userBefore

	db.First(&userBefore)
	userAfter.Name = name
	db.Model(&userBefore).Update(&userAfter)
    defer db.Close()
}

func GetUsers() []User{
	db := sqlConnect()
    var users []User
    db.Order("created_at asc").Find(&users)
    defer db.Close()

	return users
}

// func GetUserById(id int) User{
// 	db := sqlConnect()
//     var users []User
//     db.Where("id = ?", id).Find(&users)
//     defer db.Close()

// 	return users[0]
// }

func GetUserByToken(token string) User{
	db := sqlConnect()
    var users []User
    db.Where("token = ?", token).Find(&users)
    defer db.Close()

	return users[0]
}