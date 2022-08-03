package response

import "2022summer/model"

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
