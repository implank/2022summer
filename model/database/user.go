package database

type User struct {
	UserID    uint64 `gorm:"primary_key;not null;" json:"user_id"`
	Username  string `gorm:"size:255;not null;" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
	RealName  string `gorm:"size:255;not null;" json:"real_name"`
	Email     string `gorm:"size:255;not null;" json:"email"`
	Sex       string `gorm:"size:255;default:'未知';" json:"sex"`
	Age       uint   `gorm:"default:0;" json:"age"`
	UserInfo  string `gorm:"type:text;" json:"user_info"`
	AvatarUrl string `gorm:"default:'default.jpg'" json:"avatar_url"`
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
	Status     int    `gorm:"not null;" json:"status"` // 0 未读 1 已读
}
