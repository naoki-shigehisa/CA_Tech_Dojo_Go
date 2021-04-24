package main

import (
  "fmt"
  "dojo/database"
)

func main() {
	db := database.SqlConnect()
  	db.AutoMigrate(&database.User{})
  	defer db.Close()
	
	//create()
	users := database.Get()

	fmt.Println(users[0].Model.ID)
	fmt.Println(users[0].Model.CreatedAt)
	fmt.Println(users[0].Model.UpdatedAt)
	fmt.Println(users[0].Model.DeletedAt)
	fmt.Println(users[0].Token)
	fmt.Println(users[0].Name)
}