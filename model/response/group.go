package response

import (
	"2022summer/model/database"
)

type GetIdentityQ struct {
	GroupID uint64 `json:"group_id"`
}
type GetIdentityA struct {
	CommonA
	Status int `json:"status"`
}
type GetMembersQ struct {
	GroupID uint64 `json:"group_id"`
}
type GetMembersA struct {
	CommonA
	Members []database.GroupMember `json:"members"`
}
type AddMemberQ struct {
	GroupID uint64 `json:"group_id"`
	UserID  uint64 `json:"user_id"`
}
type AddMemberA struct {
	CommonA
}
type RemoveMemberQ struct {
	UserID  uint64 `json:"user_id"`
	GroupID uint64 `json:"group_id"`
}
type RemoveMemberA struct {
	CommonA
}
type SetMemberStatusQ struct {
	UserID  uint64 `json:"user_id"`
	GroupID uint64 `json:"group_id"`
	Status  int    `json:"status"`
}
type SetMemberStatusA struct {
	CommonA
}

var (
	NOAUTH = CommonA{
		Message: "没有权限",
		Success: false,
		Code:    0,
	}
	NOMENBER = CommonA{
		Message: "用户不在该团队中",
		Success: false,
		Code:    1,
	}
	PARAMERROR = CommonA{
		Message: "参数错误",
		Success: false,
		Code:    2,
	}
	DBERROR = CommonA{
		Message: "数据库错误",
		Success: false,
		Code:    3,
	}
)
