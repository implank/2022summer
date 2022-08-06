package response

import "2022summer/model/database"

type DeleteProjQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type DeleteProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePrototypeFromBinQ struct {
	PrototypeID uint64 `json:"prototype_id" binding:"required"`
}

type MovePrototypeFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveUmlFromBinQ struct {
	UmlID uint64 `json:"uml_id" binding:"required"`
}

type MoveUmlFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveDocumentFromBinQ struct {
	DocumentID int64 `json:"document_id" binding:"required"`
}

type MoveDocumentFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeletePrototypeQ struct {
	PrototypeID uint64 `json:"prototype_id" binding:"required"`
}

type DeletePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteUmlQ struct {
	UmlID uint64 `json:"uml_id" binding:"required"`
}

type DeleteUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteDocumentQ struct {
	DocumentID uint64 `json:"document_id" binding:"required"`
}

type DeleteDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type GetProjInBinQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetProjInBinA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetFilesInBinQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetFilesInBinA struct {
	Message         string               `json:"message"`
	Success         bool                 `json:"success"`
	CountPrototypes uint64               `json:"count_prototypes"`
	Prototypes      []database.Prototype `json:"prototypes"`
	CountUmls       uint64               `json:"count_umls"`
	Umls            []database.Uml       `json:"umls"`
	CountDocuments  uint64               `json:"count_documents"`
	Documents       []database.Document  `json:"documents"`
}
