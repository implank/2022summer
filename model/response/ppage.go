package response

import "2022summer/model/database"

type GetProjPPagesQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type GetProjPPagesA struct {
	Message string             `json:"message"`
	Success bool               `json:"success"`
	Count   uint64             `json:"count"`
	PPages  []database.PPageID `json:"ppages"`
}

type GetPPageByIDQ struct {
	PPageID uint64 `json:"ppage_id" binding:"required"`
}

type GetPPageByIDA struct {
	Message string         `json:"message"`
	Success bool           `json:"success"`
	PPage   database.PPage `json:"ppage"`
}

type CreatePPageQ struct {
	PPageName string `json:"ppage_name" binding:"required"`
	PPageData string `json:"ppage_data" binding:"omitempty"`
	ProjID    uint64 `json:"proj_id" binding:"required"`
}

type CreatePPageA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdatePPageQ struct {
	PPageID   uint64 `json:"ppage_id" binding:"required"`
	PPageName string `json:"ppage_name" binding:"omitempty"`
	PPageData string `json:"ppage_data" binding:"omitempty"`
}

type UpdatePPageA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type MovePPageToBinQ struct {
	PPageID uint64 `json:"ppage_id" binding:"required"`
}

type MovePPageToBinA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
