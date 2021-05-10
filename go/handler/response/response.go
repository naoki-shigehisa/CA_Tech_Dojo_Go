package handler

// ユーザー作成用
type UserCreateResponse struct {
	Token       string `json:"token"`
}

// ユーザー取得用
type UserGetResponse struct {
	Name       string `json:"name"`
}

// その他エラー等
type Response struct {
	Status       string `json:"status"`
}