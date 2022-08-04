package response

import (
	"2022summer/model/database"
)

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
	User   database.User `json:"user"`
	Poster database.User `json:"poster"`
}
type ModifyPasswordQ struct {
	Password string `json:"password" binding:"required"`
}
type ModifyPasswordA struct {
	CommonA
}
type ModifyInfoQ struct {
	Username string `json:"username"`
	Age      uint   `json:"age"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
}
type ModifyInfoA struct {
	CommonA
}
