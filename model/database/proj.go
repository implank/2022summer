package database

import "time"

type Proj struct {
	ProjID    uint64    `gorm:"primary_key;not null;" json:"proj_id"`
	ProjName  string    `gorm:"size:255;not null;" json:"proj_name"`
	ProjInfo  string    `gorm:"type:text;" json:"proj_info"`
	Status    int       `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	GroupID   uint64    `gorm:"not null;" json:"group_id"`        // 项目所属团队
	UserID    uint64    `gorm:"not null;" json:"user_id"`         // 项目创建者
	Top       int       `gorm:"not null;default:1;" json:"top"`   // 1 不置顶、2 置顶
	CreatedAt time.Time `gorm:"not null;default:Now();" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:Now();" json:"updated_at"`
	EditTimes int       `gorm:"not null;default:0;" json:"edit_times"`
}
