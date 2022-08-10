package response

import (
	"2022summer/model/database"
)

type RegisterQ struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
	Token string        `json:"token"`
	User  database.User `json:"user"`
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
type GetMessagesQ struct {
}
type GetMessagesA struct {
	CommonA
	Count    int                `json:"count"`
	Messages []database.Message `json:"messages"`
}
type DeclineInvitationQ struct {
	MessageID uint64 `json:"message_id"`
}
type DeclineInvitationA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}
type AcceptInvitationQ struct {
	MessageID uint64 `json:"message_id"`
}
type AcceptInvitationA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}
type ReadAllMessagesQ struct {
}
type ReadAllMessagesA struct {
	CommonA
}
type UploadAvatarA struct {
	CommonA
	AvatarUrl string `json:"avatar_url"`
}
