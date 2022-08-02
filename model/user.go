package model

type User struct {
	UserID   uint64 `gorm:"primary_key;not null;"`
	Username string `gorm:"size:255;not null;"`
	Password string `gorm:"size:255;not null;"`
	RealName string `gorm:"size:255;not null;"`
	Email    string `gorm:"size:255;not null;"`
	Sex      string `gorm:"size:255;default:'未知';"`
	Age      uint   `gorm:"default:0;"`
	UserInfo string `gorm:"type:text;"`
}

type Group struct {
	GroupID   uint64 `gorm:"primary_key;not null;"`
	GroupName string `gorm:"size:255;not null;"`
	GroupInfo string `gorm:"type:text;"`
	UserID    uint64 `gorm:"not null;"` // 团队创建者
}

type Identity struct {
	UserID  uint64 `gorm:"primary_key;auto_increment:false;"`
	GroupID uint64 `gorm:"primary_key;auto_increment:false;"`
	Status  int    `gorm:"not null;"` // 1 普通成员、2 管理员、3 团队创建者
}

/* * * * * * * * * * * */

type GroupMembers struct {
	UserID   uint64
	Username string
	RealName string
	Email    string
	Status   string
}
