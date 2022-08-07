package response

import (
	"2022summer/model/database"
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
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetProjAllA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjCreateQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetProjCreateA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjJoinQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
}

type GetProjJoinA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjByNameQ struct {
	ProjName string `json:"proj_name" binding:"omitempty"`
}

type GetProjByNameA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjByIDQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjByIDA struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Proj    database.Proj `json:"proj"`
}
