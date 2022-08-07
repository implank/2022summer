package service

import (
	"2022summer/global"
	"2022summer/model/database"
)

func GetProjDocuments(projID uint64, status int) (documents []database.Document) {
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&documents).RecordNotFound()
	return documents
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
