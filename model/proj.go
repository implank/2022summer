package model

type Proj struct {
	ProjID   uint64 `gorm:"primary_key;not null;"`
	ProjName string `gorm:"size:255;not null;"`
	ProjInfo string `gorm:"type:text;"`
	GroupID  uint64 `gorm:"not null;"`
	Status   int    `gorm:"default:1;not null"` // 1 正常、2 回收站
}

type Prototype struct {
	PrototypeID   uint64 `gorm:"primary_key;not null;"`
	PrototypeName string `gorm:"size:255;not null;"`
	PrototypeURL  string `gorm:"size:255;not null;"`
	ProjID        uint64 `gorm:"not null;"`
}

type Uml struct {
	UmlID   uint64 `gorm:"primary_key;not null;"`
	UmlName string `gorm:"size:255;not null;"`
	UmlURL  string `gorm:"size:255;not null;"`
	ProjID  uint64 `gorm:"not null;"`
}

type Document struct {
	DocumentID   uint64 `gorm:"primary_key;not null;"`
	DocumentName string `gorm:"size:255;not null;"`
	DocumentURL  string `gorm:"size:255;not null;"`
	ProjID       uint64 `gorm:"not null;"`
}
