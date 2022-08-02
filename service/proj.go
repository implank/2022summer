package service

import (
	"2022summer/global"
	"2022summer/model"
)

func QueryProjByProjName(projName string) (proj model.Proj, notFound bool) {
	notFound = global.DB.Where("proj_name = ?", projName).First(&proj).RecordNotFound()
	return proj, notFound
}

func QueryProjByProjID(projID uint64) (proj model.Proj, notFound bool) {
	notFound = global.DB.Where("proj_id = ?", projID).First(&proj).RecordNotFound()
	return proj, notFound
}

func CreateProj(proj *model.Proj) (err error) {
	if err = global.DB.Create(&proj).Error; err != nil {
		return err
	}
	return
}

func UpdateProj(proj *model.Proj) (err error) {
	err = global.DB.Save(proj).Error
	return err
}

func UpdateProjStatus(proj *model.Proj) (err error) {
	// 移入或移出回收站: 先改变 Prototype, Uml, Document 的 Status, 再保存 Proj
	// 【注意】传入的 Proj 的 Status 已被修改
	global.DB.Where("proj_id = ?", proj.ProjID).Updates(model.Prototype{Status: proj.Status})
	global.DB.Where("proj_id = ?", proj.ProjID).Updates(model.Uml{Status: proj.Status})
	global.DB.Where("proj_id = ?", proj.ProjID).Updates(model.Document{Status: proj.Status})
	err = global.DB.Save(proj).Error
	return err
}

func DeleteProj(proj *model.Proj) (err error) {
	// 删除回收站中的 Proj: 先删除 Prototype, Uml, Document, 再删除 Proj
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(model.Prototype{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(model.Uml{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(model.Document{})
	err = global.DB.Delete(&proj).Error
	return err
}

func QueryPrototypeByPrototypeName(prototypeName string, projID uint64) (prototype model.Prototype, notFound bool) {
	// 同一项目中不能有同名原型
	notFound = global.DB.Where("prototype_name = ? and proj_id = ?",
		prototypeName, projID).First(&prototype).RecordNotFound()
	return prototype, notFound
}

func QueryPrototypeByPrototypeID(prototypeID uint64) (prototype model.Prototype, notFound bool) {
	notFound = global.DB.Where("prototype_id = ?", prototypeID).First(&prototype).RecordNotFound()
	return prototype, notFound
}

func CreatePrototype(prototype *model.Prototype) (err error) {
	if err = global.DB.Create(&prototype).Error; err != nil {
		return err
	}
	return
}

func UpdatePrototype(prototype *model.Prototype) (err error) {
	err = global.DB.Save(prototype).Error
	return err
}

func DeletePrototype(prototype *model.Prototype) (err error) { // 用于删除回收站中的原型, 如果移到回收站, 应使用 Update, 后同
	err = global.DB.Delete(&prototype).Error
	return err
}

func QueryUmlByUmlName(umlName string, projID uint64) (uml model.Uml, notFound bool) {
	// 同一项目中不能有同名绘制图
	notFound = global.DB.Where("uml_name = ? and proj_id = ?", umlName, projID).First(&uml).RecordNotFound()
	return uml, notFound
}

func QueryUmlByUmlID(umlID uint64) (uml model.Uml, notFound bool) {
	notFound = global.DB.Where("uml_id = ?", umlID).First(&uml).RecordNotFound()
	return uml, notFound
}

func CreateUml(uml *model.Uml) (err error) {
	if err = global.DB.Create(&uml).Error; err != nil {
		return err
	}
	return
}

func UpdateUml(uml *model.Uml) (err error) {
	err = global.DB.Save(uml).Error
	return err
}

func DeleteUml(uml *model.Uml) (err error) {
	err = global.DB.Delete(&uml).Error
	return err
}

func QueryDocumentByDocumentName(documentName string, projID uint64) (document model.Document, notFound bool) {
	// 同一项目中不能有同名文档
	notFound = global.DB.Where("document_name = ? and proj_id = ?",
		documentName, projID).First(&document).RecordNotFound()
	return document, notFound
}

func QueryDocumentByDocumentID(documentID uint64) (document model.Document, notFound bool) {
	notFound = global.DB.Where("document_id = ?", documentID).First(&document).RecordNotFound()
	return document, notFound
}

func CreateDocument(document *model.Document) (err error) {
	if err = global.DB.Create(&document).Error; err != nil {
		return err
	}
	return
}

func UpdateDocument(document *model.Document) (err error) {
	err = global.DB.Save(document).Error
	return err
}

func DeleteDocument(document *model.Document) (err error) {
	err = global.DB.Delete(&document).Error
	return err
}
