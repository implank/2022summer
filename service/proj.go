package service

import (
	"2022summer/global"
	"2022summer/model"
)

func QueryProjByProjName(projName string) (proj model.Proj, notFound bool) {
	notFound = global.DB.Where("proj_name = ?", projName).First(&proj).RecordNotFound()
	return proj, notFound
}

func GetProjsByProjNameBur(projName string) (projs []model.Proj) { // 模糊搜索
	global.DB.Where("proj_name like '%" + projName + "%' and status = 1").Find(&projs).RecordNotFound()
	return projs
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
	// 【注意】传入的 Proj 的 Status 应该是已被修改的
	global.DB.Model(model.Prototype{}).Where("proj_id = ?", proj.ProjID).Updates(model.Prototype{Status: proj.Status})
	global.DB.Model(model.Uml{}).Where("proj_id = ?", proj.ProjID).Updates(model.Uml{Status: proj.Status})
	global.DB.Model(model.Document{}).Where("proj_id = ?", proj.ProjID).Updates(model.Document{Status: proj.Status})
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

func DeletePrototype(prototype *model.Prototype) (err error) {
	// 用于删除回收站中的原型, 如果移到回收站, 应使用 Update, 如果从回收站中恢复, 应先判断 Proj 是否在回收站中, 再使用 Update, 后同
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

func GetFilesByNameBur(Name string) (prototypes []model.Prototype, umls []model.Uml, documents []model.Document) {
	// 模糊搜索设计原型、Uml 和文档
	global.DB.Where("prototype_name like '%" + Name + "%'").Find(&prototypes).RecordNotFound()
	global.DB.Where("uml_name like '%" + Name + "%'").Find(&umls).RecordNotFound()
	global.DB.Where("document_name like '%" + Name + "%'").Find(&documents).RecordNotFound()
	return prototypes, umls, documents
}

func GetProjPrototypes(projID uint64, status int) (prototypes []model.Prototype) {
	// status = 2 时, 查找回收站的项目, 后同
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&prototypes).RecordNotFound()
	return prototypes
}

func GetProjUmls(projID uint64, status int) (umls []model.Uml) {
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&umls).RecordNotFound()
	return umls
}

func GetProjDocuments(projID uint64, status int) (documents []model.Document) {
	global.DB.Where("proj_id = ? and status = ?", projID, status).Find(&documents).RecordNotFound()
	return documents
}

func GetUserProjs(userID uint64, status int, op int) (projs []model.Proj) {
	// 查找该用户在所有团队中的项目
	// op = 1 时查找"我创建的", op = 2 时查找"我参与的", op = 3 时查找"全部项目"
	if op == 1 {
		global.DB.Where("user_id = ? and status = ?", userID, status).Find(&projs).RecordNotFound()
	} else if op == 2 {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE projs.group_id = identities.group_id "+
			"AND identities.user_id = ? AND projs.user_id != ? "+
			"AND projs.status = ?;", userID, userID, status).Find(&projs).RecordNotFound()
	} else {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE projs.group_id = identities.group_id "+
			"AND identities.user_id = ? AND projs.status = ?;", userID, status).Find(&projs).RecordNotFound()
	}
	return projs
}

func GetUserProjsInGroup(userID uint64, groupID uint64, status int, op int) (projs []model.Proj) {
	// 查找该用户在某个团队中的项目
	// op = 1 时查找"我创建的", op = 2 时查找"我参与的", op = 3 时查找"全部项目"
	if op == 1 {
		global.DB.Where("user_id = ? and group_id = ? and status = ?",
			userID, groupID, status).Find(&projs).RecordNotFound()
	} else if op == 2 {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? AND projs.user_id != ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?;", userID, userID, groupID, groupID, status).Find(&projs).RecordNotFound()
	} else {
		global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?;", userID, groupID, groupID, status).Find(&projs).RecordNotFound()
	}
	return projs
}
