package response

type RemoveMemberQ struct {
	UserID  uint64 `json:"user_id"`
	GroupID uint64 `json:"group_id"`
}
type RemoveMemberA struct {
	CommonA
}
