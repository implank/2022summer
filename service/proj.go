package service

import (
	"2022summer/global"
	"2022summer/model/database"
	"github.com/jinzhu/gorm"
	"strconv"
)

func QueryProjByProjName(projName string, groupID uint64) (proj database.Proj, notFound bool) { // 同一组的项目不能重名
	notFound = global.DB.Where("proj_name = ? and group_id = ?", projName, groupID).First(&proj).RecordNotFound()
	return proj, notFound
}

func GetProjsByProjNameBur(projName string, groupID uint64) (projs []database.Proj) { // 模糊搜索
	global.DB.Order("proj_id DESC").Where("proj_name like '%"+projName+"%' and status = 1 and group_id = ?", groupID).Find(&projs).RecordNotFound()
	return projs
}

func QueryProjByProjID(projID uint64) (proj database.Proj, notFound bool) {
	notFound = global.DB.Where("proj_id = ?", projID).First(&proj).RecordNotFound()
	return proj, notFound
}

func CreateProj(proj *database.Proj) (err error) {
	if err = global.DB.Create(&proj).Error; err != nil {
		return err
	}
	return
}

func UpdateProj(proj *database.Proj) (err error) {
	proj.EditTimes += 1
	err = global.DB.Save(proj).Error
	return err
}

func UpdateProjStatus(proj *database.Proj) (err error) {
	// 移入或移出回收站: 先改变 PPage, Uml, Document 的 Status, 再保存 Proj
	// 【注意】传入的 Proj 的 Status 应该是已被修改的
	global.DB.Model(database.PPage{}).Where("proj_id = ?", proj.ProjID).Updates(database.PPage{Status: proj.Status})
	global.DB.Model(database.Uml{}).Where("proj_id = ?", proj.ProjID).Updates(database.Uml{Status: proj.Status})
	global.DB.Model(database.Document{}).Where("proj_id = ?", proj.ProjID).Updates(database.Document{Status: proj.Status})
	proj.EditTimes += 1
	err = global.DB.Save(proj).Error
	return err
}

func DeleteProj(proj *database.Proj) (err error) {
	// 删除回收站中的 Proj: 先删除 PPage, Uml, Document, 再删除 Proj
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.PPage{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.Uml{})
	global.DB.Where("proj_id = ?", proj.ProjID).Delete(database.Document{})
	err = global.DB.Delete(&proj).Error
	return err
}

func GetUserProjs(userID uint64, status int, op int) (projs []database.Proj) {
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

func GetUserProjsInGroup(userID uint64, groupID uint64, status int, op int, orderBy int, isDesc bool) (projs []database.Proj) {
	// 查找该用户在某个团队中的项目
	// op = 1 时查找"我创建的", op = 2 时查找"我参与的", op = 3 时查找"全部项目"
	var query *gorm.DB
	if op == 1 {
		query = global.DB.Where("user_id = ? and group_id = ? and status = ?",
			userID, groupID, status)
	} else if op == 2 {
		query = global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? AND projs.user_id != ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?", userID, userID, groupID, groupID, status)
	} else {
		query = global.DB.Raw("SELECT * FROM projs, identities "+
			"WHERE identities.user_id = ? "+
			"AND projs.group_id = ? AND identities.group_id = ? "+
			"AND projs.status = ?", userID, groupID, groupID, status)
	}
	query = query.Order("top DESC")
	var str string
	if orderBy == 1 {
		str = "created_at"
	} else if orderBy == 2 {
		str = "updated_at "
	} else {
		str = "edit_times"
	}
	if isDesc {
		str += " DESC"
	}
	query = query.Order(str)
	query.Find(&projs).RecordNotFound()
	return projs
}

func CopyProj(projSrcID uint64, userID uint64) (projDst database.Proj, err error) {
	tx := global.DB.Begin()
	projSrc, _ := QueryProjByProjID(projSrcID)
	cnt := 0
	name := projSrc.ProjName + "的副本"
	for {
		notFound := global.DB.Where("proj_name = ? and group_id = ?", name, projSrc.GroupID).
			First(&database.Proj{}).RecordNotFound()
		if notFound {
			break
		}
		cnt += 1
		name = projSrc.ProjName + "的副本" + strconv.Itoa(cnt)
	}
	projDst = database.Proj{
		ProjName: name,
		ProjInfo: projSrc.ProjInfo,
		Status:   projSrc.Status,
		GroupID:  projSrc.GroupID,
		UserID:   userID,
		Top:      projSrc.Top}
	if err := global.DB.Create(&projDst).Error; err != nil {
		tx.Rollback()
		return projDst, err
	}
	var ppages []database.PPage
	var umls []database.Uml
	var documents []database.Document
	global.DB.Where("proj_id = ?", projSrcID).Find(&ppages)
	global.DB.Where("proj_id = ?", projSrcID).Find(&umls)
	global.DB.Where("proj_id = ?", projSrcID).Find(&documents)
	for _, value := range ppages {
		tmp := database.PPage{
			PPageName: value.PPageName,
			PPageData: value.PPageData,
			PPageURL:  "",
			Status:    value.Status,
			ProjID:    projDst.ProjID}
		if err := global.DB.Create(&tmp).Error; err != nil {
			tx.Rollback()
			return projDst, err
		}
	}
	for _, value := range umls {
		tmp := database.Uml{
			UmlName: value.UmlName,
			UmlURL:  "",
			Status:  value.Status,
			ProjID:  projDst.ProjID}
		if err := global.DB.Create(&tmp).Error; err != nil {
			tx.Rollback()
			return projDst, err
		}
	}
	for _, value := range documents {
		tmp := database.Document{
			DocumentName: value.DocumentName,
			//DocumentURL:  "",
			Status:  value.Status,
			ProjID:  projDst.ProjID,
			Content: value.Content}
		if err := global.DB.Create(&tmp).Error; err != nil {
			tx.Rollback()
			return projDst, err
		}
	}
	tx.Commit()
	return projDst, err
}
