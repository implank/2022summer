package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func GetProjPrototypesInBin(groupID uint64) (prototypes []database.Prototype) {
	// 获取某一团队在回收站的设计原型，但并不是整个项目都在回收站，后同
	global.DB.Raw("SELECT * FROM prototypes, projs "+
		"WHERE projs.proj_id = prototypes.proj_id "+
		"AND projs.status = 1 AND prototypes.status = 2 "+
		"AND projs.group_id = ?;", groupID).Find(&prototypes).RecordNotFound()
	return prototypes
}

func GetProjUmlsInBin(groupID uint64) (umls []database.Uml) {
	global.DB.Raw("SELECT * FROM umls, projs "+
		"WHERE projs.proj_id = umls.proj_id "+
		"AND projs.status = 1 AND umls.status = 2 "+
		"AND projs.group_id = ?;", groupID).Find(&umls).RecordNotFound()
	return umls
}

func GetProjDocumentsInBin(groupID uint64) (documents []database.Document) {
	global.DB.Raw("SELECT * FROM documents, projs "+
		"WHERE projs.proj_id = documents.proj_id "+
		"AND projs.status = 1 AND documents.status = 2 "+
		"AND projs.group_id = ?;", groupID).Find(&documents).RecordNotFound()
	return documents
}
