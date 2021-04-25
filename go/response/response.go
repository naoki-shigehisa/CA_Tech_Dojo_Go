package response

import (
	"fmt"
	"bytes"
	"log"
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