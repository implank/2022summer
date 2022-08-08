package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func GetFilesByNameBur(Name string, status int) (ppages []database.PPage, umls []database.Uml, documents []database.Document) {
	// 模糊搜索设计原型、Uml 和文档
	global.DB.Order("p_page_id DESC").Where("p_page_name like '%"+Name+"%' and status = ?", status).Find(&ppages).RecordNotFound()
	global.DB.Order("uml_id DESC").Where("uml_name like '%"+Name+"%' and status = ?", status).Find(&umls).RecordNotFound()
	global.DB.Order("document_id DESC").Where("document_name like '%"+Name+"%' and status = ?", status).Find(&documents).RecordNotFound()
	return ppages, umls, documents
}
