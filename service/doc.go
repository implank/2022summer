package service

import (
	"2022summer/global"
	"2022summer/model/database"
	"github.com/jinzhu/gorm"
)

func GetProjDocuments(projID uint64, status int) (documents []database.DocumentID) {
	global.DB.Table("documents").Select("document_id, document_name").Order("document_id DESC").
		Where("proj_id = ? and status = ? and is_dir != 1", projID, status).Find(&documents).RecordNotFound()
	return documents
}

func QueryDocumentByDocumentName(documentName string, projID uint64) (document database.Document, notFound bool) {
	// 同一项目中不能有同名文档
	notFound = global.DB.Where("document_name = ? and proj_id = ? and is_dir != 1",
		documentName, projID).First(&document).RecordNotFound()
	return document, notFound
}

func QueryDocumentByDocumentID(documentID uint64) (document database.Document, notFound bool) {
	notFound = global.DB.Where("document_id = ?", documentID).First(&document).RecordNotFound()
	return document, notFound
}

func CreateDocument(document *database.Document) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", document.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	if err = global.DB.Create(&document).Error; err != nil {
		return err
	}
	return
}

func UpdateDocument(document *database.Document) (err error) {
	global.DB.Model(database.Proj{}).Where("proj_id = ?", document.ProjID).
		Update("edit_times", gorm.Expr("edit_times + ?", 1))
	err = global.DB.Save(document).Error
	return err
}

func DeleteDocument(document *database.Document) (err error) {
	err = global.DB.Delete(&document).Error
	return err
}
func GetDocumentsInDir(dirID uint64) (documents []database.Document) {
	global.DB.Order("document_id DESC").Where("dir_id = ? and status != 2", dirID).Find(&documents).RecordNotFound()
	return documents
}
