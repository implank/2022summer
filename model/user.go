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
