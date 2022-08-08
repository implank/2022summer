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

/* * * * * * * * * * * */

type GroupMember struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Email    string `json:"email"`
	Status   string `json:"status"` // 1 普通成员、2 管理员、3 团队创建者
}
