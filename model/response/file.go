package response

import (
	"2022summer/model/database"
)

type GetFilesByNameQ struct {
	Name    string `json:"name"`
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetFilesByNameA struct {
	Message        string                `json:"message"`
	Success        bool                  `json:"success"`
	CountPPage     uint64                `json:"count_ppage"`
	PPage          []database.PPageID    `json:"ppage"`
	CountUmls      uint64                `json:"count_umls"`
	Umls           []database.UmlID      `json:"umls"`
	CountDocuments uint64                `json:"count_documents"`
	Documents      []database.DocumentID `json:"documents"`
}
type UploadImageA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Url     string `json:"url"`
}
type ConvertHtmlToPdfQ struct {
	Content string `json:"content"`
}
type ConvertHtmlToPdfA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Url     string `json:"url"`
}
