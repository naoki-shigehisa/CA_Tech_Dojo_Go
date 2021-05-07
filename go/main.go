package main

import (
	"dojo/response"
	"net/http"
)

func main() {
	http.HandleFunc("/", response.Handler)

	// 全てのuserを取得
	http.HandleFunc("/users", response.GetUsers)
	// token指定でuserを取得
	http.HandleFunc("/user/get", response.GetUserByToken)
	// 新規userを作成
	http.HandleFunc("/user/create", response.CreateUser)
	// user情報を更新
	http.HandleFunc("/user/update", response.UpdateUser)

	// サーバー起動
	http.ListenAndServe(":1323", nil)
}
