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
	http.HandleFunc("/users", response.GetUsers)
	http.HandleFunc("/user/get", response.GetUserByToken)
	http.HandleFunc("/user/create", response.CreateUser)
	http.HandleFunc("/user/update", response.UpdateUser)
	http.ListenAndServe(":1323", nil)
}