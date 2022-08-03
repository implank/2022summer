package response

import "2022summer/model"

type RegisterQ struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterA struct {
	CommonA
}

type LoginQ struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginA struct {
	CommonA
	Token string `json:"token"`
}

type GetUserInfoQ struct {
	UserID uint64 `json:"user_id"`
}

type GetUserInfoA struct {
	CommonA
	User   model.User `json:"user"`
	Poster model.User `json:"poster"`
}
