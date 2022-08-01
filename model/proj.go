package model

type Proj struct {
	ProjID    uint64 `gorm:"primary_key;not null;"`
	ProjName  string `gorm:"size:255;not null;"`
	ProjInfo  string `gorm:"type:text;"`
	GroupID   uint64 `gorm:"not null;"`
	Status    int    `gorm:"default:1;not null"` // 1 正常、2 回收站
	Prototype string `gorm:""`                   // url
	Uml       string `gorm:""`                   // url
	Document  string `gorm:""`                   // url
}
