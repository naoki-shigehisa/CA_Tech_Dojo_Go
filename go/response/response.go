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

// エラー出力
func httpError(w http.ResponseWriter, status string, code int) {
	content := `{"status": "` + status + `"}`
	w.WriteHeader(code)
	fmt.Fprint(w, content)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

// 全てのユーザーを取得
func GetUsers(w http.ResponseWriter, r *http.Request){
	// ユーザー情報取得
	users := database.GetUsers()

	// jsonエンコード
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {log.Fatal(err)}

	// 出力
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

// token指定でuserを取得
func GetUserByToken(w http.ResponseWriter, r *http.Request){
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドがGETか判定
	if r.Method == "GET"{
		// パラメータx-tokenが存在するかどうかを判定
		if token := r.FormValue("x-token"); token != "" {
			// userが見つかるがどうかを判定
			if user, err := database.GetUserByToken(token); err == nil{
				// 出力
				fmt.Fprint(w, `{"name": "` + user.Name + `"}`)
			}else{
				httpError(w, err.Error(), 500)
			}
		}else{
			httpError(w, "missing required parameter 'x-token'", 500)
		}
	}else{
		httpError(w, "method not allow", 500)
	}
}

// 新規ユーザーの作成
func CreateUser(w http.ResponseWriter, r *http.Request){
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドがPOSTかどうかを判定
	if r.Method == "POST"{
		// パラメータnameが存在するかどうかを判定
		if name := r.FormValue("name"); name != "" {
			// ユーザーを作成
			token := database.CreateUser(name)
			// tokenを出力
			fmt.Fprint(w, `{"token": "` + token + `"}`)
		}else{
			httpError(w, "missing required parameter 'name'", 500)
		}
	}else{
		httpError(w, "method not allow", 500)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドがPUTかどうかを判定
	if r.Method == "PUT"{
		// パラメータtokenが存在するかどうかを判定
		if token := r.FormValue("x-token"); token == "" {
			httpError(w, "missing required parameter 'x-token'", 500)
		// パラメータnameが存在するかどうかを判定
		}else if name := r.FormValue("name"); name == "" {
			httpError(w, "missing required parameter 'name'", 500)
		}else{
			// ユーザー情報更新
			if err := database.UpdateUser(token, name); err == nil{
				// 出力
				fmt.Fprint(w, `{"status": "success"}`)
			}else{
				httpError(w, err.Error(), 500)
			}
		}
	}else{
		httpError(w, "method not allow", 500)
	}
}