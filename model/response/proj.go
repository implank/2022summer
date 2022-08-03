package response

import (
	"2022summer/model"
)

type CreateProjQ struct {
	ProjName string `json:"proj_name"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
	GroupID  uint64 `json:"group_id"`
}

type CreateProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateProjQ struct {
	ProjID   uint64 `json:"proj_id"`
	ProjName string `json:"proj_name"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
}

type UpdateProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MoveProjBinQ struct {
	ProjID uint64 `json:"proj_id"`
}

type MoveProjBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type DeleteProjQ struct {
	ProjID uint64 `json:"proj_id"`
}

type DeleteProjA struct {
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
	ProjName string `json:"proj_name"`
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
}

type CreatePrototypeA struct {
}

type CreateUmlQ struct {
}

type CreateUmlA struct {
}

type CreateDocumentQ struct {
}

type CreateDocumentA struct {
}

type UpdatePrototypeQ struct {
}

type UpdatePrototypeA struct {
}

type UpdateUmlQ struct {
}

type UpdateUmlA struct {
}

type UpdateDocumentQ struct {
}

type UpdateDocumentA struct {
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

type GetSthByNameQ struct {
	Name string `json:"name"`
}

type GetSthByNameA struct {
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

type MovePrototypeFromBinQ struct {
}

type MovePrototypeFromBinA struct {
}

type MoveUmlFromBinQ struct {
}

type MoveUmlFromBinA struct {
}

type MoveDocumentFromBinQ struct {
}

type MoveDocumentFromBinA struct {
}
