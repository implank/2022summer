package database

type PPage struct { // 设计原型的页面
	PPageID   uint64 `gorm:"primary_key;not null;" json:"ppage_id"`
	PPageName string `gorm:"size:255;not null;" json:"ppage_name"`
	PPageData string `gorm:"size:max;" json:"ppage_data"`
	PPageURL  string `gorm:"size:255;" json:"ppage_url"`       // 先写上
	Status    int    `gorm:"default:1;not null" json:"status"` // 1 正常、2 回收站
	ProjID    uint64 `gorm:"not null;" json:"proj_id"`
}

type SharedPPage struct { // 分享的页面
	PPageID   uint64 `gorm:"not null;" json:"ppage_id"`
	ExpiresAt int64  `gorm:"not null;" json:"expires_at"`
	Token     string `gorm:"size:1024;not null;" json:"token"`
}

type Uml struct {
	UmlID   uint64 `gorm:"primary_key;not null;" json:"uml_id"`
	UmlName string `gorm:"size:255;not null;" json:"uml_name"`
	UmlData string `gorm:"size:max;" json:"uml_data"`
	UmlURL  string `gorm:"size:255;not null;" json:"uml_url"` // 先写上
	Status  int    `gorm:"default:1;not null" json:"status"`  // 1 正常、2 回收站
	ProjID  uint64 `gorm:"not null;" json:"proj_id"`
}

type Document struct {
	DocumentID   uint64 `gorm:"primary_key;not null;" json:"document_id"`
	DocumentName string `gorm:"size:255;not null;" json:"document_name"`
	DocumentURL  string `gorm:"size:255;not null;" json:"document_url"` // unuse
	Status       int    `gorm:"default:1;not null" json:"status"`       // 1 正常、2 回收站
	ProjID       uint64 `gorm:"not null;" json:"proj_id"`               // 非0表示项目文档并表示所属项目
	Count        uint64 `gorm:"default:0;not null" json:"count"`        // 仅为文档文件使用
	Content      string `gorm:"size:max;" json:"content"`               // 文档内容
	DirID        uint64 `gorm:"not null;" json:"dir_id"`                // 文档所属目录
	IsDir        int    `gorm:"not null;" json:"is_dir"`                // 是否为目录 0为文件 1为目录
	IsFixed      int    `gorm:"not null;" json:"is_fixed"`              // 是否为固定文件 0为非固定 1为固定
}

/* * * * * * * * * * * */

type PPageID struct {
	PPageID   uint64 `json:"ppage_id"`
	PPageName string `json:"ppage_name"`
}

type UmlID struct {
	UmlID   uint64 `json:"uml_id"`
	UmlName string `json:"uml_name"`
}

type DocumentID struct {
	DocumentID   uint64 `json:"document_id"`
	DocumentName string `json:"document_name"`
}

type File struct {
	FileID         uint64 `json:"file_id"`
	FileName       string `json:"file_name"`
	IsDir          int    `json:"is_dir"`
	ContainedFiles []File `json:"contained_files"`
}
