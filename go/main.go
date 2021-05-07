package main

import (
	"dojo/handler/user"
	"dojo/model/general"
	"net/http"
)

func main() {
	// sqlサーバーに接続
	model.SqlConnect()

	http.HandleFunc("/", handler.Handler)

	// 全てのuserを取得
	http.HandleFunc("/users", handler.GetUsers)
	// token指定でuserを取得
	http.HandleFunc("/user/get", handler.GetUserByToken)
	// 新規userを作成
	http.HandleFunc("/user/create", handler.CreateUser)
	// user情報を更新
	http.HandleFunc("/user/update", handler.UpdateUser)

	// サーバー起動
	http.ListenAndServe(":1323", nil)

	defer model.MysqlDb.Close()
}
