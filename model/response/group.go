package response

import (
	"2022summer/model/database"
)

type CreateGroupQ struct {
	GroupName string `json:"group_name" binding:"required"`
	GroupInfo string `json:"group_info"`
}
type CreateGroupA struct {
	CommonA
	Group database.Group `json:"group"`
}
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
	GroupID  uint64 `json:"group_id"`
	Username string `json:"username"`
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
	Status  int    `json:"status" binding:"oneof=1 2"`
}
type SetMemberStatusA struct {
	CommonA
}
type GetGroupsQ struct {
}
type GetGroupsA struct {
	CommonA
	Count  int              `json:"count"`
	Groups []database.Group `json:"groups"`
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
