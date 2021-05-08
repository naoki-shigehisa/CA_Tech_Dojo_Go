package model

import (
	"errors"
	"dojo/controller/user"
	"dojo/model/general"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 新規ユーザーを作成
func CreateUser(name string) string {
	db := model.MysqlDb

	// 重複しないようにtokenを生成
	token, _ := controller.MakeRandomStr(10)
	var users []User
	for db.Where("token = ?", token).Find(&users); len(users) != 0; db.Where("token = ?", token).Find(&users) {
		token, _ = controller.MakeRandomStr(10)
	}

	fmt.Println("create user " + name + " with token " + token)
	db.Create(&User{Token: token, Name: name})

	return token
}

// ユーザー情報を更新
func UpdateUser(token string, name string) error {
	db := model.MysqlDb

	var userBefore User
	userAfter := userBefore
	db.First(&userBefore, "token=?", token)

	if userBefore.Model.ID != 0 {
		userAfter.Name = name
		db.Model(&userBefore).Update(&userAfter)
		return nil
	} else {
		return errors.New("user not found")
	}
}

// 全てのユーザーを取得
func GetUsers() []User {
	db := model.MysqlDb
	var users []User
	db.Order("created_at asc").Find(&users)

	return users
}

// token指定でユーザーを取得
func GetUserByToken(token string) (User, error) {
	db := model.MysqlDb
	var users []User
	if db.Where("token = ?", token).Find(&users); len(users) != 0 {
		return users[0], nil
	} else {
		var user User
		return user, errors.New("user not found")
	}
}
