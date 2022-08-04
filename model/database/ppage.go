package database

type PPage struct { // Prototype 简写为 P, 一个 Prototype 有多个 PPage
	PPageID     uint64 `gorm:"primary_key;not null;" json:"ppage_id"`
	PPageName   string `gorm:"size:255;not null;" json:"ppage_name"`
	PPageData   string `gorm:"size:max;" json:"ppage_data"`
	PPageURL    string `gorm:"size:255;" json:"ppage_url"` // 先写上
	PrototypeID uint64 `gorm:"not null;" json:"prototype_id"`
}

/* * * * * * * * * * * */

type PPageID struct {
	PPageID   uint64 `json:"ppage_id"`
	PPageName string `json:"ppage_name"`
}
