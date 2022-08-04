package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func QueryPrototypeByPrototypeName(prototypeName string, projID uint64) (prototype database.Prototype, notFound bool) {
	// 同一项目中不能有同名原型
	notFound = global.DB.Where("prototype_name = ? and proj_id = ?",
		prototypeName, projID).First(&prototype).RecordNotFound()
	return prototype, notFound
}

func QueryPrototypeByPrototypeID(prototypeID uint64) (prototype database.Prototype, notFound bool) {
	notFound = global.DB.Where("prototype_id = ?", prototypeID).First(&prototype).RecordNotFound()
	return prototype, notFound
}

func CreatePrototype(prototype *database.Prototype) (err error) {
	if err = global.DB.Create(&prototype).Error; err != nil {
		return err
	}
	return
}

func UpdatePrototype(prototype *database.Prototype) (err error) {
	err = global.DB.Save(prototype).Error
	return err
}

func DeletePrototype(prototype *database.Prototype) (err error) {
	// 用于删除回收站中的原型, 如果移到回收站, 应使用 Update, 如果从回收站中恢复, 应先判断 Proj 是否在回收站中, 再使用 Update, 后同
	err = global.DB.Delete(&prototype).Error
	return err
}

func QueryUmlByUmlName(umlName string, projID uint64) (uml database.Uml, notFound bool) {
	// 同一项目中不能有同名绘制图
	notFound = global.DB.Where("uml_name = ? and proj_id = ?", umlName, projID).First(&uml).RecordNotFound()
	return uml, notFound
}

func QueryUmlByUmlID(umlID uint64) (uml database.Uml, notFound bool) {
	notFound = global.DB.Where("uml_id = ?", umlID).First(&uml).RecordNotFound()
	return uml, notFound
}

func CreateUml(uml *database.Uml) (err error) {
	if err = global.DB.Create(&uml).Error; err != nil {
		return err
	}
	return
}

func UpdateUml(uml *database.Uml) (err error) {
	err = global.DB.Save(uml).Error
	return err
}

func DeleteUml(uml *database.Uml) (err error) {
	err = global.DB.Delete(&uml).Error
	return err
}

func QueryDocumentByDocumentName(documentName string, projID uint64) (document database.Document, notFound bool) {
	// 同一项目中不能有同名文档
	notFound = global.DB.Where("document_name = ? and proj_id = ?",
		documentName, projID).First(&document).RecordNotFound()
	return document, notFound
}

func QueryDocumentByDocumentID(documentID uint64) (document database.Document, notFound bool) {
	notFound = global.DB.Where("document_id = ?", documentID).First(&document).RecordNotFound()
	return document, notFound
}

func CreateDocument(document *database.Document) (err error) {
	if err = global.DB.Create(&document).Error; err != nil {
		return err
	}
	return
}

func UpdateDocument(document *database.Document) (err error) {
	err = global.DB.Save(document).Error
	return err
}

func DeleteDocument(document *database.Document) (err error) {
	err = global.DB.Delete(&document).Error
	return err
}

func GetFilesByNameBur(Name string) (prototypes []database.Prototype, umls []database.Uml, documents []database.Document) {
	// 模糊搜索设计原型、Uml 和文档
	global.DB.Where("prototype_name like '%" + Name + "%'").Find(&prototypes).RecordNotFound()
	global.DB.Where("uml_name like '%" + Name + "%'").Find(&umls).RecordNotFound()
	global.DB.Where("document_name like '%" + Name + "%'").Find(&documents).RecordNotFound()
	return prototypes, umls, documents
}

func GetProjPrototypes(projID uint64, status int) (prototypes []database.Prototype) {
	// status = 2 时, 查找回收站的项目, 后同
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&prototypes).RecordNotFound()
	return prototypes
}

func GetProjUmls(projID uint64, status int) (umls []database.Uml) {
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&umls).RecordNotFound()
	return umls
}

func GetProjDocuments(projID uint64, status int) (documents []database.Document) {
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&documents).RecordNotFound()
	return documents
}
