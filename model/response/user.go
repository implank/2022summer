package response

import "2022summer/model"

type RegisterQ struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type LoginQ struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

type GetUserInfoQ struct {
	UserID uint64 `json:"user_id"`
}

type GetUserInfoA struct {
	Message string     `json:"message"`
	Success bool       `json:"success"`
	User    model.User `json:"user"`
	Poster  model.User `json:"poster"`
}
