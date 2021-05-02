package response

import (
	"fmt"
	"bytes"
	"log"
	// "reflect"
	"strconv"
	"encoding/json"
	"net/http"
	"dojo/database"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func GetUsers(w http.ResponseWriter, r *http.Request){
	users := database.GetUsers()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {log.Fatal(err)}
	fmt.Fprint(w, buf.String())
}

func GetUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		id, _ := strconv.Atoi(r.FormValue("id"))
		users := database.GetUser(id)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(users); err != nil {log.Fatal(err)}
		fmt.Fprint(w, buf.String())
	}else{
		fmt.Fprint(w, `{"status": "method not allow"}`)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		name := r.FormValue("name")
		token := database.CreateUser(name)
		fmt.Fprint(w, `{"token": "` + token + `"}`)
	}else{
		fmt.Fprint(w, `{"status": "method not allow"}`)
	}
}