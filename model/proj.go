package model

type Proj struct {
	ProjID   uint64 `gorm:"primary_key;not null;" json:"proj_id"`
	ProjName string `gorm:"size:255;not null;" json:"proj_name"`
	ProjInfo string `gorm:"type:text;" json:"proj_info"`
	Status   int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	GroupID  uint64 `gorm:"not null;" json:"group_id"`        // 项目所属团队
	UserID   uint64 `gorm:"not null;" json:"user_id"`         // 项目创建者
}

type Prototype struct {
	PrototypeID   uint64 `gorm:"primary_key;not null;" json:"prototype_id"`
	PrototypeName string `gorm:"size:255;not null;" json:"prototype_name"`
	PrototypeURL  string `gorm:"size:255;not null;" json:"prototype_url"`
	Status        int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	ProjID        uint64 `gorm:"not null;" json:"proj_id"`
}

type Uml struct {
	UmlID   uint64 `gorm:"primary_key;not null;" json:"uml_id"`
	UmlName string `gorm:"size:255;not null;" json:"uml_name"`
	UmlURL  string `gorm:"size:255;not null;" json:"uml_url"`
	Status  int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	ProjID  uint64 `gorm:"not null;" json:"proj_id"`
}

type Document struct {
	DocumentID   uint64 `gorm:"primary_key;not null;" json:"document_id"`
	DocumentName string `gorm:"size:255;not null;" json:"document_name"`
	DocumentURL  string `gorm:"size:255;not null;" json:"document_url"`
	Status       int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	ProjID       uint64 `gorm:"not null;" json:"proj_id"`
}
