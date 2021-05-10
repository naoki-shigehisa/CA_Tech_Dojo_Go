package handler

// ユーザー作成用
type UserCreateRequest struct {
	Name       string `json:"name"`
}

// ユーザー取得用
type UserGetRequest struct {
	Token       string `json:"x-token"`
}

// ユーザー作成用
type UserUpdateRequest struct {
	Token       string `json:"x-token"`
	Name       string `json:"name"`
}