package response

import "2022summer/model/database"

type GetProjDocumentsQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjDocumentsA struct {
	Message   string              `json:"message"`
	Success   bool                `json:"success"`
	Count     uint64              `json:"count"`
	Documents []database.Document `json:"documents"`
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
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UploadDocumentQ struct {
	DocumentName string `json:"document_name"`
	DocumentID   uint64 `json:"document_id"`
	ProjID       uint64 `json:"proj_id"`
	Content      string `json:"content"`
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
