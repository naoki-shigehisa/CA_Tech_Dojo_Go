package main

import (
  "net/http"
  "dojo/response"
)

func main() {
	// db := database.SqlConnect()
  	// db.AutoMigrate(&database.User{})
  	// defer db.Close()
	
	// database.Create()
	// users := database.Get()

	// fmt.Println(users[0].Model.ID)
	// fmt.Println(users[0].Model.CreatedAt)
	// fmt.Println(users[0].Model.UpdatedAt)
	// fmt.Println(users[0].Model.DeletedAt)
	// fmt.Println(users[0].Token)
	// fmt.Println(users[0].Name)

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