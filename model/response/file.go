package response

import (
	"2022summer/model/database"
)

type GetProjByIDQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjByIDA struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Proj    database.Proj `json:"proj"`
}

type GetProjPrototypesQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjPrototypesA struct {
	Message    string               `json:"message"`
	Success    bool                 `json:"success"`
	Count      uint64               `json:"count"`
	Prototypes []database.Prototype `json:"prototypes"`
}

type GetProjUmlsQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjUmlsA struct {
	Message string         `json:"message"`
	Success bool           `json:"success"`
	Count   uint64         `json:"count"`
	Umls    []database.Uml `json:"umls"`
}

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

type CreatePrototypeQ struct {
	PrototypeName string `json:"prototype_name" binding:"required"`
	ProjID        uint64 `json:"proj_id"`
}

type CreatePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type CreateUmlQ struct {
	UmlName string `json:"uml_name" binding:"required"`
	ProjID  uint64 `json:"proj_id"`
}

type CreateUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type CreateDocumentQ struct {
	DocumentName string `json:"document_name" binding:"required"`
	ProjID       uint64 `json:"proj_id"`
}

type CreateDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

//type UploadDocumentQ struct {
//	DocumentName string `json:"document_name"`
//	DocumentID   uint64 `json:"document_id"`
//	ProjID       uint64 `json:"proj_id"`
//}

//type UploadDocumentA struct {
//	CommonA
//	Document database.Document
//	Rank     uint64 `json:"count"`
//}

type UploadDocumentA struct {
	CommonA
}

type UpdatePrototypeQ struct {
	PrototypeID   uint64 `json:"prototype_id" binding:"required"`
	PrototypeName string `json:"prototype_name" binding:"required"`
}

type UpdatePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateUmlQ struct {
	UmlID   uint64 `json:"uml_id" binding:"required"`
	UmlName string `json:"uml_name" binding:"required"`
}

type UpdateUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateDocumentQ struct {
	DocumentID   uint64 `json:"document_id" binding:"required"`
	DocumentName string `json:"document_name" binding:"required"`
}

type UpdateDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePrototypeToBinQ struct {
	PrototypeID uint64 `json:"prototype_id" binding:"required"`
}

type MovePrototypeToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveUmlToBinQ struct {
	UmlID uint64 `json:"uml_id" binding:"required"`
}

type MoveUmlToBinA struct {
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

type GetFilesByNameQ struct {
	Name string `json:"name"`
}

type GetFilesByNameA struct {
	Message         string               `json:"message"`
	Success         bool                 `json:"success"`
	CountPrototypes uint64               `json:"count_prototypes"`
	Prototypes      []database.Prototype `json:"prototypes"`
	CountUmls       uint64               `json:"count_umls"`
	Umls            []database.Uml       `json:"umls"`
	CountDocuments  uint64               `json:"count_documents"`
	Documents       []database.Document  `json:"documents"`
}
