package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func GetFilesByNameBur(Name string, status int, groupID uint64) (ppages []database.PPageID, umls []database.UmlID, documents []database.DocumentID) {
	// 模糊搜索设计原型、Uml 和文档
	// global.DB.Order("p_page_id DESC").Where("p_page_name like '%"+Name+"%' and status = ?", status).Find(&ppages).RecordNotFound()
	// global.DB.Order("uml_id DESC").Where("uml_name like '%"+Name+"%' and status = ?", status).Find(&umls).RecordNotFound()
	// global.DB.Order("document_id DESC").Where("document_name like '%"+Name+"%' and status = ?", status).Find(&documents).RecordNotFound()
	global.DB.Raw("SELECT * FROM p_pages, projs "+
		"WHERE p_pages.proj_id = projs.proj_id "+
		"AND projs.group_id = ? "+
		"AND p_page_name like '%"+Name+"%' "+
		"AND p_pages.status = ? "+
		"ORDER BY p_page_id DESC;", groupID, status).Find(&ppages).RecordNotFound()
	global.DB.Raw("SELECT * FROM umls, projs "+
		"WHERE umls.proj_id = projs.proj_id "+
		"AND projs.group_id = ? "+
		"AND uml_name like '%"+Name+"%' "+
		"AND umls.status = ? "+
		"ORDER BY uml_id DESC;", groupID, status).Find(&umls).RecordNotFound()
	global.DB.Raw("SELECT * FROM documents, projs "+
		"WHERE documents.proj_id = projs.proj_id "+
		"AND projs.group_id = ? "+
		"AND document_name like '%"+Name+"%' "+
		"AND documents.status = ? "+
		"ORDER BY documents.document_id DESC;", groupID, status).Find(&documents).RecordNotFound()
	return ppages, umls, documents
}
