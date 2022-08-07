package response

import (
	"2022summer/model/database"
)

type GetFilesByNameQ struct {
	Name string `json:"name"`
}

type GetFilesByNameA struct {
	Message        string              `json:"message"`
	Success        bool                `json:"success"`
	CountPPage     uint64              `json:"count_ppage"`
	PPage          []database.PPage    `json:"ppage"`
	CountUmls      uint64              `json:"count_umls"`
	Umls           []database.Uml      `json:"umls"`
	CountDocuments uint64              `json:"count_documents"`
	Documents      []database.Document `json:"documents"`
}
