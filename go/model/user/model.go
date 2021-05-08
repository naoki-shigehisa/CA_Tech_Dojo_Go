package model

import (
	"github.com/jinzhu/gorm"
)

// ユーザー情報用構造体
type User struct {
	gorm.Model `json:"info"`
	Token      string `json:"token"`
	Name       string `json:"name"`
}
