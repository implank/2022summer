package response

import "2022summer/model/database"

type GetProjUmlsQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjUmlsA struct {
	Message string           `json:"message"`
	Success bool             `json:"success"`
	Count   uint64           `json:"count"`
	Umls    []database.UmlID `json:"umls"`
}

type GetUmlByIDQ struct {
	UmlID uint64 `json:"uml_id" binding:"required"`
}

type GetUmlByIDA struct {
	Message string       `json:"message"`
	Success bool         `json:"success"`
	Count   uint64       `json:"count"`
	Uml     database.Uml `json:"uml"`
}

type CreateUmlQ struct {
	UmlName string `json:"uml_name" binding:"required"`
	ProjID  uint64 `json:"proj_id"`
}

type CreateUmlA struct {
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

type UploadUmlQ struct {
	UmlID   uint64 `json:"uml_id" binding:"required"`
	UmlData string `json:"uml_data" binding:"omitempty"`
}

type UploadUmlA struct {
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
