package response

import (
	"fmt"
	"bytes"
	"log"
	// "reflect"
	// "strconv"
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

// func GetUserById(w http.ResponseWriter, r *http.Request){
// 	if r.Method == "GET"{
// 		id, _ := strconv.Atoi(r.FormValue("id"))
// 		users := database.GetUser(id)

// 		var buf bytes.Buffer
// 		enc := json.NewEncoder(&buf)
// 		if err := enc.Encode(users); err != nil {log.Fatal(err)}
// 		fmt.Fprint(w, buf.String())
// 	}else{
// 		fmt.Fprint(w, `{"status": "method not allow"}`)
// 	}
// }

func GetUserByToken(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		if token := r.FormValue("x-token"); token == "" {
			fmt.Fprint(w, `{"status": "missing required parameter 'x-token'"}`)
		}else{
			user := database.GetUserByToken(token)
			fmt.Fprint(w, `{"name": "` + user.Name + `"}`)
		}
	}else{
		fmt.Fprint(w, `{"status": "method not allow"}`)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		if name := r.FormValue("name"); name == "" {
			fmt.Fprint(w, `{"status": "missing required parameter 'name'"}`)
		}else{
			token := database.CreateUser(name)
			fmt.Fprint(w, `{"token": "` + token + `"}`)
		}
	}else{
		fmt.Fprint(w, `{"status": "method not allow"}`)
	}
}