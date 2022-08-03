package response

import (
	"2022summer/model"
)

type CreateProjQ struct {
	ProjName string `json:"proj_name" binding:"required"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
	GroupID  uint64 `json:"group_id" binding:"required"`
}

type CreateProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateProjQ struct {
	ProjID   uint64 `json:"proj_id" binding:"required"`
	ProjName string `json:"proj_name" binding:"required"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
}

type UpdateProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveProjBinQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type MoveProjBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type GetProjAllQ struct {
	// GroupID uint64 `json:"group_id"`
}

type GetProjAllA struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Count   uint64       `json:"count"`
	Projs   []model.Proj `json:"projs"`
}

type GetProjCreateQ struct {
	// GroupID uint64 `json:"group_id"`
}

type GetProjCreateA struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Count   uint64       `json:"count"`
	Projs   []model.Proj `json:"projs"`
}

type GetProjJoinQ struct {
	// GroupID uint64 `json:"group_id"`
}

type GetProjJoinA struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Count   uint64       `json:"count"`
	Projs   []model.Proj `json:"projs"`
}

type GetProjByNameQ struct {
	ProjName string `json:"proj_name" binding:"omitempty"`
}

type GetProjByNameA struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Count   uint64       `json:"count"`
	Projs   []model.Proj `json:"projs"`
}

/* * * * * * * * * * * */

type GetProjByIDQ struct {
	ProjID uint64 `json:"proj_id"`
}

type GetProjByIDA struct {
	Message string     `json:"message"`
	Success bool       `json:"success"`
	Proj    model.Proj `json:"proj"`
}

type GetProjPrototypesQ struct {
	ProjID uint64 `json:"proj_id"`
}

type GetProjPrototypesA struct {
	Message    string            `json:"message"`
	Success    bool              `json:"success"`
	Count      uint64            `json:"count"`
	Prototypes []model.Prototype `json:"prototypes"`
}

type GetProjUmlsQ struct {
	ProjID uint64 `json:"proj_id"`
}

type GetProjUmlsA struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Count   uint64      `json:"count"`
	Umls    []model.Uml `json:"umls"`
}

type GetProjDocumentsQ struct {
	ProjID uint64 `json:"proj_id"`
}

type GetProjDocumentsA struct {
	Message   string           `json:"message"`
	Success   bool             `json:"success"`
	Count     uint64           `json:"count"`
	Documents []model.Document `json:"documents"`
}

type CreatePrototypeQ struct {
	PrototypeName string `json:"prototype_name"`
	ProjID        uint64 `json:"proj_id"`
}

type CreatePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type CreateUmlQ struct {
	UmlName string `json:"uml_name"`
	ProjID  uint64 `json:"proj_id"`
}

type CreateUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type CreateDocumentQ struct {
	DocumentName string `json:"document_name"`
	ProjID       uint64 `json:"proj_id"`
}

type CreateDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdatePrototypeQ struct {
	PrototypeID   uint64 `json:"prototype_id"`
	PrototypeName string `json:"prototype_name"`
}

type UpdatePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateUmlQ struct {
	UmlID   uint64 `json:"uml_id"`
	UmlName string `json:"uml_name"`
}

type UpdateUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateDocumentQ struct {
	DocumentID   uint64 `json:"document_id"`
	DocumentName string `json:"document_name"`
}

type UpdateDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePrototypeToBinQ struct {
	PrototypeID uint64 `json:"prototype_id"`
}

type MovePrototypeToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveUmlToBinQ struct {
	UmlID uint64 `json:"uml_id"`
}

type MoveUmlToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveDocumentToBinQ struct {
	DocumentID uint64 `json:"document_id"`
}

type MoveDocumentToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type GetFilesByNameQ struct {
	Name string `json:"name"`
}

type GetFilesByNameA struct {
	Message         string            `json:"message"`
	Success         bool              `json:"success"`
	CountPrototypes uint64            `json:"count_prototypes"`
	Prototypes      []model.Prototype `json:"prototypes"`
	CountUmls       uint64            `json:"count_umls"`
	Umls            []model.Uml       `json:"umls"`
	CountDocuments  uint64            `json:"count_documents"`
	Documents       []model.Document  `json:"documents"`
}

/* * * * * * * * * * * */

type DeleteProjQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type DeleteProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePrototypeFromBinQ struct {
	PrototypeID uint64 `json:"prototype_id"`
}

type MovePrototypeFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveUmlFromBinQ struct {
	UmlID uint64 `json:"uml_id"`
}

type MoveUmlFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveDocumentFromBinQ struct {
	DocumentID int64 `json:"document_id"`
}

type MoveDocumentFromBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeletePrototypeQ struct {
	PrototypeID uint64 `json:"prototype_id"`
}

type DeletePrototypeA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteUmlQ struct {
	UmlID uint64 `json:"uml_id"`
}

type DeleteUmlA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteDocumentQ struct {
	DocumentID uint64 `json:"document_id"`
}

type DeleteDocumentA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
