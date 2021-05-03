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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method == "GET"{
		if token := r.FormValue("x-token"); token == "" {
			http.Error(w, fmt.Sprintf(`{"status": "missing required parameter 'x-token'"}`) , 503)
		}else{
			user := database.GetUserByToken(token)
			fmt.Fprint(w, `{"name": "` + user.Name + `"}`)
		}
	}else{
		http.Error(w, fmt.Sprintf(`{"status": "method not allow"}`) , 503)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method == "POST"{
		if name := r.FormValue("name"); name == "" {
			http.Error(w, fmt.Sprintf(`{"status": "missing required parameter 'name'"}`) , 503)
		}else{
			token := database.CreateUser(name)
			fmt.Fprint(w, `{"token": "` + token + `"}`)
		}
	}else{
		http.Error(w, fmt.Sprintf(`{"status": "method not allow"}`) , 503)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if r.Method == "PUT"{
		if token := r.FormValue("x-token"); token == "" {
			http.Error(w, fmt.Sprintf(`{"status": "missing required parameter 'x-token'"}`) , 503)
		}else if name := r.FormValue("name"); name == "" {
			http.Error(w, fmt.Sprintf(`{"status": "missing required parameter 'name'"}`) , 503)
		}else{
			database.UpdateUser(token, name)
			fmt.Fprint(w, `{"status": "success"}`)
		}
	}else{
		http.Error(w, fmt.Sprintf(`{"status": "method not allow"}`) , 503)
	}
}