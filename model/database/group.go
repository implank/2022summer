package database

type Group struct {
	GroupID   uint64 `gorm:"primary_key;not null;" json:"group_id"`
	GroupName string `gorm:"size:255;not null;" json:"group_name"`
	GroupInfo string `gorm:"type:text;" json:"group_info"`
	UserID    uint64 `gorm:"not null;" json:"user_id"` // 团队创建者
}

type Identity struct {
	UserID  uint64 `gorm:"primary_key;auto_increment:false;" json:"user_id"`
	GroupID uint64 `gorm:"primary_key;auto_increment:false;" json:"group_id"`
	Status  int    `gorm:"not null;" json:"status"` // 1 普通成员、2 管理员、3 团队创建者
}

// Message
// * 1 待Rec确认的邀请
// * 2 已被Rec拒绝的邀请
// * 3 已被Rec同意的邀请
// * 4 用户拒绝了邀请（仅发送给管理员
// * 5 广播用户加入了团队（发送给全体团队成员
type Message struct {
	MessageID  uint64 `gorm:"primary_key;" json:"message_id"`
	ReceiverID uint64 `gorm:"not null;" json:"receiver_id"`
	SenderID   uint64 `gorm:"not null;" json:"sender_id"`
	Content    string `gorm:"type:text;" json:"content"`
	GroupID    uint64 `gorm:"not null;" json:"group_id"`
	Type       int    `gorm:"not null;" json:"type"`
}

/* * * * * * * * * * * */

type GroupMember struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Email    string `json:"email"`
	Status   string `json:"status"` // 1 普通成员、2 管理员、3 团队创建者
}
