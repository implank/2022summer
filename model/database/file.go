package database

type PPage struct { // 设计原型的页面
	PPageID   uint64 `gorm:"primary_key;not null;" json:"ppage_id"`
	PPageName string `gorm:"size:255;not null;" json:"ppage_name"`
	PPageData string `gorm:"size:max;" json:"ppage_data"`
	PPageURL  string `gorm:"size:255;" json:"ppage_url"`       // 先写上
	Status    int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	ProjID    uint64 `gorm:"not null;" json:"proj_id"`
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
	Count        uint64 `gorm:"default:0;not null" json:"count"`
}

/* * * * * * * * * * * */

type PPageID struct {
	PPageID   uint64 `json:"ppage_id"`
	PPageName string `json:"ppage_name"`
}
