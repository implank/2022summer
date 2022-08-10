package response

import (
	"2022summer/model/database"
)

type CreateProjQ struct {
	ProjName string `json:"proj_name" binding:"required"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
	GroupID  uint64 `json:"group_id" binding:"required"`
	Top      int    `json:"top" binding:"omitempty"`
}

type CreateProjA struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type UpdateProjQ struct {
	ProjID   uint64 `json:"proj_id" binding:"required"`
	ProjName string `json:"proj_name" binding:"omitempty"`
	ProjInfo string `json:"proj_info" binding:"omitempty"`
	Top      int    `json:"top" binding:"omitempty"`
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
	OrderBy int    `json:"order_by"` // 1 按创建时间排序、2 按修改时间排序、3 按编辑次数排序
	IsDesc  bool   `json:"is_desc"`  // true 降序、false 升序
}

type GetProjAllA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjCreateQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
	OrderBy int    `json:"order_by"` // 1 按创建时间排序、2 按修改时间排序、3 按编辑次数排序
	IsDesc  bool   `json:"is_desc"`  // true 降序、false 升序
}

type GetProjCreateA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjJoinQ struct {
	GroupID uint64 `json:"group_id" binding:"required"`
	OrderBy int    `json:"order_by"` // 1 按创建时间排序、2 按修改时间排序、3 按编辑次数排序
	IsDesc  bool   `json:"is_desc"`  // true 降序、false 升序
}

type GetProjJoinA struct {
	Message string          `json:"message"`
	Success bool            `json:"success"`
	Count   uint64          `json:"count"`
	Projs   []database.Proj `json:"projs"`
}

type GetProjByNameQ struct {
	ProjName string `json:"proj_name" binding:"omitempty"`
	GroupID  uint64 `json:"group_id" binding:"required"`
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

type CopyProjQ struct {
	ProjID uint64 `json:"proj_id" binding:"required"`
}

type CopyProjA struct {
	Message string        `json:"message"`
	Success bool          `json:"success"`
	Proj    database.Proj `json:"proj"`
}
