package main

import (
  "fmt"
  "net/http"
//   "dojo/database"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1323", nil)

	// db := database.SqlConnect()
  	// db.AutoMigrate(&database.User{})
  	// defer db.Close()
	
	// create()
	// users := database.Get()

	// fmt.Println(users[0].Model.ID)
	// fmt.Println(users[0].Model.CreatedAt)
	// fmt.Println(users[0].Model.UpdatedAt)
	// fmt.Println(users[0].Model.DeletedAt)
	// fmt.Println(users[0].Token)
	// fmt.Println(users[0].Name)
}