package database

type User struct {
	UserID   uint64 `gorm:"primary_key;not null;" json:"user_id"`
	Username string `gorm:"size:255;not null;" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	RealName string `gorm:"size:255;not null;" json:"real_name"`
	Email    string `gorm:"size:255;not null;" json:"email"`
	Sex      string `gorm:"size:255;default:'未知';" json:"sex"`
	Age      uint   `gorm:"default:0;" json:"age"`
	UserInfo string `gorm:"type:text;" json:"user_info"`
}
