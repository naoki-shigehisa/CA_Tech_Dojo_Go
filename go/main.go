package main

import (
  "fmt"
  "bytes"
  "log"
  "encoding/json"
  "net/http"
  "dojo/database"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func getUsers(w http.ResponseWriter, r *http.Request){
	users := database.Get()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {log.Fatal(err)}
	fmt.Fprint(w, buf.String())
}

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

	http.HandleFunc("/", handler)
	http.HandleFunc("/users", getUsers)
	http.ListenAndServe(":1323", nil)
}