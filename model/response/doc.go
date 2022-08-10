package response

import "2022summer/model/database"

type GetProjDocumentsQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjDocumentsA struct {
	Message   string                `json:"message"`
	Success   bool                  `json:"success"`
	Count     uint64                `json:"count"`
	Documents []database.DocumentID `json:"documents"`
}

type EnterDocumentQ struct {
	DocumentID uint64 `json:"document_id" binding:"required"`
}

type EnterDocumentA struct {
	CommonA
	Document database.Document `json:"document"`
	Rank     uint64            `json:"rank"`
}

type QuitDocumentQ struct {
	DocumentID uint64 `json:"document_id" binding:"required"`
}

type QuitDocumentA struct {
	CommonA
	Document database.Document `json:"document"`
	Rank     uint64            `json:"rank"`
}

type CreateDocumentQ struct {
	DocumentName string `json:"document_name" binding:"required"`
	ProjID       uint64 `json:"proj_id"`
}

type CreateDocumentA struct {
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	DocumentID uint64 `json:"document_id"`
}

type UploadDocumentQ struct {
	DocumentID uint64 `json:"document_id"`
	Content    string `json:"content"`
}

type UploadDocumentA struct {
	CommonA
	Document database.Document
}

type UpdateDocumentQ struct {
	DocumentID   uint64 `json:"document_id" binding:"required"`
	DocumentName string `json:"document_name" binding:"required"`
}

type UpdateDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveDocumentToBinQ struct {
	DocumentID uint64 `json:"document_id" binding:"required"`
}

type MoveDocumentToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

//---
type CreateDocFileQ struct {
	DirID    uint64 `json:"dir_id" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
	IsDir    int    `json:"is_dir"`
}
type CreateDocFileA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
}
type GetDocFilesQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
}
type GetDocFilesA struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Code    int           `json:"code"`
	File    database.File `json:"file"`
}
