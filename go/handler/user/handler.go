package handler

import (
	"bytes"
	"dojo/model/user"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// エラー出力
func httpError(w http.ResponseWriter, status string, code int) {
	content := `{"status": "` + status + `"}`
	w.WriteHeader(code)
	fmt.Fprint(w, content)
}

// 全てのユーザーを取得
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// ユーザー情報取得
	users := model.GetUsers()

	// jsonエンコード
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(users); err != nil {
		log.Fatal(err)
	}

	// 出力
	fmt.Fprint(w, buf.String())
}

// token指定でuserを取得
func GetUserByToken(w http.ResponseWriter, r *http.Request) {
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドとパラメータを判定
	switch {
	case r.Method != "GET":
		httpError(w, "method not allow", 500)
	case r.FormValue("x-token") == "":
		httpError(w, "missing required parameter 'x-token'", 500)
	default:
		token := r.FormValue("x-token")
		// userが見つかるがどうかを判定
		if user, err := model.GetUserByToken(token); err == nil {
			// 出力
			fmt.Fprint(w, `{"name": "`+user.Name+`"}`)
		} else {
			httpError(w, err.Error(), 500)
		}
	}
}

// 新規ユーザーの作成
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドとパラメータを判定
	switch {
	case r.Method != "POST":
		httpError(w, "method not allow", 500)
	case r.FormValue("name") == "":
		httpError(w, "missing required parameter 'name'", 500)
	default:
		name := r.FormValue("name")
		// ユーザーを作成
		token := model.CreateUser(name)
		// tokenを出力
		fmt.Fprint(w, `{"token": "`+token+`"}`)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// headerの設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// メソッドとパラメータを判定
	switch {
	case r.Method != "PUT":
		httpError(w, "method not allow", 500)
	case r.FormValue("x-token") == "":
		httpError(w, "missing required parameter 'x-token'", 500)
	case r.FormValue("name") == "":
		httpError(w, "missing required parameter 'name'", 500)
	default:
		token := r.FormValue("x-token")
		name := r.FormValue("name")
		// ユーザー情報更新
		if err := model.UpdateUser(token, name); err == nil {
			// 出力
			fmt.Fprint(w, `{"status": "success"}`)
		} else {
			httpError(w, err.Error(), 500)
		}
	}
}
